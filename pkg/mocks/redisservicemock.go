// Code generated by MockGen. DO NOT EDIT.
// Source: redisservice_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRedisService is a mock of IRedisService interface.
type MockIRedisService struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisServiceMockRecorder
}

// MockIRedisServiceMockRecorder is the mock recorder for MockIRedisService.
type MockIRedisServiceMockRecorder struct {
	mock *MockIRedisService
}

// NewMockIRedisService creates a new mock instance.
func NewMockIRedisService(ctrl *gomock.Controller) *MockIRedisService {
	mock := &MockIRedisService{ctrl: ctrl}
	mock.recorder = &MockIRedisServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedisService) EXPECT() *MockIRedisServiceMockRecorder {
	return m.recorder
}

// AddKey mocks base method.
func (m *MockIRedisService) AddKey(key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKey", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddKey indicates an expected call of AddKey.
func (mr *MockIRedisServiceMockRecorder) AddKey(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKey", reflect.TypeOf((*MockIRedisService)(nil).AddKey), key, value)
}

// ReadKey mocks base method.
func (m *MockIRedisService) ReadKey(key string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadKey", key)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadKey indicates an expected call of ReadKey.
func (mr *MockIRedisServiceMockRecorder) ReadKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadKey", reflect.TypeOf((*MockIRedisService)(nil).ReadKey), key)
}
