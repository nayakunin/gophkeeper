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
	reflect "reflect"

	authcommon "github.com/nayakunin/gophkeeper/pkg/utils/authcommon"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthClient is a mock of AuthClient interface.
type MockAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthClientMockRecorder
}

// MockAuthClientMockRecorder is the mock recorder for MockAuthClient.
type MockAuthClientMockRecorder struct {
	mock *MockAuthClient
}

// NewMockAuthClient creates a new mock instance.
func NewMockAuthClient(ctrl *gomock.Controller) *MockAuthClient {
	mock := &MockAuthClient{ctrl: ctrl}
	mock.recorder = &MockAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthClient) EXPECT() *MockAuthClientMockRecorder {
	return m.recorder
}

// ParseToken mocks base method.
func (m *MockAuthClient) ParseToken(token string) (*authcommon.CustomClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(*authcommon.CustomClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthClientMockRecorder) ParseToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthClient)(nil).ParseToken), token)
}

// UserClaimFromToken mocks base method.
func (m *MockAuthClient) UserClaimFromToken(claims *authcommon.CustomClaims) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserClaimFromToken", claims)
	ret0, _ := ret[0].(int64)
	return ret0
}

// UserClaimFromToken indicates an expected call of UserClaimFromToken.
func (mr *MockAuthClientMockRecorder) UserClaimFromToken(claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserClaimFromToken", reflect.TypeOf((*MockAuthClient)(nil).UserClaimFromToken), claims)
}
