// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/charm/services (interfaces: ModelConfigService,ApplicationService)
//
// Generated by this command:
//
//	mockgen -typed -package services -destination service_mock_test.go github.com/juju/juju/internal/charm/services ModelConfigService,ApplicationService
//

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/juju/core/charm"
	charm0 "github.com/juju/juju/domain/application/charm"
	config "github.com/juju/juju/environs/config"
	charm1 "github.com/juju/juju/internal/charm"
	gomock "go.uber.org/mock/gomock"
)

// MockModelConfigService is a mock of ModelConfigService interface.
type MockModelConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigServiceMockRecorder
}

// MockModelConfigServiceMockRecorder is the mock recorder for MockModelConfigService.
type MockModelConfigServiceMockRecorder struct {
	mock *MockModelConfigService
}

// NewMockModelConfigService creates a new mock instance.
func NewMockModelConfigService(ctrl *gomock.Controller) *MockModelConfigService {
	mock := &MockModelConfigService{ctrl: ctrl}
	mock.recorder = &MockModelConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigService) EXPECT() *MockModelConfigServiceMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockModelConfigService) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelConfigServiceMockRecorder) ModelConfig(arg0 any) *MockModelConfigServiceModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModelConfigService)(nil).ModelConfig), arg0)
	return &MockModelConfigServiceModelConfigCall{Call: call}
}

// MockModelConfigServiceModelConfigCall wrap *gomock.Call
type MockModelConfigServiceModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceModelConfigCall) Return(arg0 *config.Config, arg1 error) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceModelConfigCall) Do(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceModelConfigCall) DoAndReturn(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetCharm mocks base method.
func (m *MockApplicationService) GetCharm(arg0 context.Context, arg1 charm.ID) (charm1.Charm, charm0.CharmLocator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharm", arg0, arg1)
	ret0, _ := ret[0].(charm1.Charm)
	ret1, _ := ret[1].(charm0.CharmLocator)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetCharm indicates an expected call of GetCharm.
func (mr *MockApplicationServiceMockRecorder) GetCharm(arg0, arg1 any) *MockApplicationServiceGetCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharm", reflect.TypeOf((*MockApplicationService)(nil).GetCharm), arg0, arg1)
	return &MockApplicationServiceGetCharmCall{Call: call}
}

// MockApplicationServiceGetCharmCall wrap *gomock.Call
type MockApplicationServiceGetCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmCall) Return(arg0 charm1.Charm, arg1 charm0.CharmLocator, arg2 bool, arg3 error) *MockApplicationServiceGetCharmCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmCall) Do(f func(context.Context, charm.ID) (charm1.Charm, charm0.CharmLocator, bool, error)) *MockApplicationServiceGetCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmCall) DoAndReturn(f func(context.Context, charm.ID) (charm1.Charm, charm0.CharmLocator, bool, error)) *MockApplicationServiceGetCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmID mocks base method.
func (m *MockApplicationService) GetCharmID(arg0 context.Context, arg1 charm0.GetCharmArgs) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmID", arg0, arg1)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmID indicates an expected call of GetCharmID.
func (mr *MockApplicationServiceMockRecorder) GetCharmID(arg0, arg1 any) *MockApplicationServiceGetCharmIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmID", reflect.TypeOf((*MockApplicationService)(nil).GetCharmID), arg0, arg1)
	return &MockApplicationServiceGetCharmIDCall{Call: call}
}

// MockApplicationServiceGetCharmIDCall wrap *gomock.Call
type MockApplicationServiceGetCharmIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmIDCall) Return(arg0 charm.ID, arg1 error) *MockApplicationServiceGetCharmIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmIDCall) Do(f func(context.Context, charm0.GetCharmArgs) (charm.ID, error)) *MockApplicationServiceGetCharmIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmIDCall) DoAndReturn(f func(context.Context, charm0.GetCharmArgs) (charm.ID, error)) *MockApplicationServiceGetCharmIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
