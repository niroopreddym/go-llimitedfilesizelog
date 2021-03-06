// Code generated by MockGen. DO NOT EDIT.
// Source: Ilogger.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	enums "github.com/niroopreddym/go-llimitedfilesizelog/enums"
)

// MockLoggerIface is a mock of LoggerIface interface.
type MockLoggerIface struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerIfaceMockRecorder
}

// MockLoggerIfaceMockRecorder is the mock recorder for MockLoggerIface.
type MockLoggerIfaceMockRecorder struct {
	mock *MockLoggerIface
}

// NewMockLoggerIface creates a new mock instance.
func NewMockLoggerIface(ctrl *gomock.Controller) *MockLoggerIface {
	mock := &MockLoggerIface{ctrl: ctrl}
	mock.recorder = &MockLoggerIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggerIface) EXPECT() *MockLoggerIfaceMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockLoggerIface) Log(logLevel enums.VerbosityLevel, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log", logLevel, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockLoggerIfaceMockRecorder) Log(logLevel, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLoggerIface)(nil).Log), logLevel, message)
}

// SetLogLevel mocks base method.
func (m *MockLoggerIface) SetLogLevel(logLevel enums.VerbosityLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogLevel", logLevel)
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockLoggerIfaceMockRecorder) SetLogLevel(logLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockLoggerIface)(nil).SetLogLevel), logLevel)
}
