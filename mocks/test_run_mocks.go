// Code generated by MockGen. DO NOT EDIT.
// Source: test_run.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockTestRuns is a mock of TestRuns interface.
type MockTestRuns struct {
	ctrl     *gomock.Controller
	recorder *MockTestRunsMockRecorder
}

// MockTestRunsMockRecorder is the mock recorder for MockTestRuns.
type MockTestRunsMockRecorder struct {
	mock *MockTestRuns
}

// NewMockTestRuns creates a new mock instance.
func NewMockTestRuns(ctrl *gomock.Controller) *MockTestRuns {
	mock := &MockTestRuns{ctrl: ctrl}
	mock.recorder = &MockTestRunsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestRuns) EXPECT() *MockTestRunsMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockTestRuns) Cancel(ctx context.Context, moduleID tfe.RegistryModuleID, testRunID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", ctx, moduleID, testRunID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockTestRunsMockRecorder) Cancel(ctx, moduleID, testRunID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockTestRuns)(nil).Cancel), ctx, moduleID, testRunID)
}

// Create mocks base method.
func (m *MockTestRuns) Create(ctx context.Context, options tfe.TestRunCreateOptions) (*tfe.TestRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, options)
	ret0, _ := ret[0].(*tfe.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTestRunsMockRecorder) Create(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTestRuns)(nil).Create), ctx, options)
}

// ForceCancel mocks base method.
func (m *MockTestRuns) ForceCancel(ctx context.Context, moduleID tfe.RegistryModuleID, testRunID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceCancel", ctx, moduleID, testRunID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceCancel indicates an expected call of ForceCancel.
func (mr *MockTestRunsMockRecorder) ForceCancel(ctx, moduleID, testRunID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceCancel", reflect.TypeOf((*MockTestRuns)(nil).ForceCancel), ctx, moduleID, testRunID)
}

// List mocks base method.
func (m *MockTestRuns) List(ctx context.Context, moduleID tfe.RegistryModuleID, options *tfe.TestRunListOptions) (*tfe.TestRunList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, moduleID, options)
	ret0, _ := ret[0].(*tfe.TestRunList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTestRunsMockRecorder) List(ctx, moduleID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTestRuns)(nil).List), ctx, moduleID, options)
}

// Logs mocks base method.
func (m *MockTestRuns) Logs(ctx context.Context, moduleID tfe.RegistryModuleID, testRunID string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", ctx, moduleID, testRunID)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockTestRunsMockRecorder) Logs(ctx, moduleID, testRunID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockTestRuns)(nil).Logs), ctx, moduleID, testRunID)
}

// Read mocks base method.
func (m *MockTestRuns) Read(ctx context.Context, moduleID tfe.RegistryModuleID, testRunID string) (*tfe.TestRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, moduleID, testRunID)
	ret0, _ := ret[0].(*tfe.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTestRunsMockRecorder) Read(ctx, moduleID, testRunID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTestRuns)(nil).Read), ctx, moduleID, testRunID)
}
