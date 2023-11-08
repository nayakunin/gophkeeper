// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=mocks/service.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	generated "github.com/nayakunin/gophkeeper/proto"
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

// GenerateKey mocks base method.
func (m *MockEncryption) GenerateKey() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateKey")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKey indicates an expected call of GenerateKey.
func (mr *MockEncryptionMockRecorder) GenerateKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKey", reflect.TypeOf((*MockEncryption)(nil).GenerateKey))
}

// MockApi is a mock of Api interface.
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *MockApiMockRecorder
}

// MockApiMockRecorder is the mock recorder for MockApi.
type MockApiMockRecorder struct {
	mock *MockApi
}

// NewMockApi creates a new mock instance.
func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &MockApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApi) EXPECT() *MockApiMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockApi) RegisterUser(ctx context.Context, in *generated.RegisterUserRequest) (*generated.RegisterUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, in)
	ret0, _ := ret[0].(*generated.RegisterUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockApiMockRecorder) RegisterUser(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockApi)(nil).RegisterUser), ctx, in)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// SaveCredentials mocks base method.
func (m *MockStorage) SaveCredentials(token string, encryptionKey []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCredentials", token, encryptionKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCredentials indicates an expected call of SaveCredentials.
func (mr *MockStorageMockRecorder) SaveCredentials(token, encryptionKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCredentials", reflect.TypeOf((*MockStorage)(nil).SaveCredentials), token, encryptionKey)
}