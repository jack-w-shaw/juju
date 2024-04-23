// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package bootstrap

import (
	"context"
	"fmt"
	"os"

	"github.com/juju/errors"
	"gopkg.in/tomb.v2"

	"github.com/juju/juju/agent"
	"github.com/juju/juju/controller"
	"github.com/juju/juju/core/flags"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/objectstore"
	"github.com/juju/juju/core/permission"
	usererrors "github.com/juju/juju/domain/access/errors"
	userservice "github.com/juju/juju/domain/access/service"
	domainstorage "github.com/juju/juju/domain/storage"
	storageerrors "github.com/juju/juju/domain/storage/errors"
	storageservice "github.com/juju/juju/domain/storage/service"
	"github.com/juju/juju/internal/auth"
	"github.com/juju/juju/internal/bootstrap"
	"github.com/juju/juju/internal/cloudconfig/instancecfg"
	"github.com/juju/juju/internal/password"
	"github.com/juju/juju/internal/storage"
	"github.com/juju/juju/internal/worker/gate"
	"github.com/juju/juju/state/binarystorage"
)

const (
	// States which report the state of the worker.
	stateStarted   = "started"
	stateCompleted = "completed"
)

// LegacyState is the interface that is used to get the legacy state (mongo).
type LegacyState interface {
	// ControllerModelUUID returns the UUID of the model that was
	// bootstrapped.  This is the only model that can have controller
	// machines.  The owner of this model is also considered "special", in
	// that they are the only user that is able to create other users
	// (until we have more fine grained permissions), and they cannot be
	// disabled.
	ControllerModelUUID() string
	// ToolsStorage returns a new binarystorage.StorageCloser that stores tools
	// metadata in the "juju" database "toolsmetadata" collection.
	ToolsStorage(store objectstore.ObjectStore) (binarystorage.StorageCloser, error)
}

// WorkerConfig encapsulates the configuration options for the
// bootstrap worker.
type WorkerConfig struct {
	Agent                   agent.Agent
	ObjectStoreGetter       ObjectStoreGetter
	ControllerConfigService ControllerConfigService
	CredentialService       CredentialService
	CloudService            CloudService
	UserService             UserService
	StorageService          StorageService
	ProviderRegistry        storage.ProviderRegistry
	ApplicationService      ApplicationService
	FlagService             FlagService
	NetworkService          NetworkService
	BootstrapUnlocker       gate.Unlocker
	AgentBinaryUploader     AgentBinaryBootstrapFunc
	ControllerCharmDeployer ControllerCharmDeployerFunc
	PopulateControllerCharm PopulateControllerCharmFunc
	CharmhubHTTPClient      HTTPClient
	UnitPassword            string
	NewEnviron              NewEnvironFunc
	BootstrapAddresses      BootstrapAddressesFunc
	BootstrapAddressFinder  BootstrapAddressFinderFunc
	LoggerFactory           LoggerFactory

	// Deprecated: This is only here, until we can remove the state layer.
	SystemState SystemState
}

// Validate ensures that the config values are valid.
func (c *WorkerConfig) Validate() error {
	if c.Agent == nil {
		return errors.NotValidf("nil Agent")
	}
	if c.ObjectStoreGetter == nil {
		return errors.NotValidf("nil ObjectStoreGetter")
	}
	if c.ControllerConfigService == nil {
		return errors.NotValidf("nil ControllerConfigService")
	}
	if c.CredentialService == nil {
		return errors.NotValidf("nil CredentialService")
	}
	if c.CloudService == nil {
		return errors.NotValidf("nil CloudService")
	}
	if c.UserService == nil {
		return errors.NotValidf("nil UserService")
	}
	if c.StorageService == nil {
		return errors.NotValidf("nil StorageService")
	}
	if c.ApplicationService == nil {
		return errors.NotValidf("nil ApplicationService")
	}
	if c.BootstrapUnlocker == nil {
		return errors.NotValidf("nil BootstrapUnlocker")
	}
	if c.AgentBinaryUploader == nil {
		return errors.NotValidf("nil AgentBinaryUploader")
	}
	if c.FlagService == nil {
		return errors.NotValidf("nil FlagService")
	}
	if c.NetworkService == nil {
		return errors.NotValidf("nil NetworkService")
	}
	if c.ControllerCharmDeployer == nil {
		return errors.NotValidf("nil ControllerCharmDeployer")
	}
	if c.PopulateControllerCharm == nil {
		return errors.NotValidf("nil PopulateControllerCharm")
	}
	if c.CharmhubHTTPClient == nil {
		return errors.NotValidf("nil CharmhubHTTPClient")
	}
	if c.LoggerFactory == nil {
		return errors.NotValidf("nil LoggerFactory")
	}
	if c.SystemState == nil {
		return errors.NotValidf("nil SystemState")
	}
	if c.BootstrapAddressFinder == nil {
		return errors.NotValidf("nil BootstrapAddressFinder")
	}
	if c.NewEnviron == nil {
		return errors.NotValidf("nil NewEnviron")
	}
	if c.BootstrapAddresses == nil {
		return errors.NotValidf("nil BootstrapAddresses")
	}
	return nil
}

type bootstrapWorker struct {
	internalStates chan string
	cfg            WorkerConfig
	logger         Logger
	tomb           tomb.Tomb
}

// NewWorker creates a new bootstrap worker.
func NewWorker(cfg WorkerConfig) (*bootstrapWorker, error) {
	return newWorker(cfg, nil)
}

func newWorker(cfg WorkerConfig, internalStates chan string) (*bootstrapWorker, error) {
	var err error
	if err = cfg.Validate(); err != nil {
		return nil, errors.Trace(err)
	}

	w := &bootstrapWorker{
		internalStates: internalStates,
		cfg:            cfg,
		logger:         cfg.LoggerFactory.Child("worker"),
	}
	w.tomb.Go(w.loop)
	return w, nil
}

// Kill stops the worker.
func (w *bootstrapWorker) Kill() {
	w.tomb.Kill(nil)
}

// Wait waits for the worker to stop and then returns the reason it was
// killed.
func (w *bootstrapWorker) Wait() error {
	return w.tomb.Wait()
}

func (w *bootstrapWorker) loop() error {
	// Report the initial started state.
	w.reportInternalState(stateStarted)

	ctx, cancel := w.scopedContext()
	defer cancel()

	// Insert all the initial users into the state.
	if err := w.seedInitialUsers(ctx); err != nil {
		return errors.Annotatef(err, "inserting initial users")
	}

	agentConfig := w.cfg.Agent.CurrentConfig()
	dataDir := agentConfig.DataDir()

	// Seed the agent binary to the object store.
	cleanup, err := w.seedAgentBinary(ctx, dataDir)
	if err != nil {
		return errors.Trace(err)
	}

	// Seed the controller charm to the object store.
	bootstrapParams, err := w.bootstrapParams(ctx, dataDir)
	if err != nil {
		return errors.Annotatef(err, "getting bootstrap params")
	}

	// Create the user specified storage pools.
	if err := w.seedStoragePools(ctx, bootstrapParams.StoragePools); err != nil {
		return errors.Annotate(err, "seeding storage pools")
	}

	controllerConfig, err := w.cfg.ControllerConfigService.ControllerConfig(ctx)
	if err != nil {
		return errors.Trace(err)
	}

	if err := w.seedControllerCharm(ctx, dataDir, bootstrapParams); err != nil {
		return errors.Trace(err)
	}

	// Retrieve controller addresses needed to set the API host ports.
	bootstrapAddresses, err := w.cfg.BootstrapAddressFinder(ctx, BootstrapAddressesConfig{
		BootstrapInstanceID:    bootstrapParams.BootstrapMachineInstanceId,
		SystemState:            w.cfg.SystemState,
		CloudService:           w.cfg.CloudService,
		CredentialService:      w.cfg.CredentialService,
		NewEnvironFunc:         w.cfg.NewEnviron,
		BootstrapAddressesFunc: w.cfg.BootstrapAddresses,
	})
	if err != nil {
		return errors.Trace(err)
	}
	servingInfo, ok := agentConfig.StateServingInfo()
	if !ok {
		return errors.Errorf("state serving information not available")
	}

	// Load spaces from the underlying substrate.
	if err := w.cfg.NetworkService.ReloadSpaces(ctx); err != nil {
		if !errors.Is(err, errors.NotSupported) {
			return errors.Trace(err)
		}
		w.logger.Debugf("reload spaces not supported due to a non-networking environement")
	}

	// Convert the provider addresses that we got from the bootstrap instance
	// to space ID decorated addresses.
	if err := w.initAPIHostPorts(ctx, controllerConfig, bootstrapAddresses, servingInfo.APIPort); err != nil {
		return errors.Trace(err)
	}

	// Set the bootstrap flag, to indicate that the bootstrap has completed.
	if err := w.cfg.FlagService.SetFlag(ctx, flags.BootstrapFlag, true, flags.BootstrapFlagDescription); err != nil {
		return errors.Trace(err)
	}

	// Cleanup only after the bootstrap flag has been set.
	cleanup()

	w.reportInternalState(stateCompleted)

	w.cfg.BootstrapUnlocker.Unlock()
	return nil
}

func (w *bootstrapWorker) seedInitialUsers(ctx context.Context) error {
	// Any failure should be retryable, so we can re-attempt to bootstrap.

	adminUser, err := w.cfg.UserService.GetUserByName(ctx, "admin")
	if err != nil {
		return errors.Annotatef(err, "getting admin user")
	}

	pass, err := password.RandomPassword()
	if err != nil {
		return errors.Annotatef(err, "generating metrics password")
	}
	password := auth.NewPassword(pass)
	_, _, err = w.cfg.UserService.AddUser(ctx, userservice.AddUserArg{
		Name:        "juju-metrics",
		DisplayName: "Juju Metrics",
		Password:    &password,
		CreatorUUID: adminUser.UUID,
		Permission:  permission.ControllerForAccess(permission.LoginAccess),
	})
	// User already exists, we don't need to do anything in this scenario.
	if errors.Is(err, usererrors.UserAlreadyExists) {
		return nil
	}
	return errors.Annotatef(err, "inserting initial users")
}

func (w *bootstrapWorker) seedStoragePools(ctx context.Context, poolParams map[string]storage.Attrs) error {
	storagePools, err := initialStoragePools(w.cfg.ProviderRegistry, poolParams)
	if err != nil {
		return errors.Trace(err)
	}
	for _, p := range storagePools {
		if err := w.cfg.StorageService.CreateStoragePool(ctx, p.Name(), p.Provider(), storageservice.PoolAttrs(p.Attrs())); err != nil {
			// Allow for bootstrap worker to have been restarted.
			if errors.Is(err, storageerrors.PoolAlreadyExists) {
				continue
			}
			return errors.Annotatef(err, "saving storage pool %q", p.Name())
		}
	}
	return nil
}

func (w *bootstrapWorker) reportInternalState(state string) {
	select {
	case <-w.tomb.Dying():
	case w.internalStates <- state:
	default:
	}
}

// initAPIHostPorts sets the initial API host/port addresses in state.
func (w *bootstrapWorker) initAPIHostPorts(ctx context.Context, controllerConfig controller.Config, pAddrs network.ProviderAddresses, apiPort int) error {
	allSpaces, err := w.cfg.NetworkService.GetAllSpaces(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	addrs, err := pAddrs.ToSpaceAddresses(allSpaces)
	if err != nil {
		return errors.Trace(err)
	}

	hostPorts := []network.SpaceHostPorts{network.SpaceAddressesWithPort(addrs, apiPort)}

	mgmtSpace := controllerConfig.JujuManagementSpace()
	hostPortsForAgents := w.filterHostPortsForManagementSpace(mgmtSpace, hostPorts, allSpaces)

	return w.cfg.SystemState.SetAPIHostPorts(controllerConfig, hostPorts, hostPortsForAgents)
}

// We filter the collection of API addresses based on the configured
// management space for the controller.
// If there is no space configured, or if one of the slices is filtered down
// to zero elements, just use the unfiltered slice for safety - we do not
// want to cut off communication to the controller based on erroneous config.
func (w *bootstrapWorker) filterHostPortsForManagementSpace(
	mgmtSpace string,
	apiHostPorts []network.SpaceHostPorts,
	allSpaces network.SpaceInfos,
) []network.SpaceHostPorts {
	var hostPortsForAgents []network.SpaceHostPorts

	if mgmtSpace == "" {
		hostPortsForAgents = apiHostPorts
	} else {
		mgmtSpaceInfo := allSpaces.GetByName(mgmtSpace)
		if mgmtSpaceInfo == nil {
			return apiHostPorts
		}
		hostPortsForAgents = make([]network.SpaceHostPorts, len(apiHostPorts))
		for i, apiHostPort := range apiHostPorts {
			filtered, addrsIsInSpace := apiHostPort.InSpaces(*mgmtSpaceInfo)
			if addrsIsInSpace {
				hostPortsForAgents[i] = filtered
			} else {
				w.logger.Warningf("API addresses %v not in the management space %s", apiHostPort, mgmtSpace)
				hostPortsForAgents[i] = apiHostPort
			}
		}
	}

	return hostPortsForAgents
}

// scopedContext returns a context that is in the scope of the worker lifetime.
// It returns a cancellable context that is cancelled when the action has
// completed.
func (w *bootstrapWorker) scopedContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	return w.tomb.Context(ctx), cancel
}

func (w *bootstrapWorker) seedAgentBinary(ctx context.Context, dataDir string) (func(), error) {
	objectStore, err := w.cfg.ObjectStoreGetter.GetObjectStore(ctx, w.cfg.SystemState.ControllerModelUUID())
	if err != nil {
		return nil, fmt.Errorf("failed to get object store: %w", err)
	}

	// Agent binary seeder will populate the tools for the agent.
	agentStorage := agentStorageShim{State: w.cfg.SystemState}
	cleanup, err := w.cfg.AgentBinaryUploader(ctx, dataDir, agentStorage, objectStore, w.cfg.LoggerFactory.Child("agentbinary"))
	if err != nil {
		return nil, errors.Trace(err)
	}

	return cleanup, nil
}

func (w *bootstrapWorker) seedControllerCharm(ctx context.Context, dataDir string, bootstrapArgs instancecfg.StateInitializationParams) error {
	controllerConfig, err := w.cfg.ControllerConfigService.ControllerConfig(ctx)
	if err != nil {
		return errors.Trace(err)
	}

	objectStore, err := w.cfg.ObjectStoreGetter.GetObjectStore(ctx, w.cfg.SystemState.ControllerModelUUID())
	if err != nil {
		return fmt.Errorf("failed to get object store: %w", err)
	}

	// Controller charm seeder will populate the charm for the controller.
	deployer, err := w.cfg.ControllerCharmDeployer(ControllerCharmDeployerConfig{
		StateBackend:                w.cfg.SystemState,
		ApplicationService:          w.cfg.ApplicationService,
		ObjectStore:                 objectStore,
		ControllerConfig:            controllerConfig,
		DataDir:                     dataDir,
		BootstrapMachineConstraints: bootstrapArgs.BootstrapMachineConstraints,
		ControllerCharmName:         bootstrapArgs.ControllerCharmPath,
		ControllerCharmChannel:      bootstrapArgs.ControllerCharmChannel,
		CharmhubHTTPClient:          w.cfg.CharmhubHTTPClient,
		UnitPassword:                w.cfg.UnitPassword,
		LoggerFactory:               w.cfg.LoggerFactory,
	})
	if err != nil {
		return errors.Trace(err)
	}

	return w.cfg.PopulateControllerCharm(ctx, deployer)
}

func (w *bootstrapWorker) bootstrapParams(ctx context.Context, dataDir string) (instancecfg.StateInitializationParams, error) {
	bootstrapParamsData, err := os.ReadFile(bootstrap.BootstrapParamsPath(dataDir))
	if err != nil {
		return instancecfg.StateInitializationParams{}, errors.Annotate(err, "reading bootstrap params file")
	}
	var args instancecfg.StateInitializationParams
	if err := args.Unmarshal(bootstrapParamsData); err != nil {
		return instancecfg.StateInitializationParams{}, errors.Trace(err)
	}
	return args, nil
}

// initialStoragePools extract any storage pools included with the bootstrap params.
func initialStoragePools(registry storage.ProviderRegistry, poolParams map[string]storage.Attrs) ([]*storage.Config, error) {
	var result []*storage.Config
	defaultStoragePools, err := domainstorage.DefaultStoragePools(registry)
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, p := range defaultStoragePools {
		result = append(result, p)
	}
	for name, attrs := range poolParams {
		pType, _ := attrs[domainstorage.StorageProviderType].(string)
		if pType == "" {
			return nil, errors.Errorf("missing provider type for storage pool %q", name)
		}
		delete(attrs, domainstorage.StoragePoolName)
		delete(attrs, domainstorage.StorageProviderType)
		pool, err := storage.NewConfig(name, storage.ProviderType(pType), attrs)
		if err != nil {
			return nil, errors.Trace(err)
		}
		result = append(result, pool)
	}
	return result, nil
}

type agentStorageShim struct {
	State SystemState
}

// AgentBinaryStorage returns the interface for the BinaryAgentStorage.
// This is currently a shim wrapper around the tools storage. That will be
// renamed once we re-implement the tools storage in dqlite.
func (s agentStorageShim) AgentBinaryStorage(objectStore objectstore.ObjectStore) (BinaryAgentStorage, error) {
	return s.State.ToolsStorage(objectStore)
}
