// Code generated by MockGen. DO NOT EDIT.
// Source: handler_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBankHandlerIface is a mock of BankHandlerIface interface.
type MockBankHandlerIface struct {
	ctrl     *gomock.Controller
	recorder *MockBankHandlerIfaceMockRecorder
}

// MockBankHandlerIfaceMockRecorder is the mock recorder for MockBankHandlerIface.
type MockBankHandlerIfaceMockRecorder struct {
	mock *MockBankHandlerIface
}

// NewMockBankHandlerIface creates a new mock instance.
func NewMockBankHandlerIface(ctrl *gomock.Controller) *MockBankHandlerIface {
	mock := &MockBankHandlerIface{ctrl: ctrl}
	mock.recorder = &MockBankHandlerIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankHandlerIface) EXPECT() *MockBankHandlerIfaceMockRecorder {
	return m.recorder
}

// CreateBank mocks base method.
func (m *MockBankHandlerIface) CreateBank(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateBank", w, r)
}

// CreateBank indicates an expected call of CreateBank.
func (mr *MockBankHandlerIfaceMockRecorder) CreateBank(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBank", reflect.TypeOf((*MockBankHandlerIface)(nil).CreateBank), w, r)
}

// GetAllBanks mocks base method.
func (m *MockBankHandlerIface) GetAllBanks(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllBanks", w, r)
}

// GetAllBanks indicates an expected call of GetAllBanks.
func (mr *MockBankHandlerIfaceMockRecorder) GetAllBanks(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBanks", reflect.TypeOf((*MockBankHandlerIface)(nil).GetAllBanks), w, r)
}

// GetBankDetails mocks base method.
func (m *MockBankHandlerIface) GetBankDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBankDetails", w, r)
}

// GetBankDetails indicates an expected call of GetBankDetails.
func (mr *MockBankHandlerIfaceMockRecorder) GetBankDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankDetails", reflect.TypeOf((*MockBankHandlerIface)(nil).GetBankDetails), w, r)
}

// RemoveBank mocks base method.
func (m *MockBankHandlerIface) RemoveBank(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveBank", w, r)
}

// RemoveBank indicates an expected call of RemoveBank.
func (mr *MockBankHandlerIfaceMockRecorder) RemoveBank(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBank", reflect.TypeOf((*MockBankHandlerIface)(nil).RemoveBank), w, r)
}

// UpdateBankDetails mocks base method.
func (m *MockBankHandlerIface) UpdateBankDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateBankDetails", w, r)
}

// UpdateBankDetails indicates an expected call of UpdateBankDetails.
func (mr *MockBankHandlerIfaceMockRecorder) UpdateBankDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBankDetails", reflect.TypeOf((*MockBankHandlerIface)(nil).UpdateBankDetails), w, r)
}

// MockAccountHandlerIface is a mock of AccountHandlerIface interface.
type MockAccountHandlerIface struct {
	ctrl     *gomock.Controller
	recorder *MockAccountHandlerIfaceMockRecorder
}

// MockAccountHandlerIfaceMockRecorder is the mock recorder for MockAccountHandlerIface.
type MockAccountHandlerIfaceMockRecorder struct {
	mock *MockAccountHandlerIface
}

// NewMockAccountHandlerIface creates a new mock instance.
func NewMockAccountHandlerIface(ctrl *gomock.Controller) *MockAccountHandlerIface {
	mock := &MockAccountHandlerIface{ctrl: ctrl}
	mock.recorder = &MockAccountHandlerIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountHandlerIface) EXPECT() *MockAccountHandlerIfaceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountHandlerIface) CreateAccount(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateAccount", w, r)
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountHandlerIfaceMockRecorder) CreateAccount(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountHandlerIface)(nil).CreateAccount), w, r)
}

// GetAccountDetails mocks base method.
func (m *MockAccountHandlerIface) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAccountDetails", w, r)
}

// GetAccountDetails indicates an expected call of GetAccountDetails.
func (mr *MockAccountHandlerIfaceMockRecorder) GetAccountDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDetails", reflect.TypeOf((*MockAccountHandlerIface)(nil).GetAccountDetails), w, r)
}

// GetAccountDetailsResponse mocks base method.
func (m *MockAccountHandlerIface) GetAccountDetailsResponse(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAccountDetailsResponse", w, r)
}

// GetAccountDetailsResponse indicates an expected call of GetAccountDetailsResponse.
func (mr *MockAccountHandlerIfaceMockRecorder) GetAccountDetailsResponse(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDetailsResponse", reflect.TypeOf((*MockAccountHandlerIface)(nil).GetAccountDetailsResponse), w, r)
}

// UpdateAccountDetails mocks base method.
func (m *MockAccountHandlerIface) UpdateAccountDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateAccountDetails", w, r)
}

// UpdateAccountDetails indicates an expected call of UpdateAccountDetails.
func (mr *MockAccountHandlerIfaceMockRecorder) UpdateAccountDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountDetails", reflect.TypeOf((*MockAccountHandlerIface)(nil).UpdateAccountDetails), w, r)
}

// MockCompositeIface is a mock of CompositeIface interface.
type MockCompositeIface struct {
	ctrl     *gomock.Controller
	recorder *MockCompositeIfaceMockRecorder
}

// MockCompositeIfaceMockRecorder is the mock recorder for MockCompositeIface.
type MockCompositeIfaceMockRecorder struct {
	mock *MockCompositeIface
}

// NewMockCompositeIface creates a new mock instance.
func NewMockCompositeIface(ctrl *gomock.Controller) *MockCompositeIface {
	mock := &MockCompositeIface{ctrl: ctrl}
	mock.recorder = &MockCompositeIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompositeIface) EXPECT() *MockCompositeIfaceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockCompositeIface) CreateAccount(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateAccount", w, r)
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockCompositeIfaceMockRecorder) CreateAccount(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockCompositeIface)(nil).CreateAccount), w, r)
}

// CreateBank mocks base method.
func (m *MockCompositeIface) CreateBank(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateBank", w, r)
}

// CreateBank indicates an expected call of CreateBank.
func (mr *MockCompositeIfaceMockRecorder) CreateBank(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBank", reflect.TypeOf((*MockCompositeIface)(nil).CreateBank), w, r)
}

// GetAccountDetails mocks base method.
func (m *MockCompositeIface) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAccountDetails", w, r)
}

// GetAccountDetails indicates an expected call of GetAccountDetails.
func (mr *MockCompositeIfaceMockRecorder) GetAccountDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDetails", reflect.TypeOf((*MockCompositeIface)(nil).GetAccountDetails), w, r)
}

// GetAccountDetailsResponse mocks base method.
func (m *MockCompositeIface) GetAccountDetailsResponse(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAccountDetailsResponse", w, r)
}

// GetAccountDetailsResponse indicates an expected call of GetAccountDetailsResponse.
func (mr *MockCompositeIfaceMockRecorder) GetAccountDetailsResponse(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDetailsResponse", reflect.TypeOf((*MockCompositeIface)(nil).GetAccountDetailsResponse), w, r)
}

// GetAllBanks mocks base method.
func (m *MockCompositeIface) GetAllBanks(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllBanks", w, r)
}

// GetAllBanks indicates an expected call of GetAllBanks.
func (mr *MockCompositeIfaceMockRecorder) GetAllBanks(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBanks", reflect.TypeOf((*MockCompositeIface)(nil).GetAllBanks), w, r)
}

// GetBankDetails mocks base method.
func (m *MockCompositeIface) GetBankDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBankDetails", w, r)
}

// GetBankDetails indicates an expected call of GetBankDetails.
func (mr *MockCompositeIfaceMockRecorder) GetBankDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankDetails", reflect.TypeOf((*MockCompositeIface)(nil).GetBankDetails), w, r)
}

// RemoveBank mocks base method.
func (m *MockCompositeIface) RemoveBank(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveBank", w, r)
}

// RemoveBank indicates an expected call of RemoveBank.
func (mr *MockCompositeIfaceMockRecorder) RemoveBank(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBank", reflect.TypeOf((*MockCompositeIface)(nil).RemoveBank), w, r)
}

// UpdateAccountDetails mocks base method.
func (m *MockCompositeIface) UpdateAccountDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateAccountDetails", w, r)
}

// UpdateAccountDetails indicates an expected call of UpdateAccountDetails.
func (mr *MockCompositeIfaceMockRecorder) UpdateAccountDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountDetails", reflect.TypeOf((*MockCompositeIface)(nil).UpdateAccountDetails), w, r)
}

// UpdateBankDetails mocks base method.
func (m *MockCompositeIface) UpdateBankDetails(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateBankDetails", w, r)
}

// UpdateBankDetails indicates an expected call of UpdateBankDetails.
func (mr *MockCompositeIfaceMockRecorder) UpdateBankDetails(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBankDetails", reflect.TypeOf((*MockCompositeIface)(nil).UpdateBankDetails), w, r)
}
