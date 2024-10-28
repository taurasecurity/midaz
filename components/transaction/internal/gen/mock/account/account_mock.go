// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/transaction/internal/domain/account (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen --destination=../../gen/mock/account/account_mock.go --package=mock . Repository
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	account "github.com/LerianStudio/midaz/common/mgrpc/account"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetAccountsByAlias mocks base method.
func (m *MockRepository) GetAccountsByAlias(ctx context.Context, token string, organizationID, ledgerID uuid.UUID, aliases []string) (*account.AccountsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsByAlias", ctx, token, organizationID, ledgerID, aliases)
	ret0, _ := ret[0].(*account.AccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsByAlias indicates an expected call of GetAccountsByAlias.
func (mr *MockRepositoryMockRecorder) GetAccountsByAlias(ctx, token, organizationID, ledgerID, aliases any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsByAlias", reflect.TypeOf((*MockRepository)(nil).GetAccountsByAlias), ctx, token, organizationID, ledgerID, aliases)
}

// GetAccountsByIds mocks base method.
func (m *MockRepository) GetAccountsByIds(ctx context.Context, token string, organizationID, ledgerID uuid.UUID, ids []string) (*account.AccountsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsByIds", ctx, token, organizationID, ledgerID, ids)
	ret0, _ := ret[0].(*account.AccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsByIds indicates an expected call of GetAccountsByIds.
func (mr *MockRepositoryMockRecorder) GetAccountsByIds(ctx, token, organizationID, ledgerID, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsByIds", reflect.TypeOf((*MockRepository)(nil).GetAccountsByIds), ctx, token, organizationID, ledgerID, ids)
}

// UpdateAccounts mocks base method.
func (m *MockRepository) UpdateAccounts(ctx context.Context, token string, accounts []*account.Account) (*account.AccountsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccounts", ctx, token, accounts)
	ret0, _ := ret[0].(*account.AccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccounts indicates an expected call of UpdateAccounts.
func (mr *MockRepositoryMockRecorder) UpdateAccounts(ctx, token, accounts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccounts", reflect.TypeOf((*MockRepository)(nil).UpdateAccounts), ctx, token, accounts)
}
