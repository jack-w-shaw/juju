// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/cloud/service (interfaces: State)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cloud "github.com/juju/juju/cloud"
	state "github.com/juju/juju/domain/cloud/state"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockState) List(arg0 context.Context, arg1 *state.Filter) ([]cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStateMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockState)(nil).List), arg0, arg1)
}

// Save mocks base method.
func (m *MockState) Save(arg0 context.Context, arg1 cloud.Cloud) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockStateMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockState)(nil).Save), arg0, arg1)
}
