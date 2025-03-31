// Code generated by MockGen. DO NOT EDIT.
// Source: ./components/onboarding/internal/adapters/postgres/account/account.postgresql.go
//
// Generated by this command:
//
//	mockgen -source=./components/onboarding/internal/adapters/postgres/account/account.postgresql.go -destination=./components/onboarding/internal/adapters/postgres/account/account.postgresql_mock.go -package=account
//

// Package account is a generated GoMock package.
package account

import (
	context "context"
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/pkg/mmodel"
	http "github.com/LerianStudio/midaz/pkg/net/http"
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

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, acc *mmodel.Account) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, acc)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, acc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, acc)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, organizationID, ledgerID uuid.UUID, portfolioID *uuid.UUID, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organizationID, ledgerID, portfolioID, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, organizationID, ledgerID, portfolioID, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, organizationID, ledgerID, portfolioID, id)
}

// Find mocks base method.
func (m *MockRepository) Find(ctx context.Context, organizationID, ledgerID uuid.UUID, portfolioID *uuid.UUID, id uuid.UUID) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, organizationID, ledgerID, portfolioID, id)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(ctx, organizationID, ledgerID, portfolioID, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), ctx, organizationID, ledgerID, portfolioID, id)
}

// FindAlias mocks base method.
func (m *MockRepository) FindAlias(ctx context.Context, organizationID, ledgerID uuid.UUID, portfolioID *uuid.UUID, alias string) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAlias", ctx, organizationID, ledgerID, portfolioID, alias)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAlias indicates an expected call of FindAlias.
func (mr *MockRepositoryMockRecorder) FindAlias(ctx, organizationID, ledgerID, portfolioID, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAlias", reflect.TypeOf((*MockRepository)(nil).FindAlias), ctx, organizationID, ledgerID, portfolioID, alias)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(ctx context.Context, organizationID, ledgerID uuid.UUID, portfolioID *uuid.UUID, filter http.Pagination) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, organizationID, ledgerID, portfolioID, filter)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(ctx, organizationID, ledgerID, portfolioID, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), ctx, organizationID, ledgerID, portfolioID, filter)
}

// FindByAlias mocks base method.
func (m *MockRepository) FindByAlias(ctx context.Context, organizationID, ledgerID uuid.UUID, alias string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAlias", ctx, organizationID, ledgerID, alias)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAlias indicates an expected call of FindByAlias.
func (mr *MockRepositoryMockRecorder) FindByAlias(ctx, organizationID, ledgerID, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAlias", reflect.TypeOf((*MockRepository)(nil).FindByAlias), ctx, organizationID, ledgerID, alias)
}

// FindWithDeleted mocks base method.
func (m *MockRepository) FindWithDeleted(ctx context.Context, organizationID, ledgerID uuid.UUID, portfolioID *uuid.UUID, id uuid.UUID) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWithDeleted", ctx, organizationID, ledgerID, portfolioID, id)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWithDeleted indicates an expected call of FindWithDeleted.
func (mr *MockRepositoryMockRecorder) FindWithDeleted(ctx, organizationID, ledgerID, portfolioID, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWithDeleted", reflect.TypeOf((*MockRepository)(nil).FindWithDeleted), ctx, organizationID, ledgerID, portfolioID, id)
}

// ListAccountsByAlias mocks base method.
func (m *MockRepository) ListAccountsByAlias(ctx context.Context, organizationID, ledgerID uuid.UUID, aliases []string) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountsByAlias", ctx, organizationID, ledgerID, aliases)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountsByAlias indicates an expected call of ListAccountsByAlias.
func (mr *MockRepositoryMockRecorder) ListAccountsByAlias(ctx, organizationID, ledgerID, aliases any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsByAlias", reflect.TypeOf((*MockRepository)(nil).ListAccountsByAlias), ctx, organizationID, ledgerID, aliases)
}

// ListAccountsByIDs mocks base method.
func (m *MockRepository) ListAccountsByIDs(ctx context.Context, organizationID, ledgerID uuid.UUID, ids []uuid.UUID) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountsByIDs", ctx, organizationID, ledgerID, ids)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountsByIDs indicates an expected call of ListAccountsByIDs.
func (mr *MockRepositoryMockRecorder) ListAccountsByIDs(ctx, organizationID, ledgerID, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsByIDs", reflect.TypeOf((*MockRepository)(nil).ListAccountsByIDs), ctx, organizationID, ledgerID, ids)
}

// ListByAlias mocks base method.
func (m *MockRepository) ListByAlias(ctx context.Context, organizationID, ledgerID, portfolioID uuid.UUID, alias []string) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByAlias", ctx, organizationID, ledgerID, portfolioID, alias)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByAlias indicates an expected call of ListByAlias.
func (mr *MockRepositoryMockRecorder) ListByAlias(ctx, organizationID, ledgerID, portfolioID, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByAlias", reflect.TypeOf((*MockRepository)(nil).ListByAlias), ctx, organizationID, ledgerID, portfolioID, alias)
}

// ListByIDs mocks base method.
func (m *MockRepository) ListByIDs(ctx context.Context, organizationID, ledgerID uuid.UUID, portfolioID *uuid.UUID, ids []uuid.UUID) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDs", ctx, organizationID, ledgerID, portfolioID, ids)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDs indicates an expected call of ListByIDs.
func (mr *MockRepositoryMockRecorder) ListByIDs(ctx, organizationID, ledgerID, portfolioID, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDs", reflect.TypeOf((*MockRepository)(nil).ListByIDs), ctx, organizationID, ledgerID, portfolioID, ids)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, organizationID, ledgerID uuid.UUID, portfolioID *uuid.UUID, id uuid.UUID, acc *mmodel.Account) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, organizationID, ledgerID, portfolioID, id, acc)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, organizationID, ledgerID, portfolioID, id, acc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, organizationID, ledgerID, portfolioID, id, acc)
}
