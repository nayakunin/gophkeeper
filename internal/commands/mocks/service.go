// Code generated by MockGen. DO NOT EDIT.
// Source: root.go
//
// Generated by this command:
//
//	mockgen -source=root.go -destination=mocks/service.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	generated "github.com/nayakunin/gophkeeper/proto"
	gomock "go.uber.org/mock/gomock"
)

// MockLocalStorage is a mock of LocalStorage interface.
type MockLocalStorage struct {
	ctrl     *gomock.Controller
	recorder *MockLocalStorageMockRecorder
}

// MockLocalStorageMockRecorder is the mock recorder for MockLocalStorage.
type MockLocalStorageMockRecorder struct {
	mock *MockLocalStorage
}

// NewMockLocalStorage creates a new mock instance.
func NewMockLocalStorage(ctrl *gomock.Controller) *MockLocalStorage {
	mock := &MockLocalStorage{ctrl: ctrl}
	mock.recorder = &MockLocalStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalStorage) EXPECT() *MockLocalStorageMockRecorder {
	return m.recorder
}

// DeleteCredentials mocks base method.
func (m *MockLocalStorage) DeleteCredentials() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCredentials")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCredentials indicates an expected call of DeleteCredentials.
func (mr *MockLocalStorageMockRecorder) DeleteCredentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCredentials", reflect.TypeOf((*MockLocalStorage)(nil).DeleteCredentials))
}

// GetCredentials mocks base method.
func (m *MockLocalStorage) GetCredentials() (string, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentials")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCredentials indicates an expected call of GetCredentials.
func (mr *MockLocalStorageMockRecorder) GetCredentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*MockLocalStorage)(nil).GetCredentials))
}

// SaveCredentials mocks base method.
func (m *MockLocalStorage) SaveCredentials(token string, encryptionKey []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCredentials", token, encryptionKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCredentials indicates an expected call of SaveCredentials.
func (mr *MockLocalStorageMockRecorder) SaveCredentials(token, encryptionKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCredentials", reflect.TypeOf((*MockLocalStorage)(nil).SaveCredentials), token, encryptionKey)
}

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

// Encrypt mocks base method.
func (m *MockEncryption) Encrypt(text, key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", text, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockEncryptionMockRecorder) Encrypt(text, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEncryption)(nil).Encrypt), text, key)
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

// AddBinaryData mocks base method.
func (m *MockApi) AddBinaryData(ctx context.Context, in *generated.AddBinaryDataRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinaryData", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinaryData indicates an expected call of AddBinaryData.
func (mr *MockApiMockRecorder) AddBinaryData(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinaryData", reflect.TypeOf((*MockApi)(nil).AddBinaryData), ctx, in)
}

// AddCardData mocks base method.
func (m *MockApi) AddCardData(ctx context.Context, in *generated.AddBankCardDetailRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCardData", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCardData indicates an expected call of AddCardData.
func (mr *MockApiMockRecorder) AddCardData(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCardData", reflect.TypeOf((*MockApi)(nil).AddCardData), ctx, in)
}

// AddPasswordData mocks base method.
func (m *MockApi) AddPasswordData(ctx context.Context, in *generated.AddLoginPasswordPairRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPasswordData", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPasswordData indicates an expected call of AddPasswordData.
func (mr *MockApiMockRecorder) AddPasswordData(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPasswordData", reflect.TypeOf((*MockApi)(nil).AddPasswordData), ctx, in)
}

// AddTextData mocks base method.
func (m *MockApi) AddTextData(ctx context.Context, in *generated.AddTextDataRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTextData", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTextData indicates an expected call of AddTextData.
func (mr *MockApiMockRecorder) AddTextData(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTextData", reflect.TypeOf((*MockApi)(nil).AddTextData), ctx, in)
}
