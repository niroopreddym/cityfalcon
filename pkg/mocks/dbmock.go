// Code generated by MockGen. DO NOT EDIT.
// Source: db_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

// MockDbIface is a mock of DbIface interface.
type MockDbIface struct {
	ctrl     *gomock.Controller
	recorder *MockDbIfaceMockRecorder
}

// MockDbIfaceMockRecorder is the mock recorder for MockDbIface.
type MockDbIfaceMockRecorder struct {
	mock *MockDbIface
}

// NewMockDbIface creates a new mock instance.
func NewMockDbIface(ctrl *gomock.Controller) *MockDbIface {
	mock := &MockDbIface{ctrl: ctrl}
	mock.recorder = &MockDbIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDbIface) EXPECT() *MockDbIfaceMockRecorder {
	return m.recorder
}

// CreateConnection mocks base method.
func (m *MockDbIface) CreateConnection(connstring string) (*pgx.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConnection", connstring)
	ret0, _ := ret[0].(*pgx.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConnection indicates an expected call of CreateConnection.
func (mr *MockDbIfaceMockRecorder) CreateConnection(connstring interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConnection", reflect.TypeOf((*MockDbIface)(nil).CreateConnection), connstring)
}

// DbClose mocks base method.
func (m *MockDbIface) DbClose() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DbClose")
}

// DbClose indicates an expected call of DbClose.
func (mr *MockDbIfaceMockRecorder) DbClose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbClose", reflect.TypeOf((*MockDbIface)(nil).DbClose))
}

// DbExecuteConflict mocks base method.
func (m *MockDbIface) DbExecuteConflict(query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DbExecuteConflict", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DbExecuteConflict indicates an expected call of DbExecuteConflict.
func (mr *MockDbIfaceMockRecorder) DbExecuteConflict(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbExecuteConflict", reflect.TypeOf((*MockDbIface)(nil).DbExecuteConflict), varargs...)
}

// DbExecuteQuery mocks base method.
func (m *MockDbIface) DbExecuteQuery(query string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DbExecuteQuery", query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DbExecuteQuery indicates an expected call of DbExecuteQuery.
func (mr *MockDbIfaceMockRecorder) DbExecuteQuery(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbExecuteQuery", reflect.TypeOf((*MockDbIface)(nil).DbExecuteQuery), query)
}

// DbExecuteScalar mocks base method.
func (m *MockDbIface) DbExecuteScalar(query string, args ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DbExecuteScalar", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DbExecuteScalar indicates an expected call of DbExecuteScalar.
func (mr *MockDbIfaceMockRecorder) DbExecuteScalar(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbExecuteScalar", reflect.TypeOf((*MockDbIface)(nil).DbExecuteScalar), varargs...)
}

// DbGetMultipleRow mocks base method.
func (m *MockDbIface) DbGetMultipleRow(query string, args ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DbGetMultipleRow", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DbGetMultipleRow indicates an expected call of DbGetMultipleRow.
func (mr *MockDbIfaceMockRecorder) DbGetMultipleRow(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbGetMultipleRow", reflect.TypeOf((*MockDbIface)(nil).DbGetMultipleRow), varargs...)
}

// DbWriter mocks base method.
func (m *MockDbIface) DbWriter(query string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DbWriter", query)
	ret0, _ := ret[0].(error)
	return ret0
}

// DbWriter indicates an expected call of DbWriter.
func (mr *MockDbIfaceMockRecorder) DbWriter(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DbWriter", reflect.TypeOf((*MockDbIface)(nil).DbWriter), query)
}

// InitDbReader mocks base method.
func (m *MockDbIface) InitDbReader() (*pgx.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitDbReader")
	ret0, _ := ret[0].(*pgx.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitDbReader indicates an expected call of InitDbReader.
func (mr *MockDbIfaceMockRecorder) InitDbReader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitDbReader", reflect.TypeOf((*MockDbIface)(nil).InitDbReader))
}

// InitDbWriter mocks base method.
func (m *MockDbIface) InitDbWriter() (*pgx.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitDbWriter")
	ret0, _ := ret[0].(*pgx.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitDbWriter indicates an expected call of InitDbWriter.
func (mr *MockDbIfaceMockRecorder) InitDbWriter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitDbWriter", reflect.TypeOf((*MockDbIface)(nil).InitDbWriter))
}

// TxBegin mocks base method.
func (m *MockDbIface) TxBegin() (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxBegin")
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxBegin indicates an expected call of TxBegin.
func (mr *MockDbIfaceMockRecorder) TxBegin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxBegin", reflect.TypeOf((*MockDbIface)(nil).TxBegin))
}

// TxComplete mocks base method.
func (m *MockDbIface) TxComplete(tx pgx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxComplete", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// TxComplete indicates an expected call of TxComplete.
func (mr *MockDbIfaceMockRecorder) TxComplete(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxComplete", reflect.TypeOf((*MockDbIface)(nil).TxComplete), tx)
}

// TxExecuteStmt mocks base method.
func (m *MockDbIface) TxExecuteStmt(tx pgx.Tx, query string, args ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TxExecuteStmt", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxExecuteStmt indicates an expected call of TxExecuteStmt.
func (mr *MockDbIfaceMockRecorder) TxExecuteStmt(tx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxExecuteStmt", reflect.TypeOf((*MockDbIface)(nil).TxExecuteStmt), varargs...)
}

// TxQuery mocks base method.
func (m *MockDbIface) TxQuery(tx pgx.Tx, query string) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxQuery", tx, query)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxQuery indicates an expected call of TxQuery.
func (mr *MockDbIfaceMockRecorder) TxQuery(tx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxQuery", reflect.TypeOf((*MockDbIface)(nil).TxQuery), tx, query)
}