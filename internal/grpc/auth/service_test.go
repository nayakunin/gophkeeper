package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/database"
	"github.com/nayakunin/gophkeeper/internal/grpc/auth/mocks"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_AuthenticateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type mockGetUserReply struct {
		user *database.User
		err  error
	}

	type mockComparePasswordReply struct {
		err error
	}

	type mockGenerateJWTReply struct {
		token string
		err   error
	}

	type mockDecryptReply struct {
		decryptedMasterKey []byte
		err                error
	}

	defaultIn := &api.AuthenticateUserRequest{
		Username: "username",
		Password: []byte("password"),
	}

	defaultUserMock := &database.User{
		ID:                 1,
		Username:           "username",
		Email:              "email",
		PasswordHash:       []byte("password hash"),
		EncryptedMasterKey: []byte("encrypted master key"),
	}

	testcases := []struct {
		name                 string
		out                  *api.AuthenticateUserResponse
		mockGetUserReply     *mockGetUserReply
		mockComparePassword  *mockComparePasswordReply
		mockGenerateJWTReply *mockGenerateJWTReply
		mockDecryptReply     *mockDecryptReply
		hasError             bool
	}{{
		name: "success",
		mockGetUserReply: &mockGetUserReply{
			user: defaultUserMock,
		},
		mockComparePassword: &mockComparePasswordReply{},
		mockGenerateJWTReply: &mockGenerateJWTReply{
			token: "token",
		},
		mockDecryptReply: &mockDecryptReply{
			decryptedMasterKey: []byte("decrypted master key"),
		},
		out: &api.AuthenticateUserResponse{
			Token:         "token",
			EncryptionKey: []byte("decrypted master key"),
		},
	}, {
		name: "error getting user",
		mockGetUserReply: &mockGetUserReply{
			err: errors.New("get user error"),
		},
		hasError: true,
	}, {
		name: "error comparing password",
		mockGetUserReply: &mockGetUserReply{
			user: defaultUserMock,
		},
		mockComparePassword: &mockComparePasswordReply{
			err: errors.New("compare password error"),
		},
		hasError: true,
	}, {
		name: "error generating jwt",
		mockGetUserReply: &mockGetUserReply{
			user: defaultUserMock,
		},
		mockComparePassword: &mockComparePasswordReply{},
		mockGenerateJWTReply: &mockGenerateJWTReply{
			err: errors.New("generate jwt error"),
		},
		hasError: true,
	}, {
		name: "error decrypting master key",
		mockGetUserReply: &mockGetUserReply{
			user: defaultUserMock,
		},
		mockComparePassword: &mockComparePasswordReply{},
		mockGenerateJWTReply: &mockGenerateJWTReply{
			token: "token",
		},
		mockDecryptReply: &mockDecryptReply{
			err: errors.New("decrypt error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			mockEncryption := mocks.NewMockEncryption(ctrl)
			mockAuthService := mocks.NewMockAuthService(ctrl)
			if tc.mockGetUserReply != nil {
				mockStorage.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(tc.mockGetUserReply.user, tc.mockGetUserReply.err)
			}
			if tc.mockComparePassword != nil {
				mockAuthService.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(tc.mockComparePassword.err)
			}
			if tc.mockGenerateJWTReply != nil {
				mockAuthService.EXPECT().GenerateJWT(gomock.Any()).Return(tc.mockGenerateJWTReply.token, tc.mockGenerateJWTReply.err)
			}
			if tc.mockDecryptReply != nil {
				mockEncryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tc.mockDecryptReply.decryptedMasterKey, tc.mockDecryptReply.err)
			}

			s := NewService(mockStorage, mockEncryption, mockAuthService)

			out, err := s.AuthenticateUser(context.Background(), defaultIn)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != nil {
				assert.Equal(t, tc.out, out)
			}
		})
	}
}
