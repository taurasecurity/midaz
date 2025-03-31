// Code generated by MockGen. DO NOT EDIT.
// Source: ./components/mdz/internal/domain/repository/auth.go
//
// Generated by this command:
//
//	mockgen -source=./components/mdz/internal/domain/repository/auth.go -destination=./components/mdz/internal/domain/repository/auth_mock.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	model "github.com/LerianStudio/midaz/components/mdz/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
	isgomock struct{}
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// AuthenticateWithCredentials mocks base method.
func (m *MockAuth) AuthenticateWithCredentials(username, password string) (*model.TokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateWithCredentials", username, password)
	ret0, _ := ret[0].(*model.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateWithCredentials indicates an expected call of AuthenticateWithCredentials.
func (mr *MockAuthMockRecorder) AuthenticateWithCredentials(username, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateWithCredentials", reflect.TypeOf((*MockAuth)(nil).AuthenticateWithCredentials), username, password)
}

// ExchangeToken mocks base method.
func (m *MockAuth) ExchangeToken(code string) (*model.TokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeToken", code)
	ret0, _ := ret[0].(*model.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExchangeToken indicates an expected call of ExchangeToken.
func (mr *MockAuthMockRecorder) ExchangeToken(code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeToken", reflect.TypeOf((*MockAuth)(nil).ExchangeToken), code)
}
