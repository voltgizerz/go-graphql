// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/pokemon.service.go

// Package mock_service is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/go-graphql/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPokemonServiceInterface is a mock of PokemonServiceInterface interface.
type MockPokemonServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPokemonServiceInterfaceMockRecorder
}

// MockPokemonServiceInterfaceMockRecorder is the mock recorder for MockPokemonServiceInterface.
type MockPokemonServiceInterfaceMockRecorder struct {
	mock *MockPokemonServiceInterface
}

// NewMockPokemonServiceInterface creates a new mock instance.
func NewMockPokemonServiceInterface(ctrl *gomock.Controller) *MockPokemonServiceInterface {
	mock := &MockPokemonServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPokemonServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPokemonServiceInterface) EXPECT() *MockPokemonServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPokemonServiceInterface) Create(ctx context.Context, input models.CreatePokemonInput) (*models.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*models.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPokemonServiceInterfaceMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPokemonServiceInterface)(nil).Create), ctx, input)
}

// Delete mocks base method.
func (m *MockPokemonServiceInterface) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPokemonServiceInterfaceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPokemonServiceInterface)(nil).Delete), ctx, id)
}

// FetchAll mocks base method.
func (m *MockPokemonServiceInterface) FetchAll(ctx context.Context, limit, offset int) ([]*models.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAll", ctx, limit, offset)
	ret0, _ := ret[0].([]*models.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAll indicates an expected call of FetchAll.
func (mr *MockPokemonServiceInterfaceMockRecorder) FetchAll(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAll", reflect.TypeOf((*MockPokemonServiceInterface)(nil).FetchAll), ctx, limit, offset)
}

// FetchOne mocks base method.
func (m *MockPokemonServiceInterface) FetchOne(ctx context.Context, id int) (*models.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOne", ctx, id)
	ret0, _ := ret[0].(*models.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOne indicates an expected call of FetchOne.
func (mr *MockPokemonServiceInterfaceMockRecorder) FetchOne(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOne", reflect.TypeOf((*MockPokemonServiceInterface)(nil).FetchOne), ctx, id)
}
