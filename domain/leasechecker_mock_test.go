// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain (interfaces: LeaseCheckerWaiter)
//
// Generated by this command:
//
//	mockgen -typed -package domain -destination leasechecker_mock_test.go github.com/juju/juju/domain LeaseCheckerWaiter
//

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	lease "github.com/juju/juju/core/lease"
	gomock "go.uber.org/mock/gomock"
)

// MockLeaseCheckerWaiter is a mock of LeaseCheckerWaiter interface.
type MockLeaseCheckerWaiter struct {
	ctrl     *gomock.Controller
	recorder *MockLeaseCheckerWaiterMockRecorder
}

// MockLeaseCheckerWaiterMockRecorder is the mock recorder for MockLeaseCheckerWaiter.
type MockLeaseCheckerWaiterMockRecorder struct {
	mock *MockLeaseCheckerWaiter
}

// NewMockLeaseCheckerWaiter creates a new mock instance.
func NewMockLeaseCheckerWaiter(ctrl *gomock.Controller) *MockLeaseCheckerWaiter {
	mock := &MockLeaseCheckerWaiter{ctrl: ctrl}
	mock.recorder = &MockLeaseCheckerWaiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaseCheckerWaiter) EXPECT() *MockLeaseCheckerWaiterMockRecorder {
	return m.recorder
}

// Token mocks base method.
func (m *MockLeaseCheckerWaiter) Token(arg0, arg1 string) lease.Token {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", arg0, arg1)
	ret0, _ := ret[0].(lease.Token)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockLeaseCheckerWaiterMockRecorder) Token(arg0, arg1 any) *MockLeaseCheckerWaiterTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockLeaseCheckerWaiter)(nil).Token), arg0, arg1)
	return &MockLeaseCheckerWaiterTokenCall{Call: call}
}

// MockLeaseCheckerWaiterTokenCall wrap *gomock.Call
type MockLeaseCheckerWaiterTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLeaseCheckerWaiterTokenCall) Return(arg0 lease.Token) *MockLeaseCheckerWaiterTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLeaseCheckerWaiterTokenCall) Do(f func(string, string) lease.Token) *MockLeaseCheckerWaiterTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLeaseCheckerWaiterTokenCall) DoAndReturn(f func(string, string) lease.Token) *MockLeaseCheckerWaiterTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WaitUntilExpired mocks base method.
func (m *MockLeaseCheckerWaiter) WaitUntilExpired(arg0 context.Context, arg1 string, arg2 chan<- struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilExpired", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilExpired indicates an expected call of WaitUntilExpired.
func (mr *MockLeaseCheckerWaiterMockRecorder) WaitUntilExpired(arg0, arg1, arg2 any) *MockLeaseCheckerWaiterWaitUntilExpiredCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilExpired", reflect.TypeOf((*MockLeaseCheckerWaiter)(nil).WaitUntilExpired), arg0, arg1, arg2)
	return &MockLeaseCheckerWaiterWaitUntilExpiredCall{Call: call}
}

// MockLeaseCheckerWaiterWaitUntilExpiredCall wrap *gomock.Call
type MockLeaseCheckerWaiterWaitUntilExpiredCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLeaseCheckerWaiterWaitUntilExpiredCall) Return(arg0 error) *MockLeaseCheckerWaiterWaitUntilExpiredCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLeaseCheckerWaiterWaitUntilExpiredCall) Do(f func(context.Context, string, chan<- struct{}) error) *MockLeaseCheckerWaiterWaitUntilExpiredCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLeaseCheckerWaiterWaitUntilExpiredCall) DoAndReturn(f func(context.Context, string, chan<- struct{}) error) *MockLeaseCheckerWaiterWaitUntilExpiredCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
