// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSQLProvider is a mock of SQLProvider interface.
type MockSQLProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSQLProviderMockRecorder
}

// MockSQLProviderMockRecorder is the mock recorder for MockSQLProvider.
type MockSQLProviderMockRecorder struct {
	mock *MockSQLProvider
}

// NewMockSQLProvider creates a new mock instance.
func NewMockSQLProvider(ctrl *gomock.Controller) *MockSQLProvider {
	mock := &MockSQLProvider{ctrl: ctrl}
	mock.recorder = &MockSQLProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLProvider) EXPECT() *MockSQLProviderMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockSQLProvider) Begin() (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockSQLProviderMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockSQLProvider)(nil).Begin))
}

// BeginTx mocks base method.
func (m *MockSQLProvider) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, options)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockSQLProviderMockRecorder) BeginTx(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockSQLProvider)(nil).BeginTx), ctx, options)
}

// BindNamed mocks base method.
func (m *MockSQLProvider) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindNamed", query, arg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BindNamed indicates an expected call of BindNamed.
func (mr *MockSQLProviderMockRecorder) BindNamed(query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindNamed", reflect.TypeOf((*MockSQLProvider)(nil).BindNamed), query, arg)
}

// ExecContext mocks base method.
func (m *MockSQLProvider) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockSQLProviderMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockSQLProvider)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockSQLProvider) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockSQLProviderMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockSQLProvider)(nil).GetContext), varargs...)
}

// NamedExecContext mocks base method.
func (m *MockSQLProvider) NamedExecContext(context context.Context, query string, arg interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExecContext", context, query, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExecContext indicates an expected call of NamedExecContext.
func (mr *MockSQLProviderMockRecorder) NamedExecContext(context, query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExecContext", reflect.TypeOf((*MockSQLProvider)(nil).NamedExecContext), context, query, arg)
}

// NamedQueryContext mocks base method.
func (m *MockSQLProvider) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedQueryContext", ctx, query, arg)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedQueryContext indicates an expected call of NamedQueryContext.
func (mr *MockSQLProviderMockRecorder) NamedQueryContext(ctx, query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedQueryContext", reflect.TypeOf((*MockSQLProvider)(nil).NamedQueryContext), ctx, query, arg)
}

// Rebind mocks base method.
func (m *MockSQLProvider) Rebind(query string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebind", query)
	ret0, _ := ret[0].(string)
	return ret0
}

// Rebind indicates an expected call of Rebind.
func (mr *MockSQLProviderMockRecorder) Rebind(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebind", reflect.TypeOf((*MockSQLProvider)(nil).Rebind), query)
}

// SelectContext mocks base method.
func (m *MockSQLProvider) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockSQLProviderMockRecorder) SelectContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockSQLProvider)(nil).SelectContext), varargs...)
}
