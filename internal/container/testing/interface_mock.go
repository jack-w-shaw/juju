// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/container (interfaces: TestLXDManager)

package testing

import (
	reflect "reflect"

	instancecfg "github.com/juju/juju/internal/cloudconfig/instancecfg"
	base "github.com/juju/juju/core/base"
	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	lxdprofile "github.com/juju/juju/core/lxdprofile"
	environs "github.com/juju/juju/environs"
	instances "github.com/juju/juju/environs/instances"
	container "github.com/juju/juju/internal/container"
	gomock "go.uber.org/mock/gomock"
)

// MockTestLXDManager is a mock of TestLXDManager interface.
type MockTestLXDManager struct {
	ctrl     *gomock.Controller
	recorder *MockTestLXDManagerMockRecorder
}

// MockTestLXDManagerMockRecorder is the mock recorder for MockTestLXDManager.
type MockTestLXDManagerMockRecorder struct {
	mock *MockTestLXDManager
}

// NewMockTestLXDManager creates a new mock instance.
func NewMockTestLXDManager(ctrl *gomock.Controller) *MockTestLXDManager {
	mock := &MockTestLXDManager{ctrl: ctrl}
	mock.recorder = &MockTestLXDManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestLXDManager) EXPECT() *MockTestLXDManagerMockRecorder {
	return m.recorder
}

// AssignLXDProfiles mocks base method.
func (m *MockTestLXDManager) AssignLXDProfiles(arg0 string, arg1 []string, arg2 []lxdprofile.ProfilePost) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignLXDProfiles", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignLXDProfiles indicates an expected call of AssignLXDProfiles.
func (mr *MockTestLXDManagerMockRecorder) AssignLXDProfiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignLXDProfiles", reflect.TypeOf((*MockTestLXDManager)(nil).AssignLXDProfiles), arg0, arg1, arg2)
}

// CreateContainer mocks base method.
func (m *MockTestLXDManager) CreateContainer(arg0 *instancecfg.InstanceConfig, arg1 constraints.Value, arg2 base.Base, arg3 *container.NetworkConfig, arg4 *container.StorageConfig, arg5 environs.StatusCallbackFunc) (instances.Instance, *instance.HardwareCharacteristics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(instances.Instance)
	ret1, _ := ret[1].(*instance.HardwareCharacteristics)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockTestLXDManagerMockRecorder) CreateContainer(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockTestLXDManager)(nil).CreateContainer), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DestroyContainer mocks base method.
func (m *MockTestLXDManager) DestroyContainer(arg0 instance.Id) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyContainer indicates an expected call of DestroyContainer.
func (mr *MockTestLXDManagerMockRecorder) DestroyContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyContainer", reflect.TypeOf((*MockTestLXDManager)(nil).DestroyContainer), arg0)
}

// IsInitialized mocks base method.
func (m *MockTestLXDManager) IsInitialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInitialized indicates an expected call of IsInitialized.
func (mr *MockTestLXDManagerMockRecorder) IsInitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitialized", reflect.TypeOf((*MockTestLXDManager)(nil).IsInitialized))
}

// LXDProfileNames mocks base method.
func (m *MockTestLXDManager) LXDProfileNames(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfileNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LXDProfileNames indicates an expected call of LXDProfileNames.
func (mr *MockTestLXDManagerMockRecorder) LXDProfileNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfileNames", reflect.TypeOf((*MockTestLXDManager)(nil).LXDProfileNames), arg0)
}

// ListContainers mocks base method.
func (m *MockTestLXDManager) ListContainers() ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainers")
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockTestLXDManagerMockRecorder) ListContainers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockTestLXDManager)(nil).ListContainers))
}

// MaybeWriteLXDProfile mocks base method.
func (m *MockTestLXDManager) MaybeWriteLXDProfile(arg0 string, arg1 lxdprofile.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybeWriteLXDProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaybeWriteLXDProfile indicates an expected call of MaybeWriteLXDProfile.
func (mr *MockTestLXDManagerMockRecorder) MaybeWriteLXDProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeWriteLXDProfile", reflect.TypeOf((*MockTestLXDManager)(nil).MaybeWriteLXDProfile), arg0, arg1)
}

// Namespace mocks base method.
func (m *MockTestLXDManager) Namespace() instance.Namespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(instance.Namespace)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockTestLXDManagerMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockTestLXDManager)(nil).Namespace))
}
