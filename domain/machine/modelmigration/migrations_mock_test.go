// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/machine/modelmigration (interfaces: Coordinator,ImportService,ExportService)
//
// Generated by this command:
//
//	mockgen -typed -package modelmigration -destination migrations_mock_test.go github.com/juju/juju/domain/machine/modelmigration Coordinator,ImportService,ExportService
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	reflect "reflect"

	instance "github.com/juju/juju/core/instance"
	machine "github.com/juju/juju/core/machine"
	modelmigration "github.com/juju/juju/core/modelmigration"
	gomock "go.uber.org/mock/gomock"
)

// MockCoordinator is a mock of Coordinator interface.
type MockCoordinator struct {
	ctrl     *gomock.Controller
	recorder *MockCoordinatorMockRecorder
}

// MockCoordinatorMockRecorder is the mock recorder for MockCoordinator.
type MockCoordinatorMockRecorder struct {
	mock *MockCoordinator
}

// NewMockCoordinator creates a new mock instance.
func NewMockCoordinator(ctrl *gomock.Controller) *MockCoordinator {
	mock := &MockCoordinator{ctrl: ctrl}
	mock.recorder = &MockCoordinatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoordinator) EXPECT() *MockCoordinatorMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCoordinator) Add(arg0 modelmigration.Operation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockCoordinatorMockRecorder) Add(arg0 any) *MockCoordinatorAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCoordinator)(nil).Add), arg0)
	return &MockCoordinatorAddCall{Call: call}
}

// MockCoordinatorAddCall wrap *gomock.Call
type MockCoordinatorAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCoordinatorAddCall) Return() *MockCoordinatorAddCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCoordinatorAddCall) Do(f func(modelmigration.Operation)) *MockCoordinatorAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCoordinatorAddCall) DoAndReturn(f func(modelmigration.Operation)) *MockCoordinatorAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockImportService is a mock of ImportService interface.
type MockImportService struct {
	ctrl     *gomock.Controller
	recorder *MockImportServiceMockRecorder
}

// MockImportServiceMockRecorder is the mock recorder for MockImportService.
type MockImportServiceMockRecorder struct {
	mock *MockImportService
}

// NewMockImportService creates a new mock instance.
func NewMockImportService(ctrl *gomock.Controller) *MockImportService {
	mock := &MockImportService{ctrl: ctrl}
	mock.recorder = &MockImportServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImportService) EXPECT() *MockImportServiceMockRecorder {
	return m.recorder
}

// CreateMachine mocks base method.
func (m *MockImportService) CreateMachine(arg0 context.Context, arg1 machine.Name) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachine", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMachine indicates an expected call of CreateMachine.
func (mr *MockImportServiceMockRecorder) CreateMachine(arg0, arg1 any) *MockImportServiceCreateMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachine", reflect.TypeOf((*MockImportService)(nil).CreateMachine), arg0, arg1)
	return &MockImportServiceCreateMachineCall{Call: call}
}

// MockImportServiceCreateMachineCall wrap *gomock.Call
type MockImportServiceCreateMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceCreateMachineCall) Return(arg0 string, arg1 error) *MockImportServiceCreateMachineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceCreateMachineCall) Do(f func(context.Context, machine.Name) (string, error)) *MockImportServiceCreateMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceCreateMachineCall) DoAndReturn(f func(context.Context, machine.Name) (string, error)) *MockImportServiceCreateMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetMachineCloudInstance mocks base method.
func (m *MockImportService) SetMachineCloudInstance(arg0 context.Context, arg1 string, arg2 instance.Id, arg3 string, arg4 *instance.HardwareCharacteristics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMachineCloudInstance", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMachineCloudInstance indicates an expected call of SetMachineCloudInstance.
func (mr *MockImportServiceMockRecorder) SetMachineCloudInstance(arg0, arg1, arg2, arg3, arg4 any) *MockImportServiceSetMachineCloudInstanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMachineCloudInstance", reflect.TypeOf((*MockImportService)(nil).SetMachineCloudInstance), arg0, arg1, arg2, arg3, arg4)
	return &MockImportServiceSetMachineCloudInstanceCall{Call: call}
}

// MockImportServiceSetMachineCloudInstanceCall wrap *gomock.Call
type MockImportServiceSetMachineCloudInstanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceSetMachineCloudInstanceCall) Return(arg0 error) *MockImportServiceSetMachineCloudInstanceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceSetMachineCloudInstanceCall) Do(f func(context.Context, string, instance.Id, string, *instance.HardwareCharacteristics) error) *MockImportServiceSetMachineCloudInstanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceSetMachineCloudInstanceCall) DoAndReturn(f func(context.Context, string, instance.Id, string, *instance.HardwareCharacteristics) error) *MockImportServiceSetMachineCloudInstanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockExportService is a mock of ExportService interface.
type MockExportService struct {
	ctrl     *gomock.Controller
	recorder *MockExportServiceMockRecorder
}

// MockExportServiceMockRecorder is the mock recorder for MockExportService.
type MockExportServiceMockRecorder struct {
	mock *MockExportService
}

// NewMockExportService creates a new mock instance.
func NewMockExportService(ctrl *gomock.Controller) *MockExportService {
	mock := &MockExportService{ctrl: ctrl}
	mock.recorder = &MockExportServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExportService) EXPECT() *MockExportServiceMockRecorder {
	return m.recorder
}

// AllMachineNames mocks base method.
func (m *MockExportService) AllMachineNames(arg0 context.Context) ([]machine.Name, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMachineNames", arg0)
	ret0, _ := ret[0].([]machine.Name)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMachineNames indicates an expected call of AllMachineNames.
func (mr *MockExportServiceMockRecorder) AllMachineNames(arg0 any) *MockExportServiceAllMachineNamesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMachineNames", reflect.TypeOf((*MockExportService)(nil).AllMachineNames), arg0)
	return &MockExportServiceAllMachineNamesCall{Call: call}
}

// MockExportServiceAllMachineNamesCall wrap *gomock.Call
type MockExportServiceAllMachineNamesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceAllMachineNamesCall) Return(arg0 []machine.Name, arg1 error) *MockExportServiceAllMachineNamesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceAllMachineNamesCall) Do(f func(context.Context) ([]machine.Name, error)) *MockExportServiceAllMachineNamesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceAllMachineNamesCall) DoAndReturn(f func(context.Context) ([]machine.Name, error)) *MockExportServiceAllMachineNamesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetMachineUUID mocks base method.
func (m *MockExportService) GetMachineUUID(arg0 context.Context, arg1 machine.Name) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineUUID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineUUID indicates an expected call of GetMachineUUID.
func (mr *MockExportServiceMockRecorder) GetMachineUUID(arg0, arg1 any) *MockExportServiceGetMachineUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineUUID", reflect.TypeOf((*MockExportService)(nil).GetMachineUUID), arg0, arg1)
	return &MockExportServiceGetMachineUUIDCall{Call: call}
}

// MockExportServiceGetMachineUUIDCall wrap *gomock.Call
type MockExportServiceGetMachineUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceGetMachineUUIDCall) Return(arg0 string, arg1 error) *MockExportServiceGetMachineUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceGetMachineUUIDCall) Do(f func(context.Context, machine.Name) (string, error)) *MockExportServiceGetMachineUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceGetMachineUUIDCall) DoAndReturn(f func(context.Context, machine.Name) (string, error)) *MockExportServiceGetMachineUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HardwareCharacteristics mocks base method.
func (m *MockExportService) HardwareCharacteristics(arg0 context.Context, arg1 string) (*instance.HardwareCharacteristics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardwareCharacteristics", arg0, arg1)
	ret0, _ := ret[0].(*instance.HardwareCharacteristics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardwareCharacteristics indicates an expected call of HardwareCharacteristics.
func (mr *MockExportServiceMockRecorder) HardwareCharacteristics(arg0, arg1 any) *MockExportServiceHardwareCharacteristicsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareCharacteristics", reflect.TypeOf((*MockExportService)(nil).HardwareCharacteristics), arg0, arg1)
	return &MockExportServiceHardwareCharacteristicsCall{Call: call}
}

// MockExportServiceHardwareCharacteristicsCall wrap *gomock.Call
type MockExportServiceHardwareCharacteristicsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceHardwareCharacteristicsCall) Return(arg0 *instance.HardwareCharacteristics, arg1 error) *MockExportServiceHardwareCharacteristicsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceHardwareCharacteristicsCall) Do(f func(context.Context, string) (*instance.HardwareCharacteristics, error)) *MockExportServiceHardwareCharacteristicsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceHardwareCharacteristicsCall) DoAndReturn(f func(context.Context, string) (*instance.HardwareCharacteristics, error)) *MockExportServiceHardwareCharacteristicsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InstanceID mocks base method.
func (m *MockExportService) InstanceID(arg0 context.Context, arg1 string) (instance.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceID", arg0, arg1)
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceID indicates an expected call of InstanceID.
func (mr *MockExportServiceMockRecorder) InstanceID(arg0, arg1 any) *MockExportServiceInstanceIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockExportService)(nil).InstanceID), arg0, arg1)
	return &MockExportServiceInstanceIDCall{Call: call}
}

// MockExportServiceInstanceIDCall wrap *gomock.Call
type MockExportServiceInstanceIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceInstanceIDCall) Return(arg0 instance.Id, arg1 error) *MockExportServiceInstanceIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceInstanceIDCall) Do(f func(context.Context, string) (instance.Id, error)) *MockExportServiceInstanceIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceInstanceIDCall) DoAndReturn(f func(context.Context, string) (instance.Id, error)) *MockExportServiceInstanceIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}