// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/authentication (interfaces: EntityAuthenticator)
//
// Generated by this command:
//
//	mockgen -typed -package stateauthenticator -destination authentication_mock_test.go github.com/juju/juju/apiserver/authentication EntityAuthenticator
//

// Package stateauthenticator is a generated GoMock package.
package stateauthenticator

import (
	context "context"
	reflect "reflect"

	authentication "github.com/juju/juju/apiserver/authentication"
	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockEntityAuthenticator is a mock of EntityAuthenticator interface.
type MockEntityAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockEntityAuthenticatorMockRecorder
}

// MockEntityAuthenticatorMockRecorder is the mock recorder for MockEntityAuthenticator.
type MockEntityAuthenticatorMockRecorder struct {
	mock *MockEntityAuthenticator
}

// NewMockEntityAuthenticator creates a new mock instance.
func NewMockEntityAuthenticator(ctrl *gomock.Controller) *MockEntityAuthenticator {
	mock := &MockEntityAuthenticator{ctrl: ctrl}
	mock.recorder = &MockEntityAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityAuthenticator) EXPECT() *MockEntityAuthenticatorMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockEntityAuthenticator) Authenticate(arg0 context.Context, arg1 authentication.AuthParams) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockEntityAuthenticatorMockRecorder) Authenticate(arg0, arg1 any) *MockEntityAuthenticatorAuthenticateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockEntityAuthenticator)(nil).Authenticate), arg0, arg1)
	return &MockEntityAuthenticatorAuthenticateCall{Call: call}
}

// MockEntityAuthenticatorAuthenticateCall wrap *gomock.Call
type MockEntityAuthenticatorAuthenticateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEntityAuthenticatorAuthenticateCall) Return(arg0 state.Entity, arg1 error) *MockEntityAuthenticatorAuthenticateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEntityAuthenticatorAuthenticateCall) Do(f func(context.Context, authentication.AuthParams) (state.Entity, error)) *MockEntityAuthenticatorAuthenticateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEntityAuthenticatorAuthenticateCall) DoAndReturn(f func(context.Context, authentication.AuthParams) (state.Entity, error)) *MockEntityAuthenticatorAuthenticateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}