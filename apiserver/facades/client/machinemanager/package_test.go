// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package machinemanager

import (
	"testing"

	coretesting "github.com/juju/juju/internal/testing"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package machinemanager -destination package_mock_test.go github.com/juju/juju/apiserver/facades/client/machinemanager Leadership,Authorizer,ControllerBackend,InstanceConfigBackend,Backend,StorageInterface,Pool,Machine,Application,Unit,Charm,CharmhubClient,ControllerConfigService,MachineService,NetworkService,KeyUpdaterService,ModelConfigService
//go:generate go run go.uber.org/mock/mockgen -typed -package machinemanager -destination state_mock_test.go github.com/juju/juju/state StorageAttachment,StorageInstance,Block
//go:generate go run go.uber.org/mock/mockgen -typed -package machinemanager -destination state_storage_mock_test.go github.com/juju/juju/state/binarystorage StorageCloser
//go:generate go run go.uber.org/mock/mockgen -typed -package machinemanager -destination volume_access_mock_test.go github.com/juju/juju/apiserver/common/storagecommon VolumeAccess
//go:generate go run go.uber.org/mock/mockgen -typed -package machinemanager -destination environ_mock_test.go github.com/juju/juju/environs Environ,InstanceTypesFetcher,BootstrapEnviron
//go:generate go run go.uber.org/mock/mockgen -typed -package machinemanager -destination objectstore_mock_test.go github.com/juju/juju/core/objectstore ObjectStore
func TestPackage(t *testing.T) {
	// TODO(wallyworld) - needed until instance config tests converted to gomock
	coretesting.MgoTestPackage(t)
}
