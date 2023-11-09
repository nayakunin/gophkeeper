package register

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/auth/register/mocks"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_RegisterCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		username string
		password string
	}

	type encryptionMock struct {
		encryptionKey []byte
		err           error
	}

	type apiMock struct {
		response *generated.RegisterUserResponse
		err      error
	}

	type storageMock struct {
		err error
	}

	type myMocks struct {
		api        *apiMock
		storage    *storageMock
		encryption *encryptionMock
	}

	defaultArgs := args{
		username: "test",
		password: "test",
	}

	tests := []struct {
		name    string
		args    args
		mocks   myMocks
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if encryption returns error",
			mocks: myMocks{
				encryption: &encryptionMock{
					err: assert.AnError,
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return error if api returns error",
			mocks: myMocks{
				encryption: &encryptionMock{
					encryptionKey: []byte("encryption key"),
				},
				api: &apiMock{
					err: assert.AnError,
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return error if storage returns error",
			mocks: myMocks{
				encryption: &encryptionMock{
					encryptionKey: []byte("encryption key"),
				},
				api: &apiMock{
					response: &generated.RegisterUserResponse{
						Token: "token",
					},
				},
				storage: &storageMock{
					err: assert.AnError,
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			mocks: myMocks{
				encryption: &encryptionMock{
					encryptionKey: []byte("encryption key"),
				},
				api: &apiMock{
					response: &generated.RegisterUserResponse{
						Token: "token",
					},
				},
				storage: &storageMock{},
			},
			args:    defaultArgs,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)
			a := mocks.NewMockApi(ctrl)
			s := mocks.NewMockStorage(ctrl)

			if tt.mocks.encryption != nil {
				e.EXPECT().GenerateKey().Return(tt.mocks.encryption.encryptionKey, tt.mocks.encryption.err)
			}

			if tt.mocks.api != nil {
				a.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(tt.mocks.api.response, tt.mocks.api.err)
			}

			if tt.mocks.storage != nil {
				s.EXPECT().SaveCredentials(gomock.Any(), gomock.Any()).Return(tt.mocks.storage.err)
			}

			service := NewService(e, s, a)

			cmd := service.GetCmd()

			cmd.Flags().Set("username", tt.args.username)
			cmd.Flags().Set("password", tt.args.password)

			got := cmd.RunE(cmd, nil)

			tt.wantErr(t, got)
		})
	}
}
