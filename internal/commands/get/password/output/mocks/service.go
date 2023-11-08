// Code generated by MockGen. DO NOT EDIT.
// Source: output.go
//
// Generated by this command:
//
//	mockgen -source=output.go -destination=mocks/service.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEncryption is a mock of Encryption interface.
type MockEncryption struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionMockRecorder
}

// MockEncryptionMockRecorder is the mock recorder for MockEncryption.
type MockEncryptionMockRecorder struct {
	mock *MockEncryption
}

// NewMockEncryption creates a new mock instance.
func NewMockEncryption(ctrl *gomock.Controller) *MockEncryption {
	mock := &MockEncryption{ctrl: ctrl}
	mock.recorder = &MockEncryptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryption) EXPECT() *MockEncryptionMockRecorder {
	return m.recorder
}

// Decrypt mocks base method.
func (m *MockEncryption) Decrypt(text, key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", text, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockEncryptionMockRecorder) Decrypt(text, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockEncryption)(nil).Decrypt), text, key)
}