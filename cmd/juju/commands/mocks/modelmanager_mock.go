// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/commands (interfaces: ModelManagerAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockModelManagerAPI is a mock of ModelManagerAPI interface.
type MockModelManagerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockModelManagerAPIMockRecorder
}

// MockModelManagerAPIMockRecorder is the mock recorder for MockModelManagerAPI.
type MockModelManagerAPIMockRecorder struct {
	mock *MockModelManagerAPI
}

// NewMockModelManagerAPI creates a new mock instance.
func NewMockModelManagerAPI(ctrl *gomock.Controller) *MockModelManagerAPI {
	mock := &MockModelManagerAPI{ctrl: ctrl}
	mock.recorder = &MockModelManagerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelManagerAPI) EXPECT() *MockModelManagerAPIMockRecorder {
	return m.recorder
}

// BestAPIVersion mocks base method.
func (m *MockModelManagerAPI) BestAPIVersion() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestAPIVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// BestAPIVersion indicates an expected call of BestAPIVersion.
func (mr *MockModelManagerAPIMockRecorder) BestAPIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestAPIVersion", reflect.TypeOf((*MockModelManagerAPI)(nil).BestAPIVersion))
}

// Close mocks base method.
func (m *MockModelManagerAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockModelManagerAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockModelManagerAPI)(nil).Close))
}

// ValidateModelUpgrade mocks base method.
func (m *MockModelManagerAPI) ValidateModelUpgrade(arg0 names.ModelTag, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateModelUpgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateModelUpgrade indicates an expected call of ValidateModelUpgrade.
func (mr *MockModelManagerAPIMockRecorder) ValidateModelUpgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateModelUpgrade", reflect.TypeOf((*MockModelManagerAPI)(nil).ValidateModelUpgrade), arg0, arg1)
}
