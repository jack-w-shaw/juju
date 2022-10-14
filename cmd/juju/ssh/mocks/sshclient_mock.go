// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/ssh (interfaces: SSHClientAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cloudspec "github.com/juju/juju/environs/cloudspec"
)

// MockSSHClientAPI is a mock of SSHClientAPI interface.
type MockSSHClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSSHClientAPIMockRecorder
}

// MockSSHClientAPIMockRecorder is the mock recorder for MockSSHClientAPI.
type MockSSHClientAPIMockRecorder struct {
	mock *MockSSHClientAPI
}

// NewMockSSHClientAPI creates a new mock instance.
func NewMockSSHClientAPI(ctrl *gomock.Controller) *MockSSHClientAPI {
	mock := &MockSSHClientAPI{ctrl: ctrl}
	mock.recorder = &MockSSHClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHClientAPI) EXPECT() *MockSSHClientAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSSHClientAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSSHClientAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSSHClientAPI)(nil).Close))
}

// ModelCredentialForSSH mocks base method.
func (m *MockSSHClientAPI) ModelCredentialForSSH() (cloudspec.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelCredentialForSSH")
	ret0, _ := ret[0].(cloudspec.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelCredentialForSSH indicates an expected call of ModelCredentialForSSH.
func (mr *MockSSHClientAPIMockRecorder) ModelCredentialForSSH() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelCredentialForSSH", reflect.TypeOf((*MockSSHClientAPI)(nil).ModelCredentialForSSH))
}