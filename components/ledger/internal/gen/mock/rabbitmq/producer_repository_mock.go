// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/ledger/internal/domain/rabbitmq (interfaces: ProducerRepository)
//
// Generated by this command:
//
//	mockgen --destination=../../gen/mock/rabbitmq/producer_repository_mock.go --package=mock . ProducerRepository
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProducerRepository is a mock of ProducerRepository interface.
type MockProducerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProducerRepositoryMockRecorder
}

// MockProducerRepositoryMockRecorder is the mock recorder for MockProducerRepository.
type MockProducerRepositoryMockRecorder struct {
	mock *MockProducerRepository
}

// NewMockProducerRepository creates a new mock instance.
func NewMockProducerRepository(ctrl *gomock.Controller) *MockProducerRepository {
	mock := &MockProducerRepository{ctrl: ctrl}
	mock.recorder = &MockProducerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducerRepository) EXPECT() *MockProducerRepositoryMockRecorder {
	return m.recorder
}

// ProducerDefault mocks base method.
func (m *MockProducerRepository) ProducerDefault(arg0 string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProducerDefault", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProducerDefault indicates an expected call of ProducerDefault.
func (mr *MockProducerRepositoryMockRecorder) ProducerDefault(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProducerDefault", reflect.TypeOf((*MockProducerRepository)(nil).ProducerDefault), arg0)
}
