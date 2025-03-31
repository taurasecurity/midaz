// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/fredamaral/Repos/lerianstudio/midaz/components/onboarding/internal/adapters/postgres/portfolio/portfolio.postgresql.go
//
// Generated by this command:
//
//	mockgen -source=/Users/fredamaral/Repos/lerianstudio/midaz/components/onboarding/internal/adapters/postgres/portfolio/portfolio.postgresql.go -destination=/Users/fredamaral/Repos/lerianstudio/midaz/components/onboarding/internal/adapters/postgres/portfolio/portfolio.postgresql_mock.go -package=portfolio
//

// Package portfolio is a generated GoMock package.
package portfolio

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
func (m *MockRepository) Create(ctx context.Context, portfolio *mmodel.Portfolio) (*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, portfolio)
	ret0, _ := ret[0].(*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, portfolio any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, portfolio)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, organizationID, ledgerID, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organizationID, ledgerID, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, organizationID, ledgerID, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, organizationID, ledgerID, id)
}

// Find mocks base method.
func (m *MockRepository) Find(ctx context.Context, organizationID, ledgerID, id uuid.UUID) (*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, organizationID, ledgerID, id)
	ret0, _ := ret[0].(*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(ctx, organizationID, ledgerID, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), ctx, organizationID, ledgerID, id)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(ctx context.Context, organizationID, ledgerID uuid.UUID, filter http.Pagination) ([]*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, organizationID, ledgerID, filter)
	ret0, _ := ret[0].([]*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(ctx, organizationID, ledgerID, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), ctx, organizationID, ledgerID, filter)
}

// FindByIDEntity mocks base method.
func (m *MockRepository) FindByIDEntity(ctx context.Context, organizationID, ledgerID, entityID uuid.UUID) (*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDEntity", ctx, organizationID, ledgerID, entityID)
	ret0, _ := ret[0].(*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDEntity indicates an expected call of FindByIDEntity.
func (mr *MockRepositoryMockRecorder) FindByIDEntity(ctx, organizationID, ledgerID, entityID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDEntity", reflect.TypeOf((*MockRepository)(nil).FindByIDEntity), ctx, organizationID, ledgerID, entityID)
}

// ListByIDs mocks base method.
func (m *MockRepository) ListByIDs(ctx context.Context, organizationID, ledgerID uuid.UUID, ids []uuid.UUID) ([]*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDs", ctx, organizationID, ledgerID, ids)
	ret0, _ := ret[0].([]*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDs indicates an expected call of ListByIDs.
func (mr *MockRepositoryMockRecorder) ListByIDs(ctx, organizationID, ledgerID, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDs", reflect.TypeOf((*MockRepository)(nil).ListByIDs), ctx, organizationID, ledgerID, ids)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, organizationID, ledgerID, id uuid.UUID, portfolio *mmodel.Portfolio) (*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, organizationID, ledgerID, id, portfolio)
	ret0, _ := ret[0].(*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, organizationID, ledgerID, id, portfolio any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, organizationID, ledgerID, id, portfolio)
}
