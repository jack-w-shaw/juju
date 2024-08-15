// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/access/modelmigration (interfaces: Coordinator,ImportService,ExportService)
//
// Generated by this command:
//
//	mockgen -typed -package modelmigration -destination migrations_mock_test.go github.com/juju/juju/domain/access/modelmigration Coordinator,ImportService,ExportService
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/juju/juju/core/model"
	modelmigration "github.com/juju/juju/core/modelmigration"
	permission "github.com/juju/juju/core/permission"
	user "github.com/juju/juju/core/user"
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

// CreatePermission mocks base method.
func (m *MockImportService) CreatePermission(arg0 context.Context, arg1 permission.UserAccessSpec) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", arg0, arg1)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockImportServiceMockRecorder) CreatePermission(arg0, arg1 any) *MockImportServiceCreatePermissionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockImportService)(nil).CreatePermission), arg0, arg1)
	return &MockImportServiceCreatePermissionCall{Call: call}
}

// MockImportServiceCreatePermissionCall wrap *gomock.Call
type MockImportServiceCreatePermissionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceCreatePermissionCall) Return(arg0 permission.UserAccess, arg1 error) *MockImportServiceCreatePermissionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceCreatePermissionCall) Do(f func(context.Context, permission.UserAccessSpec) (permission.UserAccess, error)) *MockImportServiceCreatePermissionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceCreatePermissionCall) DoAndReturn(f func(context.Context, permission.UserAccessSpec) (permission.UserAccess, error)) *MockImportServiceCreatePermissionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetLastModelLogin mocks base method.
func (m *MockImportService) SetLastModelLogin(arg0 context.Context, arg1 user.Name, arg2 model.UUID, arg3 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLastModelLogin", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLastModelLogin indicates an expected call of SetLastModelLogin.
func (mr *MockImportServiceMockRecorder) SetLastModelLogin(arg0, arg1, arg2, arg3 any) *MockImportServiceSetLastModelLoginCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastModelLogin", reflect.TypeOf((*MockImportService)(nil).SetLastModelLogin), arg0, arg1, arg2, arg3)
	return &MockImportServiceSetLastModelLoginCall{Call: call}
}

// MockImportServiceSetLastModelLoginCall wrap *gomock.Call
type MockImportServiceSetLastModelLoginCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceSetLastModelLoginCall) Return(arg0 error) *MockImportServiceSetLastModelLoginCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceSetLastModelLoginCall) Do(f func(context.Context, user.Name, model.UUID, time.Time) error) *MockImportServiceSetLastModelLoginCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceSetLastModelLoginCall) DoAndReturn(f func(context.Context, user.Name, model.UUID, time.Time) error) *MockImportServiceSetLastModelLoginCall {
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

// LastModelLogin mocks base method.
func (m *MockExportService) LastModelLogin(arg0 context.Context, arg1 user.Name, arg2 model.UUID) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastModelLogin", arg0, arg1, arg2)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastModelLogin indicates an expected call of LastModelLogin.
func (mr *MockExportServiceMockRecorder) LastModelLogin(arg0, arg1, arg2 any) *MockExportServiceLastModelLoginCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastModelLogin", reflect.TypeOf((*MockExportService)(nil).LastModelLogin), arg0, arg1, arg2)
	return &MockExportServiceLastModelLoginCall{Call: call}
}

// MockExportServiceLastModelLoginCall wrap *gomock.Call
type MockExportServiceLastModelLoginCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceLastModelLoginCall) Return(arg0 time.Time, arg1 error) *MockExportServiceLastModelLoginCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceLastModelLoginCall) Do(f func(context.Context, user.Name, model.UUID) (time.Time, error)) *MockExportServiceLastModelLoginCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceLastModelLoginCall) DoAndReturn(f func(context.Context, user.Name, model.UUID) (time.Time, error)) *MockExportServiceLastModelLoginCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadAllUserAccessForTarget mocks base method.
func (m *MockExportService) ReadAllUserAccessForTarget(arg0 context.Context, arg1 permission.ID) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForTarget", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForTarget indicates an expected call of ReadAllUserAccessForTarget.
func (mr *MockExportServiceMockRecorder) ReadAllUserAccessForTarget(arg0, arg1 any) *MockExportServiceReadAllUserAccessForTargetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForTarget", reflect.TypeOf((*MockExportService)(nil).ReadAllUserAccessForTarget), arg0, arg1)
	return &MockExportServiceReadAllUserAccessForTargetCall{Call: call}
}

// MockExportServiceReadAllUserAccessForTargetCall wrap *gomock.Call
type MockExportServiceReadAllUserAccessForTargetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockExportServiceReadAllUserAccessForTargetCall) Return(arg0 []permission.UserAccess, arg1 error) *MockExportServiceReadAllUserAccessForTargetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockExportServiceReadAllUserAccessForTargetCall) Do(f func(context.Context, permission.ID) ([]permission.UserAccess, error)) *MockExportServiceReadAllUserAccessForTargetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockExportServiceReadAllUserAccessForTargetCall) DoAndReturn(f func(context.Context, permission.ID) ([]permission.UserAccess, error)) *MockExportServiceReadAllUserAccessForTargetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
