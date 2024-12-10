// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/leadership (interfaces: Pinner)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/leadership_mock.go github.com/juju/juju/core/leadership Pinner
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPinner is a mock of Pinner interface.
type MockPinner struct {
	ctrl     *gomock.Controller
	recorder *MockPinnerMockRecorder
}

// MockPinnerMockRecorder is the mock recorder for MockPinner.
type MockPinnerMockRecorder struct {
	mock *MockPinner
}

// NewMockPinner creates a new mock instance.
func NewMockPinner(ctrl *gomock.Controller) *MockPinner {
	mock := &MockPinner{ctrl: ctrl}
	mock.recorder = &MockPinnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPinner) EXPECT() *MockPinnerMockRecorder {
	return m.recorder
}

// PinLeadership mocks base method.
func (m *MockPinner) PinLeadership(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinLeadership", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PinLeadership indicates an expected call of PinLeadership.
func (mr *MockPinnerMockRecorder) PinLeadership(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinLeadership", reflect.TypeOf((*MockPinner)(nil).PinLeadership), arg0, arg1)
}

// PinnedLeadership mocks base method.
func (m *MockPinner) PinnedLeadership() map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinnedLeadership")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// PinnedLeadership indicates an expected call of PinnedLeadership.
func (mr *MockPinnerMockRecorder) PinnedLeadership() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinnedLeadership", reflect.TypeOf((*MockPinner)(nil).PinnedLeadership))
}

// UnpinLeadership mocks base method.
func (m *MockPinner) UnpinLeadership(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpinLeadership", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpinLeadership indicates an expected call of UnpinLeadership.
func (mr *MockPinnerMockRecorder) UnpinLeadership(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpinLeadership", reflect.TypeOf((*MockPinner)(nil).UnpinLeadership), arg0, arg1)
}
