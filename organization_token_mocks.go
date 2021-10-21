// Code generated by MockGen. DO NOT EDIT.
// Source: organization_token.go

// Package tfe is a generated GoMock package.
package tfe

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrganizationTokens is a mock of OrganizationTokens interface.
type MockOrganizationTokens struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationTokensMockRecorder
}

// MockOrganizationTokensMockRecorder is the mock recorder for MockOrganizationTokens.
type MockOrganizationTokensMockRecorder struct {
	mock *MockOrganizationTokens
}

// NewMockOrganizationTokens creates a new mock instance.
func NewMockOrganizationTokens(ctrl *gomock.Controller) *MockOrganizationTokens {
	mock := &MockOrganizationTokens{ctrl: ctrl}
	mock.recorder = &MockOrganizationTokensMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationTokens) EXPECT() *MockOrganizationTokensMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationTokens) Delete(ctx context.Context, organization string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organization)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationTokensMockRecorder) Delete(ctx, organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationTokens)(nil).Delete), ctx, organization)
}

// Generate mocks base method.
func (m *MockOrganizationTokens) Generate(ctx context.Context, organization string) (*OrganizationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx, organization)
	ret0, _ := ret[0].(*OrganizationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockOrganizationTokensMockRecorder) Generate(ctx, organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockOrganizationTokens)(nil).Generate), ctx, organization)
}

// Read mocks base method.
func (m *MockOrganizationTokens) Read(ctx context.Context, organization string) (*OrganizationToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, organization)
	ret0, _ := ret[0].(*OrganizationToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockOrganizationTokensMockRecorder) Read(ctx, organization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockOrganizationTokens)(nil).Read), ctx, organization)
}