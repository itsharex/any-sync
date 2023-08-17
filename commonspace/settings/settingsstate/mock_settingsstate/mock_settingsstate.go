// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/settings/settingsstate (interfaces: StateBuilder,ChangeFactory)

// Package mock_settingsstate is a generated GoMock package.
package mock_settingsstate

import (
	reflect "reflect"

	objecttree "github.com/anyproto/any-sync/commonspace/object/tree/objecttree"
	settingsstate "github.com/anyproto/any-sync/commonspace/settings/settingsstate"
	gomock "go.uber.org/mock/gomock"
)

// MockStateBuilder is a mock of StateBuilder interface.
type MockStateBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockStateBuilderMockRecorder
}

// MockStateBuilderMockRecorder is the mock recorder for MockStateBuilder.
type MockStateBuilderMockRecorder struct {
	mock *MockStateBuilder
}

// NewMockStateBuilder creates a new mock instance.
func NewMockStateBuilder(ctrl *gomock.Controller) *MockStateBuilder {
	mock := &MockStateBuilder{ctrl: ctrl}
	mock.recorder = &MockStateBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateBuilder) EXPECT() *MockStateBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockStateBuilder) Build(arg0 objecttree.ReadableObjectTree, arg1 *settingsstate.State) (*settingsstate.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1)
	ret0, _ := ret[0].(*settingsstate.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockStateBuilderMockRecorder) Build(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockStateBuilder)(nil).Build), arg0, arg1)
}

// MockChangeFactory is a mock of ChangeFactory interface.
type MockChangeFactory struct {
	ctrl     *gomock.Controller
	recorder *MockChangeFactoryMockRecorder
}

// MockChangeFactoryMockRecorder is the mock recorder for MockChangeFactory.
type MockChangeFactoryMockRecorder struct {
	mock *MockChangeFactory
}

// NewMockChangeFactory creates a new mock instance.
func NewMockChangeFactory(ctrl *gomock.Controller) *MockChangeFactory {
	mock := &MockChangeFactory{ctrl: ctrl}
	mock.recorder = &MockChangeFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChangeFactory) EXPECT() *MockChangeFactoryMockRecorder {
	return m.recorder
}

// CreateObjectDeleteChange mocks base method.
func (m *MockChangeFactory) CreateObjectDeleteChange(arg0 string, arg1 *settingsstate.State, arg2 bool) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObjectDeleteChange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObjectDeleteChange indicates an expected call of CreateObjectDeleteChange.
func (mr *MockChangeFactoryMockRecorder) CreateObjectDeleteChange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObjectDeleteChange", reflect.TypeOf((*MockChangeFactory)(nil).CreateObjectDeleteChange), arg0, arg1, arg2)
}
