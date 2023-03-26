// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ngenohkevin/flock_manager/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateBreed mocks base method.
func (m *MockStore) CreateBreed(arg0 context.Context, arg1 string) (db.Breed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBreed", arg0, arg1)
	ret0, _ := ret[0].(db.Breed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBreed indicates an expected call of CreateBreed.
func (mr *MockStoreMockRecorder) CreateBreed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBreed", reflect.TypeOf((*MockStore)(nil).CreateBreed), arg0, arg1)
}

// CreateHatchery mocks base method.
func (m *MockStore) CreateHatchery(arg0 context.Context, arg1 db.CreateHatcheryParams) (db.Hatchery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHatchery", arg0, arg1)
	ret0, _ := ret[0].(db.Hatchery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHatchery indicates an expected call of CreateHatchery.
func (mr *MockStoreMockRecorder) CreateHatchery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHatchery", reflect.TypeOf((*MockStore)(nil).CreateHatchery), arg0, arg1)
}

// CreatePremises mocks base method.
func (m *MockStore) CreatePremises(arg0 context.Context, arg1 db.CreatePremisesParams) (db.Premise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePremises", arg0, arg1)
	ret0, _ := ret[0].(db.Premise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePremises indicates an expected call of CreatePremises.
func (mr *MockStoreMockRecorder) CreatePremises(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePremises", reflect.TypeOf((*MockStore)(nil).CreatePremises), arg0, arg1)
}

// CreateProduction mocks base method.
func (m *MockStore) CreateProduction(arg0 context.Context, arg1 db.CreateProductionParams) (db.Production, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduction", arg0, arg1)
	ret0, _ := ret[0].(db.Production)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduction indicates an expected call of CreateProduction.
func (mr *MockStoreMockRecorder) CreateProduction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduction", reflect.TypeOf((*MockStore)(nil).CreateProduction), arg0, arg1)
}

// DeleteBreed mocks base method.
func (m *MockStore) DeleteBreed(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBreed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBreed indicates an expected call of DeleteBreed.
func (mr *MockStoreMockRecorder) DeleteBreed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBreed", reflect.TypeOf((*MockStore)(nil).DeleteBreed), arg0, arg1)
}

// DeleteHatchery mocks base method.
func (m *MockStore) DeleteHatchery(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHatchery", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHatchery indicates an expected call of DeleteHatchery.
func (mr *MockStoreMockRecorder) DeleteHatchery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHatchery", reflect.TypeOf((*MockStore)(nil).DeleteHatchery), arg0, arg1)
}

// DeletePremises mocks base method.
func (m *MockStore) DeletePremises(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePremises", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePremises indicates an expected call of DeletePremises.
func (mr *MockStoreMockRecorder) DeletePremises(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePremises", reflect.TypeOf((*MockStore)(nil).DeletePremises), arg0, arg1)
}

// DeleteProduction mocks base method.
func (m *MockStore) DeleteProduction(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduction indicates an expected call of DeleteProduction.
func (mr *MockStoreMockRecorder) DeleteProduction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduction", reflect.TypeOf((*MockStore)(nil).DeleteProduction), arg0, arg1)
}

// GetBreed mocks base method.
func (m *MockStore) GetBreed(arg0 context.Context, arg1 int64) (db.Breed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBreed", arg0, arg1)
	ret0, _ := ret[0].(db.Breed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBreed indicates an expected call of GetBreed.
func (mr *MockStoreMockRecorder) GetBreed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBreed", reflect.TypeOf((*MockStore)(nil).GetBreed), arg0, arg1)
}

// GetHatchery mocks base method.
func (m *MockStore) GetHatchery(arg0 context.Context, arg1 int64) (db.Hatchery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHatchery", arg0, arg1)
	ret0, _ := ret[0].(db.Hatchery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHatchery indicates an expected call of GetHatchery.
func (mr *MockStoreMockRecorder) GetHatchery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHatchery", reflect.TypeOf((*MockStore)(nil).GetHatchery), arg0, arg1)
}

// GetPremises mocks base method.
func (m *MockStore) GetPremises(arg0 context.Context, arg1 int64) (db.Premise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPremises", arg0, arg1)
	ret0, _ := ret[0].(db.Premise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPremises indicates an expected call of GetPremises.
func (mr *MockStoreMockRecorder) GetPremises(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPremises", reflect.TypeOf((*MockStore)(nil).GetPremises), arg0, arg1)
}

// GetProduction mocks base method.
func (m *MockStore) GetProduction(arg0 context.Context, arg1 int64) (db.Production, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduction", arg0, arg1)
	ret0, _ := ret[0].(db.Production)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduction indicates an expected call of GetProduction.
func (mr *MockStoreMockRecorder) GetProduction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduction", reflect.TypeOf((*MockStore)(nil).GetProduction), arg0, arg1)
}

// ListBreeds mocks base method.
func (m *MockStore) ListBreeds(arg0 context.Context, arg1 db.ListBreedsParams) ([]db.Breed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBreeds", arg0, arg1)
	ret0, _ := ret[0].([]db.Breed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBreeds indicates an expected call of ListBreeds.
func (mr *MockStoreMockRecorder) ListBreeds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBreeds", reflect.TypeOf((*MockStore)(nil).ListBreeds), arg0, arg1)
}

// ListHatchery mocks base method.
func (m *MockStore) ListHatchery(arg0 context.Context, arg1 db.ListHatcheryParams) ([]db.Hatchery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHatchery", arg0, arg1)
	ret0, _ := ret[0].([]db.Hatchery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHatchery indicates an expected call of ListHatchery.
func (mr *MockStoreMockRecorder) ListHatchery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHatchery", reflect.TypeOf((*MockStore)(nil).ListHatchery), arg0, arg1)
}

// ListPremises mocks base method.
func (m *MockStore) ListPremises(arg0 context.Context, arg1 db.ListPremisesParams) ([]db.Premise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPremises", arg0, arg1)
	ret0, _ := ret[0].([]db.Premise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPremises indicates an expected call of ListPremises.
func (mr *MockStoreMockRecorder) ListPremises(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPremises", reflect.TypeOf((*MockStore)(nil).ListPremises), arg0, arg1)
}

// ListProduction mocks base method.
func (m *MockStore) ListProduction(arg0 context.Context, arg1 db.ListProductionParams) ([]db.Production, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProduction", arg0, arg1)
	ret0, _ := ret[0].([]db.Production)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProduction indicates an expected call of ListProduction.
func (mr *MockStoreMockRecorder) ListProduction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProduction", reflect.TypeOf((*MockStore)(nil).ListProduction), arg0, arg1)
}

// UpdateBreed mocks base method.
func (m *MockStore) UpdateBreed(arg0 context.Context, arg1 db.UpdateBreedParams) (db.Breed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBreed", arg0, arg1)
	ret0, _ := ret[0].(db.Breed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBreed indicates an expected call of UpdateBreed.
func (mr *MockStoreMockRecorder) UpdateBreed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBreed", reflect.TypeOf((*MockStore)(nil).UpdateBreed), arg0, arg1)
}

// UpdateHatchery mocks base method.
func (m *MockStore) UpdateHatchery(arg0 context.Context, arg1 db.UpdateHatcheryParams) (db.Hatchery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHatchery", arg0, arg1)
	ret0, _ := ret[0].(db.Hatchery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHatchery indicates an expected call of UpdateHatchery.
func (mr *MockStoreMockRecorder) UpdateHatchery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHatchery", reflect.TypeOf((*MockStore)(nil).UpdateHatchery), arg0, arg1)
}

// UpdatePremises mocks base method.
func (m *MockStore) UpdatePremises(arg0 context.Context, arg1 db.UpdatePremisesParams) (db.Premise, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePremises", arg0, arg1)
	ret0, _ := ret[0].(db.Premise)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePremises indicates an expected call of UpdatePremises.
func (mr *MockStoreMockRecorder) UpdatePremises(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePremises", reflect.TypeOf((*MockStore)(nil).UpdatePremises), arg0, arg1)
}

// UpdateProduction mocks base method.
func (m *MockStore) UpdateProduction(arg0 context.Context, arg1 db.UpdateProductionParams) (db.Production, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduction", arg0, arg1)
	ret0, _ := ret[0].(db.Production)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduction indicates an expected call of UpdateProduction.
func (mr *MockStoreMockRecorder) UpdateProduction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduction", reflect.TypeOf((*MockStore)(nil).UpdateProduction), arg0, arg1)
}