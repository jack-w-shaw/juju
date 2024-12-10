// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/machinemanager (interfaces: Leadership)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/leadership_mock.go github.com/juju/juju/apiserver/facades/client/machinemanager Leadership
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	params "github.com/juju/juju/rpc/params"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockLeadership is a mock of Leadership interface.
type MockLeadership struct {
	ctrl     *gomock.Controller
	recorder *MockLeadershipMockRecorder
}

// MockLeadershipMockRecorder is the mock recorder for MockLeadership.
type MockLeadershipMockRecorder struct {
	mock *MockLeadership
}

// NewMockLeadership creates a new mock instance.
func NewMockLeadership(ctrl *gomock.Controller) *MockLeadership {
	mock := &MockLeadership{ctrl: ctrl}
	mock.recorder = &MockLeadershipMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeadership) EXPECT() *MockLeadershipMockRecorder {
	return m.recorder
}

// GetMachineApplicationNames mocks base method.
func (m *MockLeadership) GetMachineApplicationNames(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineApplicationNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineApplicationNames indicates an expected call of GetMachineApplicationNames.
func (mr *MockLeadershipMockRecorder) GetMachineApplicationNames(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineApplicationNames", reflect.TypeOf((*MockLeadership)(nil).GetMachineApplicationNames), arg0)
}

// UnpinApplicationLeadersByName mocks base method.
func (m *MockLeadership) UnpinApplicationLeadersByName(arg0 names.Tag, arg1 []string) (params.PinApplicationsResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpinApplicationLeadersByName", arg0, arg1)
	ret0, _ := ret[0].(params.PinApplicationsResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpinApplicationLeadersByName indicates an expected call of UnpinApplicationLeadersByName.
func (mr *MockLeadershipMockRecorder) UnpinApplicationLeadersByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpinApplicationLeadersByName", reflect.TypeOf((*MockLeadership)(nil).UnpinApplicationLeadersByName), arg0, arg1)
}
