// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/uniter/runner/context (interfaces: State)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/state_mock.go github.com/juju/juju/worker/uniter/runner/context State
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	application "github.com/juju/juju/core/application"
	network "github.com/juju/juju/core/network"
	params "github.com/juju/juju/rpc/params"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
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

// ActionBegin mocks base method.
func (m *MockState) ActionBegin(arg0 names.ActionTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionBegin", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionBegin indicates an expected call of ActionBegin.
func (mr *MockStateMockRecorder) ActionBegin(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionBegin", reflect.TypeOf((*MockState)(nil).ActionBegin), arg0)
}

// ActionFinish mocks base method.
func (m *MockState) ActionFinish(arg0 names.ActionTag, arg1 string, arg2 map[string]any, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionFinish", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActionFinish indicates an expected call of ActionFinish.
func (mr *MockStateMockRecorder) ActionFinish(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionFinish", reflect.TypeOf((*MockState)(nil).ActionFinish), arg0, arg1, arg2, arg3)
}

// CloudSpec mocks base method.
func (m *MockState) CloudSpec() (*params.CloudSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudSpec")
	ret0, _ := ret[0].(*params.CloudSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudSpec indicates an expected call of CloudSpec.
func (mr *MockStateMockRecorder) CloudSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudSpec", reflect.TypeOf((*MockState)(nil).CloudSpec))
}

// GetPodSpec mocks base method.
func (m *MockState) GetPodSpec(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodSpec", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodSpec indicates an expected call of GetPodSpec.
func (mr *MockStateMockRecorder) GetPodSpec(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodSpec", reflect.TypeOf((*MockState)(nil).GetPodSpec), arg0)
}

// GetRawK8sSpec mocks base method.
func (m *MockState) GetRawK8sSpec(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawK8sSpec", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawK8sSpec indicates an expected call of GetRawK8sSpec.
func (mr *MockStateMockRecorder) GetRawK8sSpec(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawK8sSpec", reflect.TypeOf((*MockState)(nil).GetRawK8sSpec), arg0)
}

// GoalState mocks base method.
func (m *MockState) GoalState() (application.GoalState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoalState")
	ret0, _ := ret[0].(application.GoalState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GoalState indicates an expected call of GoalState.
func (mr *MockStateMockRecorder) GoalState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoalState", reflect.TypeOf((*MockState)(nil).GoalState))
}

// OpenedMachinePortRangesByEndpoint mocks base method.
func (m *MockState) OpenedMachinePortRangesByEndpoint(arg0 names.MachineTag) (map[names.UnitTag]network.GroupedPortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenedMachinePortRangesByEndpoint", arg0)
	ret0, _ := ret[0].(map[names.UnitTag]network.GroupedPortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenedMachinePortRangesByEndpoint indicates an expected call of OpenedMachinePortRangesByEndpoint.
func (mr *MockStateMockRecorder) OpenedMachinePortRangesByEndpoint(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenedMachinePortRangesByEndpoint", reflect.TypeOf((*MockState)(nil).OpenedMachinePortRangesByEndpoint), arg0)
}

// SetUnitWorkloadVersion mocks base method.
func (m *MockState) SetUnitWorkloadVersion(arg0 names.UnitTag, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnitWorkloadVersion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitWorkloadVersion indicates an expected call of SetUnitWorkloadVersion.
func (mr *MockStateMockRecorder) SetUnitWorkloadVersion(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitWorkloadVersion", reflect.TypeOf((*MockState)(nil).SetUnitWorkloadVersion), arg0, arg1)
}

// StorageAttachment mocks base method.
func (m *MockState) StorageAttachment(arg0 names.StorageTag, arg1 names.UnitTag) (params.StorageAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAttachment", arg0, arg1)
	ret0, _ := ret[0].(params.StorageAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAttachment indicates an expected call of StorageAttachment.
func (mr *MockStateMockRecorder) StorageAttachment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAttachment", reflect.TypeOf((*MockState)(nil).StorageAttachment), arg0, arg1)
}

// UnitStorageAttachments mocks base method.
func (m *MockState) UnitStorageAttachments(arg0 names.UnitTag) ([]params.StorageAttachmentId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitStorageAttachments", arg0)
	ret0, _ := ret[0].([]params.StorageAttachmentId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitStorageAttachments indicates an expected call of UnitStorageAttachments.
func (mr *MockStateMockRecorder) UnitStorageAttachments(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitStorageAttachments", reflect.TypeOf((*MockState)(nil).UnitStorageAttachments), arg0)
}

// UnitWorkloadVersion mocks base method.
func (m *MockState) UnitWorkloadVersion(arg0 names.UnitTag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitWorkloadVersion", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitWorkloadVersion indicates an expected call of UnitWorkloadVersion.
func (mr *MockStateMockRecorder) UnitWorkloadVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitWorkloadVersion", reflect.TypeOf((*MockState)(nil).UnitWorkloadVersion), arg0)
}
