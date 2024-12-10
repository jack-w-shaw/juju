// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/modelupgrader (interfaces: StatePool,State,Model)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/state_mock.go github.com/juju/juju/apiserver/facades/client/modelupgrader StatePool,State,Model
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	modelupgrader "github.com/juju/juju/apiserver/facades/client/modelupgrader"
	controller "github.com/juju/juju/controller"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	replicaset "github.com/juju/replicaset/v2"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockStatePool is a mock of StatePool interface.
type MockStatePool struct {
	ctrl     *gomock.Controller
	recorder *MockStatePoolMockRecorder
}

// MockStatePoolMockRecorder is the mock recorder for MockStatePool.
type MockStatePoolMockRecorder struct {
	mock *MockStatePool
}

// NewMockStatePool creates a new mock instance.
func NewMockStatePool(ctrl *gomock.Controller) *MockStatePool {
	mock := &MockStatePool{ctrl: ctrl}
	mock.recorder = &MockStatePoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatePool) EXPECT() *MockStatePoolMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockStatePool) Get(arg0 string) (modelupgrader.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(modelupgrader.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStatePoolMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStatePool)(nil).Get), arg0)
}

// MongoVersion mocks base method.
func (m *MockStatePool) MongoVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoVersion indicates an expected call of MongoVersion.
func (mr *MockStatePoolMockRecorder) MongoVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoVersion", reflect.TypeOf((*MockStatePool)(nil).MongoVersion))
}

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

// AbortCurrentUpgrade mocks base method.
func (m *MockState) AbortCurrentUpgrade() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortCurrentUpgrade")
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortCurrentUpgrade indicates an expected call of AbortCurrentUpgrade.
func (mr *MockStateMockRecorder) AbortCurrentUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortCurrentUpgrade", reflect.TypeOf((*MockState)(nil).AbortCurrentUpgrade))
}

// AllModelUUIDs mocks base method.
func (m *MockState) AllModelUUIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelUUIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelUUIDs indicates an expected call of AllModelUUIDs.
func (mr *MockStateMockRecorder) AllModelUUIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelUUIDs", reflect.TypeOf((*MockState)(nil).AllModelUUIDs))
}

// ControllerConfig mocks base method.
func (m *MockState) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockStateMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockState)(nil).ControllerConfig))
}

// HasUpgradeSeriesLocks mocks base method.
func (m *MockState) HasUpgradeSeriesLocks() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUpgradeSeriesLocks")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasUpgradeSeriesLocks indicates an expected call of HasUpgradeSeriesLocks.
func (mr *MockStateMockRecorder) HasUpgradeSeriesLocks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUpgradeSeriesLocks", reflect.TypeOf((*MockState)(nil).HasUpgradeSeriesLocks))
}

// MachineCountForSeries mocks base method.
func (m *MockState) MachineCountForSeries(arg0 ...string) (map[string]int, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MachineCountForSeries", varargs...)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineCountForSeries indicates an expected call of MachineCountForSeries.
func (mr *MockStateMockRecorder) MachineCountForSeries(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineCountForSeries", reflect.TypeOf((*MockState)(nil).MachineCountForSeries), arg0...)
}

// Model mocks base method.
func (m *MockState) Model() (modelupgrader.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(modelupgrader.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockStateMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockState)(nil).Model))
}

// MongoCurrentStatus mocks base method.
func (m *MockState) MongoCurrentStatus() (*replicaset.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MongoCurrentStatus")
	ret0, _ := ret[0].(*replicaset.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MongoCurrentStatus indicates an expected call of MongoCurrentStatus.
func (mr *MockStateMockRecorder) MongoCurrentStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MongoCurrentStatus", reflect.TypeOf((*MockState)(nil).MongoCurrentStatus))
}

// Release mocks base method.
func (m *MockState) Release() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Release indicates an expected call of Release.
func (mr *MockStateMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockState)(nil).Release))
}

// SetModelAgentVersion mocks base method.
func (m *MockState) SetModelAgentVersion(arg0 version.Number, arg1 *string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModelAgentVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModelAgentVersion indicates an expected call of SetModelAgentVersion.
func (mr *MockStateMockRecorder) SetModelAgentVersion(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModelAgentVersion", reflect.TypeOf((*MockState)(nil).SetModelAgentVersion), arg0, arg1, arg2)
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// AgentVersion mocks base method.
func (m *MockModel) AgentVersion() (version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentVersion")
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentVersion indicates an expected call of AgentVersion.
func (mr *MockModelMockRecorder) AgentVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentVersion", reflect.TypeOf((*MockModel)(nil).AgentVersion))
}

// IsControllerModel mocks base method.
func (m *MockModel) IsControllerModel() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsControllerModel")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsControllerModel indicates an expected call of IsControllerModel.
func (mr *MockModelMockRecorder) IsControllerModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsControllerModel", reflect.TypeOf((*MockModel)(nil).IsControllerModel))
}

// Life mocks base method.
func (m *MockModel) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockModelMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockModel)(nil).Life))
}

// MigrationMode mocks base method.
func (m *MockModel) MigrationMode() state.MigrationMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrationMode")
	ret0, _ := ret[0].(state.MigrationMode)
	return ret0
}

// MigrationMode indicates an expected call of MigrationMode.
func (mr *MockModelMockRecorder) MigrationMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrationMode", reflect.TypeOf((*MockModel)(nil).MigrationMode))
}

// Name mocks base method.
func (m *MockModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModel)(nil).Name))
}

// Owner mocks base method.
func (m *MockModel) Owner() names.UserTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(names.UserTag)
	return ret0
}

// Owner indicates an expected call of Owner.
func (mr *MockModelMockRecorder) Owner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockModel)(nil).Owner))
}

// Type mocks base method.
func (m *MockModel) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockModelMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockModel)(nil).Type))
}
