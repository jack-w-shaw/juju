// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package containerprovisioner_test

import (
	"context"
	"fmt"
	reflect "reflect"
	"sync"
	"time"

	"github.com/juju/errors"
	"github.com/juju/names/v5"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/version/v2"
	"github.com/juju/worker/v4/workertest"
	"github.com/kr/pretty"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/agent"
	"github.com/juju/juju/api"
	apiprovisioner "github.com/juju/juju/api/agent/provisioner"
	"github.com/juju/juju/controller"
	"github.com/juju/juju/core/arch"
	corebase "github.com/juju/juju/core/base"
	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/life"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/status"
	jujuversion "github.com/juju/juju/core/version"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/core/watcher/watchertest"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/environs/envcontext"
	"github.com/juju/juju/environs/instances"
	environmocks "github.com/juju/juju/environs/testing"
	"github.com/juju/juju/internal/cloudconfig/instancecfg"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	coretesting "github.com/juju/juju/internal/testing"
	"github.com/juju/juju/internal/tools"
	"github.com/juju/juju/internal/worker/containerprovisioner"
	"github.com/juju/juju/rpc/params"
)

type lxdProvisionerSuite struct {
	controllerAPI *MockControllerAPI
	machinesAPI   *MockMachinesAPI
	broker        *environmocks.MockEnviron

	containersCh  chan []string
	modelConfigCh chan struct{}

	provisionerStarted chan bool
}

func (s *lxdProvisionerSuite) setUpMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)
	s.controllerAPI = NewMockControllerAPI(ctrl)
	s.machinesAPI = NewMockMachinesAPI(ctrl)
	s.broker = environmocks.NewMockEnviron(ctrl)
	s.expectAuth()
	s.expectStartup(c)
	return ctrl
}

func (s *lxdProvisionerSuite) expectAuth() {
	s.controllerAPI.EXPECT().APIAddresses(gomock.Any()).Return([]string{"10.0.0.1"}, nil).AnyTimes()
	s.controllerAPI.EXPECT().ModelUUID(gomock.Any()).Return(coretesting.ModelTag.Id(), nil).AnyTimes()
	s.controllerAPI.EXPECT().CACert(gomock.Any()).Return(coretesting.CACert, nil).AnyTimes()
}

func (s *lxdProvisionerSuite) expectStartup(c *gc.C) {
	s.modelConfigCh = make(chan struct{})
	watchCfg := watchertest.NewMockNotifyWatcher(s.modelConfigCh)
	s.controllerAPI.EXPECT().WatchForModelConfigChanges(gomock.Any()).Return(watchCfg, nil)

	cfg := coretesting.CustomModelConfig(c, coretesting.Attrs{config.ProvisionerHarvestModeKey: config.HarvestDestroyed.String()})
	s.controllerAPI.EXPECT().ModelConfig(gomock.Any()).Return(cfg, nil).MaxTimes(2)

	s.provisionerStarted = make(chan bool)
	controllerCfg := coretesting.FakeControllerConfig()
	s.controllerAPI.EXPECT().ControllerConfig(gomock.Any()).DoAndReturn(func(context.Context) (controller.Config, error) {
		defer close(s.provisionerStarted)
		return controllerCfg, nil
	})
}

var _ = gc.Suite(&lxdProvisionerSuite{})

func (s *lxdProvisionerSuite) newLXDProvisioner(c *gc.C, ctrl *gomock.Controller) containerprovisioner.Provisioner {
	mTag := names.NewMachineTag("0")
	defaultPaths := agent.DefaultPaths
	defaultPaths.DataDir = c.MkDir()
	cfg, err := agent.NewAgentConfig(
		agent.AgentConfigParams{
			Paths:             defaultPaths,
			Tag:               mTag,
			UpgradedToVersion: jujuversion.Current,
			Password:          "password",
			Nonce:             "nonce",
			APIAddresses:      []string{"0.0.0.0:12345"},
			CACert:            coretesting.CACert,
			Controller:        coretesting.ControllerTag,
			Model:             coretesting.ModelTag,
		})
	c.Assert(err, jc.ErrorIsNil)

	s.containersCh = make(chan []string)
	m0 := &testMachine{containersCh: s.containersCh}
	s.machinesAPI.EXPECT().Machines(gomock.Any(), mTag).Return([]apiprovisioner.MachineResult{{
		Machine: m0,
	}}, nil)

	w, err := containerprovisioner.NewContainerProvisioner(
		instance.LXD, s.controllerAPI, s.machinesAPI, loggertesting.WrapCheckLog(c),
		cfg, s.broker, &mockToolsFinder{}, &mockDistributionGroupFinder{}, &credentialAPIForTest{})
	c.Assert(err, jc.ErrorIsNil)

	s.waitForProvisioner(c)
	return w
}

func (s *lxdProvisionerSuite) TestProvisionerStartStop(c *gc.C) {
	ctrl := s.setUpMocks(c)
	defer ctrl.Finish()

	p := s.newLXDProvisioner(c, ctrl)
	workertest.CleanKill(c, p)
}

func (s *lxdProvisionerSuite) sendMachineContainersChange(c *gc.C, ids ...string) {
	select {
	case s.containersCh <- ids:
	case <-time.After(coretesting.LongWait):
		c.Fatal("timed out sending containers change")
	}
}

func (s *lxdProvisionerSuite) waitForProvisioner(c *gc.C) {
	select {
	case <-s.provisionerStarted:
	case <-time.After(coretesting.LongWait):
		c.Fatalf("timed out waiting for provisioner")
	}
}

func (s *lxdProvisionerSuite) checkStartInstance(c *gc.C, m *testMachine) {
	for attempt := coretesting.LongAttempt.Start(); attempt.Next(); {
		_, err := m.InstanceId(context.Background())
		if err == nil {
			return
		}
	}
	c.Fatalf("machine %v not started", m.id)
}

func (s *lxdProvisionerSuite) TestContainerStartedAndStopped(c *gc.C) {
	ctrl := s.setUpMocks(c)
	defer ctrl.Finish()

	p := s.newLXDProvisioner(c, ctrl)
	defer workertest.CleanKill(c, p)

	cTag := names.NewMachineTag("0/lxd/666")

	c666 := &testMachine{id: "0/lxd/666"}
	s.broker.EXPECT().AllRunningInstances(gomock.Any()).Return([]instances.Instance{&testInstance{id: "inst-666"}}, nil).Times(2)
	s.machinesAPI.EXPECT().Machines(gomock.Any(), cTag).Return([]apiprovisioner.MachineResult{{
		Machine: c666,
	}}, nil).Times(2)
	s.machinesAPI.EXPECT().ProvisioningInfo(gomock.Any(), []names.MachineTag{cTag}).Return(params.ProvisioningInfoResults{
		Results: []params.ProvisioningInfoResult{{
			Result: &params.ProvisioningInfo{
				ControllerConfig: coretesting.FakeControllerConfig(),
				Constraints:      constraints.MustParse("mem=666G"),
				Base:             params.Base{Name: "ubuntu", Channel: "22.04"},
				Jobs:             []model.MachineJob{model.JobHostUnits},
			},
		}},
	}, nil)
	startArg := machineStartInstanceArg(c666.id)
	startArg.Constraints = constraints.MustParse("mem=666G")
	s.broker.EXPECT().StartInstance(gomock.Any(), newDefaultStartInstanceParamsMatcher(c, startArg)).Return(&environs.StartInstanceResult{
		Instance: &testInstance{id: "inst-666"},
	}, nil)

	s.sendMachineContainersChange(c, c666.Id())
	s.checkStartInstance(c, c666)

	s.broker.EXPECT().StopInstances(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx envcontext.ProviderCallContext, ids ...instance.Id) error {
		c.Assert(len(ids), gc.Equals, 1)
		c.Assert(ids[0], gc.DeepEquals, instance.Id("inst-666"))
		return nil
	})

	c666.SetLife(life.Dead)
	s.sendMachineContainersChange(c, c666.Id())
	s.waitForRemovalMark(c, c666)
}

func (s *lxdProvisionerSuite) TestKVMProvisionerObservesConfigChanges(c *gc.C) {
	ctrl := s.setUpMocks(c)
	defer ctrl.Finish()

	p := s.newLXDProvisioner(c, ctrl)
	defer workertest.CleanKill(c, p)

	s.assertProvisionerObservesConfigChanges(c, p, true)
}

func (s *lxdProvisionerSuite) TestKVMProvisionerObservesConfigChangesWorkerCount(c *gc.C) {
	ctrl := s.setUpMocks(c)
	defer ctrl.Finish()

	p := s.newLXDProvisioner(c, ctrl)
	defer workertest.CleanKill(c, p)

	s.assertProvisionerObservesConfigChangesWorkerCount(c, p, true)
}

// waitForRemovalMark waits for the supplied machine to be marked for removal.
func (s *lxdProvisionerSuite) waitForRemovalMark(c *gc.C, m *testMachine) {
	for attempt := coretesting.LongAttempt.Start(); attempt.Next(); {
		if m.GetMarkForRemoval() {
			return
		}
	}
	c.Fatalf("machine %q not marked for removal", m.id)
}

func (s *lxdProvisionerSuite) assertProvisionerObservesConfigChanges(c *gc.C, p containerprovisioner.Provisioner, container bool) {
	// Inject our observer into the provisioner
	cfgObserver := make(chan *config.Config)
	containerprovisioner.SetObserver(p, cfgObserver)

	attrs := coretesting.FakeConfig()
	attrs[config.ProvisionerHarvestModeKey] = config.HarvestDestroyed.String()
	modelCfg, err := config.New(config.UseDefaults, attrs)
	c.Assert(err, jc.ErrorIsNil)
	s.controllerAPI.EXPECT().ModelConfig(gomock.Any()).Return(modelCfg, nil)

	if !container {
		s.broker.EXPECT().SetConfig(gomock.Any(), modelCfg).Return(nil)
	}

	s.sendModelConfigChange(c)

	// Wait for the PA to load the new configuration. We wait for the change we expect
	// like this because sometimes we pick up the initial harvest config (destroyed)
	// rather than the one we change to (all).
	var received []int
	timeout := time.After(coretesting.LongWait)
	for {
		select {
		case newCfg := <-cfgObserver:
			if newCfg.ProvisionerHarvestMode() == config.HarvestDestroyed {
				return
			}
			received = append(received, newCfg.NumProvisionWorkers())
		case <-timeout:
			if len(received) == 0 {
				c.Fatalf("PA did not action config change")
			} else {
				c.Fatalf("timed out waiting for config to change to '%v', received %+v",
					config.HarvestDestroyed.String(), received)
			}
		}
	}
}

func (s *lxdProvisionerSuite) assertProvisionerObservesConfigChangesWorkerCount(c *gc.C, p containerprovisioner.Provisioner, container bool) {
	// Inject our observer into the provisioner
	cfgObserver := make(chan *config.Config)
	containerprovisioner.SetObserver(p, cfgObserver)

	attrs := coretesting.FakeConfig().Merge(coretesting.Attrs{
		config.ProvisionerHarvestModeKey: config.HarvestDestroyed.String(),
	})
	if container {
		attrs[config.NumContainerProvisionWorkersKey] = 20
	} else {
		attrs[config.NumProvisionWorkersKey] = 20
	}
	modelCfg, err := config.New(config.UseDefaults, attrs)
	c.Assert(err, jc.ErrorIsNil)
	s.controllerAPI.EXPECT().ModelConfig(gomock.Any()).Return(modelCfg, nil)

	if !container {
		s.broker.EXPECT().SetConfig(gomock.Any(), modelCfg).Return(nil)
	}

	s.sendModelConfigChange(c)

	// Wait for the PA to load the new configuration. We wait for the change we expect
	// like this because sometimes we pick up the initial harvest config (destroyed)
	// rather than the one we change to (all).
	var received []int
	timeout := time.After(coretesting.LongWait)
	for {
		select {
		case newCfg := <-cfgObserver:
			if container {
				if newCfg.NumContainerProvisionWorkers() == 20 {
					return
				}
				received = append(received, newCfg.NumContainerProvisionWorkers())
			} else {
				if newCfg.NumProvisionWorkers() == 20 {
					return
				}
				received = append(received, newCfg.NumProvisionWorkers())
			}
		case <-timeout:
			if len(received) == 0 {
				c.Fatalf("PA did not action config change")
			} else {
				c.Fatalf("timed out waiting for config to change to '%v', received %+v",
					20, received)
			}
		}
	}
}

func (s *lxdProvisionerSuite) sendModelConfigChange(c *gc.C) {
	select {
	case s.modelConfigCh <- struct{}{}:
	case <-time.After(coretesting.LongWait):
		c.Fatal("timed out sending model config change")
	}
}

type credentialAPIForTest struct{}

func (*credentialAPIForTest) InvalidateModelCredential(_ context.Context, reason string) error {
	return nil
}

var (
	startInstanceArgTemplate = environs.StartInstanceParams{
		ControllerUUID: coretesting.ControllerTag.Id(),
		Tools:          tools.List{{Version: version.MustParseBinary("2.99.0-ubuntu-amd64")}},
	}
	instanceConfigTemplate = instancecfg.InstanceConfig{
		ControllerTag:    coretesting.ControllerTag,
		ControllerConfig: coretesting.FakeControllerConfig(),
		Jobs:             []model.MachineJob{model.JobHostUnits},
		APIInfo: &api.Info{
			ModelTag: coretesting.ModelTag,
			Addrs:    []string{"10.0.0.1"},
			CACert:   coretesting.CACert,
		},
		Base:               corebase.MustParseBaseFromString("ubuntu@22.04"),
		TransientDataDir:   "/var/run/juju",
		DataDir:            "/var/lib/juju",
		LogDir:             "/var/log/juju",
		MetricsSpoolDir:    "/var/lib/juju/metricspool",
		CloudInitOutputLog: "/var/log/cloud-init-output.log",
		ImageStream:        "released",
	}
)

func machineStartInstanceArg(id string) *environs.StartInstanceParams {
	result := startInstanceArgTemplate
	instCfg := instanceConfigTemplate
	result.InstanceConfig = &instCfg
	tag := names.NewMachineTag(id)
	result.InstanceConfig.APIInfo.Tag = tag
	result.InstanceConfig.MachineId = id
	result.InstanceConfig.MachineAgentServiceName = fmt.Sprintf("jujud-%s", tag)
	return &result
}

func newDefaultStartInstanceParamsMatcher(c *gc.C, want *environs.StartInstanceParams) *startInstanceParamsMatcher {
	match := func(p environs.StartInstanceParams) bool {
		p.Abort = nil
		p.StatusCallback = nil
		p.CleanupCallback = nil
		if p.InstanceConfig != nil {
			cfgCopy := *p.InstanceConfig
			// The api password and machine nonce are generated to random values.
			// Just ensure they are not empty and tweak it so that the compare succeeds.
			if cfgCopy.APIInfo != nil {
				if cfgCopy.APIInfo.Password == "" {
					return false
				}
				cfgCopy.APIInfo.Password = want.InstanceConfig.APIInfo.Password
			}
			if cfgCopy.MachineNonce == "" {
				return false
			}
			cfgCopy.MachineNonce = ""
			p.InstanceConfig = &cfgCopy
		}
		if len(p.EndpointBindings) == 0 {
			p.EndpointBindings = nil
		}
		if len(p.Volumes) == 0 {
			p.Volumes = nil
		}
		if len(p.VolumeAttachments) == 0 {
			p.VolumeAttachments = nil
		}
		if len(p.ImageMetadata) == 0 {
			p.ImageMetadata = nil
		}
		match := reflect.DeepEqual(p, *want)
		if !match {
			c.Logf("got: %s\n", pretty.Sprint(p))
		}
		return match
	}
	m := newStartInstanceParamsMatcher(map[string]func(environs.StartInstanceParams) bool{
		fmt.Sprintf("core start params: %s\n", pretty.Sprint(*want)): match,
	})
	return m
}

func newStartInstanceParamsMatcher(
	matchers map[string]func(environs.StartInstanceParams) bool,
) *startInstanceParamsMatcher {
	if matchers == nil {
		matchers = make(map[string]func(environs.StartInstanceParams) bool)
	}
	return &startInstanceParamsMatcher{matchers: matchers}
}

// startInstanceParamsMatcher is a GoMock matcher that applies a collection of
// conditions to an environs.StartInstanceParams.
// All conditions must be true in order for a positive match.
type startInstanceParamsMatcher struct {
	matchers map[string]func(environs.StartInstanceParams) bool
	failMsg  string
}

func (m *startInstanceParamsMatcher) Matches(params interface{}) bool {
	siParams := params.(environs.StartInstanceParams)
	for msg, match := range m.matchers {
		if !match(siParams) {
			m.failMsg = msg
			return false
		}
	}
	return true
}

func (m *startInstanceParamsMatcher) String() string {
	return m.failMsg
}

type testInstance struct {
	instances.Instance
	id string
}

func (i *testInstance) Id() instance.Id {
	return instance.Id(i.id)
}

type testMachine struct {
	*apiprovisioner.Machine

	mu sync.Mutex

	id             string
	life           life.Value
	agentVersion   version.Number
	instance       *testInstance
	keepInstance   bool
	markForRemoval bool
	machineStatus  status.Status
	instStatus     status.Status
	instStatusMsg  string
	modStatusMsg   string
	password       string

	containersCh chan []string
}

func (m *testMachine) Id() string {
	return m.id
}

func (m *testMachine) String() string {
	return m.Id()
}

func (m *testMachine) Life() life.Value {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.life
}

func (m *testMachine) SetLife(life life.Value) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.life = life
}

func (m *testMachine) WatchContainers(_ context.Context, cType instance.ContainerType) (watcher.StringsWatcher, error) {
	if m.containersCh == nil {
		return nil, errors.Errorf("unexpected call to watch %q containers on %q", cType, m.id)
	}
	w := watchertest.NewMockStringsWatcher(m.containersCh)
	return w, nil
}

func (m *testMachine) InstanceId(context.Context) (instance.Id, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.instance == nil {
		return "", params.Error{Code: "not provisioned"}
	}
	return m.instance.Id(), nil
}

func (m *testMachine) InstanceNames() (instance.Id, string, error) {
	instId, err := m.InstanceId(context.Background())
	return instId, "", err
}

func (m *testMachine) KeepInstance(context.Context) (bool, error) {
	return m.keepInstance, nil
}

func (m *testMachine) MarkForRemoval(context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.markForRemoval = true
	return nil
}

func (m *testMachine) GetMarkForRemoval() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.markForRemoval
}

func (m *testMachine) Tag() names.Tag {
	return m.MachineTag()
}

func (m *testMachine) MachineTag() names.MachineTag {
	return names.NewMachineTag(m.id)
}

func (m *testMachine) SetInstanceStatus(ctx context.Context, status status.Status, message string, _ map[string]interface{}) error {
	m.mu.Lock()
	m.instStatus = status
	m.instStatusMsg = message
	m.mu.Unlock()
	return nil
}

func (m *testMachine) InstanceStatus(context.Context) (status.Status, string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.instStatus == "" {
		return "pending", "", nil
	}
	return m.instStatus, m.instStatusMsg, nil
}

func (m *testMachine) SetModificationStatus(_ context.Context, _ status.Status, message string, _ map[string]interface{}) error {
	m.mu.Lock()
	m.modStatusMsg = message
	m.mu.Unlock()
	return nil
}

func (m *testMachine) ModificationStatus() (status.Status, string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return "", m.modStatusMsg, nil
}

func (m *testMachine) SetStatus(_ context.Context, status status.Status, _ string, _ map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.machineStatus = status
	return nil
}

func (m *testMachine) Status(context.Context) (status.Status, string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.machineStatus == "" {
		return "pending", "", nil
	}
	return m.machineStatus, "", nil
}

func (m *testMachine) ModelAgentVersion(context.Context) (*version.Number, error) {
	if m.agentVersion == version.Zero {
		return &coretesting.FakeVersionNumber, nil
	}
	return &m.agentVersion, nil
}

func (m *testMachine) SetUnprovisioned() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.instance = nil
}

func (m *testMachine) SetInstanceInfo(
	_ context.Context,
	instId instance.Id, _ string, _ string, _ *instance.HardwareCharacteristics, _ []params.NetworkConfig, _ []params.Volume,
	_ map[string]params.VolumeAttachmentInfo, _ []string,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.instance = &testInstance{id: string(instId)}
	return nil
}

func (m *testMachine) SetPassword(_ context.Context, password string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.password = password
	return nil
}

func (m *testMachine) GetPassword() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.password
}

func (m *testMachine) EnsureDead(context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.markForRemoval = true
	return nil
}

type mockToolsFinder struct {
}

func (f mockToolsFinder) FindTools(ctx context.Context, number version.Number, os string, a string) (tools.List, error) {
	if number.Compare(version.MustParse("6.6.6")) == 0 {
		return nil, tools.ErrNoMatches
	}
	v, err := version.ParseBinary(fmt.Sprintf("%s-%s-%s", number, os, arch.HostArch()))
	if err != nil {
		return nil, err
	}
	if a == "" {
		return nil, errors.New("missing arch")
	}
	v.Arch = a
	return tools.List{&tools.Tools{Version: v}}, nil
}

type mockDistributionGroupFinder struct {
	groups map[names.MachineTag][]string
}

func (mock *mockDistributionGroupFinder) DistributionGroupByMachineId(
	ctx context.Context,
	tags ...names.MachineTag,
) ([]apiprovisioner.DistributionGroupResult, error) {
	result := make([]apiprovisioner.DistributionGroupResult, len(tags))
	if len(mock.groups) == 0 {
		for i := range tags {
			result[i] = apiprovisioner.DistributionGroupResult{MachineIds: []string{}}
		}
	} else {
		for i, tag := range tags {
			if dg, ok := mock.groups[tag]; ok {
				result[i] = apiprovisioner.DistributionGroupResult{MachineIds: dg}
			} else {
				result[i] = apiprovisioner.DistributionGroupResult{
					MachineIds: []string{}, Err: &params.Error{Code: params.CodeNotFound, Message: "Fail"}}
			}
		}
	}
	return result, nil
}