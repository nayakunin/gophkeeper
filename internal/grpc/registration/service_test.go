package registration

import (
	"errors"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/grpc/registration/mocks"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type MockCreateUserReply struct {
		userID int64
		err    error
	}

	type MockEncryptReply struct {
		encryptedMasterKey []byte
		err                error
	}

	type MockGenerateJWTReply struct {
		token string
		err   error
	}

	type MockHashPassword struct {
		hash []byte
		err  error
	}

	defaultIn := &api.RegisterUserRequest{
		Username:      "username",
		Password:      "password",
		EncryptionKey: []byte("encryption key"),
	}
	defaultOut := &api.RegisterUserResponse{
		Token: "token",
	}
	defaultHashMock := &MockHashPassword{
		hash: []byte("hash"),
	}
	defaultGenerateJWTMock := &MockGenerateJWTReply{
		token: "token",
	}
	defaultStorageMock := &MockCreateUserReply{
		userID: 1,
	}
	defaultEncryptionMock := &MockEncryptReply{
		encryptedMasterKey: []byte("encrypted master key"),
	}

	testCases := []struct {
		name                 string
		in                   *api.RegisterUserRequest
		out                  *api.RegisterUserResponse
		hasError             bool
		mockStorageReply     *MockCreateUserReply
		mockEncryptionReply  *MockEncryptReply
		mockGenerateJWTReply *MockGenerateJWTReply
		mockHashPassword     *MockHashPassword
	}{{
		name:                 "success",
		in:                   defaultIn,
		out:                  defaultOut,
		mockStorageReply:     defaultStorageMock,
		mockEncryptionReply:  defaultEncryptionMock,
		mockHashPassword:     defaultHashMock,
		mockGenerateJWTReply: defaultGenerateJWTMock,
	}, {
		name:     "unable to hash password",
		in:       defaultIn,
		hasError: true,
		mockHashPassword: &MockHashPassword{
			err: errors.New("hash error"),
		},
	}, {
		name:             "unable to encrypt master key",
		in:               defaultIn,
		hasError:         true,
		mockHashPassword: defaultHashMock,
		mockEncryptionReply: &MockEncryptReply{
			err: errors.New("encrypt error"),
		},
	}, {
		name:                "unable to create user",
		in:                  defaultIn,
		hasError:            true,
		mockHashPassword:    defaultHashMock,
		mockEncryptionReply: defaultEncryptionMock,
		mockStorageReply: &MockCreateUserReply{
			err: errors.New("create user error"),
		},
	}, {
		name:                "unable to generate jwt",
		in:                  defaultIn,
		hasError:            true,
		mockHashPassword:    defaultHashMock,
		mockEncryptionReply: defaultEncryptionMock,
		mockStorageReply:    defaultStorageMock,
		mockGenerateJWTReply: &MockGenerateJWTReply{
			err: errors.New("generate jwt error"),
		},
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockStorageReply != nil {
				mockStorage.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockStorageReply.userID, tc.mockStorageReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockEncryptionReply != nil {
				mockEncryption.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tc.mockEncryptionReply.encryptedMasterKey, tc.mockEncryptionReply.err)
			}
			mockAuth := mocks.NewMockAuthService(ctrl)
			if tc.mockHashPassword != nil {
				mockAuth.EXPECT().HashPassword(gomock.Any()).Return(tc.mockHashPassword.hash, tc.mockHashPassword.err)
			}
			if tc.mockGenerateJWTReply != nil {
				mockAuth.EXPECT().GenerateJWT(gomock.Any()).Return(tc.mockGenerateJWTReply.token, tc.mockGenerateJWTReply.err)
			}

			s := NewService(mockStorage, mockEncryption, mockAuth)

			out, err := s.RegisterUser(nil, tc.in)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.out, out)
		})
	}
}
