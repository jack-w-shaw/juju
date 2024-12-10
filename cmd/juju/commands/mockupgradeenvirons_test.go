// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/commands (interfaces: UpgradePrecheckEnviron)
//
// Generated by this command:
//
//	mockgen -package commands -destination mockupgradeenvirons_test.go github.com/juju/juju/cmd/juju/commands UpgradePrecheckEnviron
//

// Package commands is a generated GoMock package.
package commands

import (
	reflect "reflect"

	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	environs "github.com/juju/juju/environs"
	config "github.com/juju/juju/environs/config"
	context "github.com/juju/juju/environs/context"
	instances "github.com/juju/juju/environs/instances"
	storage "github.com/juju/juju/storage"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockUpgradePrecheckEnviron is a mock of UpgradePrecheckEnviron interface.
type MockUpgradePrecheckEnviron struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradePrecheckEnvironMockRecorder
}

// MockUpgradePrecheckEnvironMockRecorder is the mock recorder for MockUpgradePrecheckEnviron.
type MockUpgradePrecheckEnvironMockRecorder struct {
	mock *MockUpgradePrecheckEnviron
}

// NewMockUpgradePrecheckEnviron creates a new mock instance.
func NewMockUpgradePrecheckEnviron(ctrl *gomock.Controller) *MockUpgradePrecheckEnviron {
	mock := &MockUpgradePrecheckEnviron{ctrl: ctrl}
	mock.recorder = &MockUpgradePrecheckEnvironMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradePrecheckEnviron) EXPECT() *MockUpgradePrecheckEnvironMockRecorder {
	return m.recorder
}

// AdoptResources mocks base method.
func (m *MockUpgradePrecheckEnviron) AdoptResources(arg0 context.ProviderCallContext, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdoptResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdoptResources indicates an expected call of AdoptResources.
func (mr *MockUpgradePrecheckEnvironMockRecorder) AdoptResources(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdoptResources", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).AdoptResources), arg0, arg1, arg2)
}

// AllInstances mocks base method.
func (m *MockUpgradePrecheckEnviron) AllInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllInstances indicates an expected call of AllInstances.
func (mr *MockUpgradePrecheckEnvironMockRecorder) AllInstances(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllInstances", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).AllInstances), arg0)
}

// AllRunningInstances mocks base method.
func (m *MockUpgradePrecheckEnviron) AllRunningInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRunningInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRunningInstances indicates an expected call of AllRunningInstances.
func (mr *MockUpgradePrecheckEnvironMockRecorder) AllRunningInstances(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRunningInstances", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).AllRunningInstances), arg0)
}

// Bootstrap mocks base method.
func (m *MockUpgradePrecheckEnviron) Bootstrap(arg0 environs.BootstrapContext, arg1 context.ProviderCallContext, arg2 environs.BootstrapParams) (*environs.BootstrapResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*environs.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap.
func (mr *MockUpgradePrecheckEnvironMockRecorder) Bootstrap(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).Bootstrap), arg0, arg1, arg2)
}

// Config mocks base method.
func (m *MockUpgradePrecheckEnviron) Config() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockUpgradePrecheckEnvironMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).Config))
}

// ConstraintsValidator mocks base method.
func (m *MockUpgradePrecheckEnviron) ConstraintsValidator(arg0 context.ProviderCallContext) (constraints.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstraintsValidator", arg0)
	ret0, _ := ret[0].(constraints.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstraintsValidator indicates an expected call of ConstraintsValidator.
func (mr *MockUpgradePrecheckEnvironMockRecorder) ConstraintsValidator(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstraintsValidator", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).ConstraintsValidator), arg0)
}

// ControllerInstances mocks base method.
func (m *MockUpgradePrecheckEnviron) ControllerInstances(arg0 context.ProviderCallContext, arg1 string) ([]instance.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerInstances", arg0, arg1)
	ret0, _ := ret[0].([]instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerInstances indicates an expected call of ControllerInstances.
func (mr *MockUpgradePrecheckEnvironMockRecorder) ControllerInstances(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerInstances", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).ControllerInstances), arg0, arg1)
}

// Create mocks base method.
func (m *MockUpgradePrecheckEnviron) Create(arg0 context.ProviderCallContext, arg1 environs.CreateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUpgradePrecheckEnvironMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).Create), arg0, arg1)
}

// Destroy mocks base method.
func (m *MockUpgradePrecheckEnviron) Destroy(arg0 context.ProviderCallContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockUpgradePrecheckEnvironMockRecorder) Destroy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).Destroy), arg0)
}

// DestroyController mocks base method.
func (m *MockUpgradePrecheckEnviron) DestroyController(arg0 context.ProviderCallContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyController indicates an expected call of DestroyController.
func (mr *MockUpgradePrecheckEnvironMockRecorder) DestroyController(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyController", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).DestroyController), arg0, arg1)
}

// InstanceTypes mocks base method.
func (m *MockUpgradePrecheckEnviron) InstanceTypes(arg0 context.ProviderCallContext, arg1 constraints.Value) (instances.InstanceTypesWithCostMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceTypes", arg0, arg1)
	ret0, _ := ret[0].(instances.InstanceTypesWithCostMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceTypes indicates an expected call of InstanceTypes.
func (mr *MockUpgradePrecheckEnvironMockRecorder) InstanceTypes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceTypes", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).InstanceTypes), arg0, arg1)
}

// Instances mocks base method.
func (m *MockUpgradePrecheckEnviron) Instances(arg0 context.ProviderCallContext, arg1 []instance.Id) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instances", arg0, arg1)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Instances indicates an expected call of Instances.
func (mr *MockUpgradePrecheckEnvironMockRecorder) Instances(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instances", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).Instances), arg0, arg1)
}

// PrecheckInstance mocks base method.
func (m *MockUpgradePrecheckEnviron) PrecheckInstance(arg0 context.ProviderCallContext, arg1 environs.PrecheckInstanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrecheckInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrecheckInstance indicates an expected call of PrecheckInstance.
func (mr *MockUpgradePrecheckEnvironMockRecorder) PrecheckInstance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrecheckInstance", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).PrecheckInstance), arg0, arg1)
}

// PrecheckUpgradeOperations mocks base method.
func (m *MockUpgradePrecheckEnviron) PrecheckUpgradeOperations() []environs.PrecheckJujuUpgradeOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrecheckUpgradeOperations")
	ret0, _ := ret[0].([]environs.PrecheckJujuUpgradeOperation)
	return ret0
}

// PrecheckUpgradeOperations indicates an expected call of PrecheckUpgradeOperations.
func (mr *MockUpgradePrecheckEnvironMockRecorder) PrecheckUpgradeOperations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrecheckUpgradeOperations", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).PrecheckUpgradeOperations))
}

// PrepareForBootstrap mocks base method.
func (m *MockUpgradePrecheckEnviron) PrepareForBootstrap(arg0 environs.BootstrapContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareForBootstrap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareForBootstrap indicates an expected call of PrepareForBootstrap.
func (mr *MockUpgradePrecheckEnvironMockRecorder) PrepareForBootstrap(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForBootstrap", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).PrepareForBootstrap), arg0, arg1)
}

// PreparePrechecker mocks base method.
func (m *MockUpgradePrecheckEnviron) PreparePrechecker() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreparePrechecker")
	ret0, _ := ret[0].(error)
	return ret0
}

// PreparePrechecker indicates an expected call of PreparePrechecker.
func (mr *MockUpgradePrecheckEnvironMockRecorder) PreparePrechecker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreparePrechecker", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).PreparePrechecker))
}

// Provider mocks base method.
func (m *MockUpgradePrecheckEnviron) Provider() environs.EnvironProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(environs.EnvironProvider)
	return ret0
}

// Provider indicates an expected call of Provider.
func (mr *MockUpgradePrecheckEnvironMockRecorder) Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).Provider))
}

// SetConfig mocks base method.
func (m *MockUpgradePrecheckEnviron) SetConfig(arg0 *config.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockUpgradePrecheckEnvironMockRecorder) SetConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).SetConfig), arg0)
}

// StartInstance mocks base method.
func (m *MockUpgradePrecheckEnviron) StartInstance(arg0 context.ProviderCallContext, arg1 environs.StartInstanceParams) (*environs.StartInstanceResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstance", arg0, arg1)
	ret0, _ := ret[0].(*environs.StartInstanceResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartInstance indicates an expected call of StartInstance.
func (mr *MockUpgradePrecheckEnvironMockRecorder) StartInstance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstance", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).StartInstance), arg0, arg1)
}

// StopInstances mocks base method.
func (m *MockUpgradePrecheckEnviron) StopInstances(arg0 context.ProviderCallContext, arg1 ...instance.Id) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopInstances", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopInstances indicates an expected call of StopInstances.
func (mr *MockUpgradePrecheckEnvironMockRecorder) StopInstances(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInstances", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).StopInstances), varargs...)
}

// StorageProvider mocks base method.
func (m *MockUpgradePrecheckEnviron) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockUpgradePrecheckEnvironMockRecorder) StorageProvider(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).StorageProvider), arg0)
}

// StorageProviderTypes mocks base method.
func (m *MockUpgradePrecheckEnviron) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockUpgradePrecheckEnvironMockRecorder) StorageProviderTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockUpgradePrecheckEnviron)(nil).StorageProviderTypes))
}
