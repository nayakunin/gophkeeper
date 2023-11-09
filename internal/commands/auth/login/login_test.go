package login

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/auth/login/mocks"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		username string
		password string
	}

	type apiMock struct {
		response *generated.AuthenticateUserResponse
		err      error
	}

	type storageMock struct {
		err error
	}

	type myMocks struct {
		api     *apiMock
		storage *storageMock
	}

	tests := []struct {
		name    string
		mocks   myMocks
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if api returns error",
			mocks: myMocks{
				api: &apiMock{
					err: assert.AnError,
				},
			},
			args: args{
				username: "test",
				password: "test",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if storage returns error",
			mocks: myMocks{
				api: &apiMock{
					response: &generated.AuthenticateUserResponse{
						Token:         "token",
						EncryptionKey: []byte("encryption key"),
					},
				},
				storage: &storageMock{
					err: assert.AnError,
				},
			},
			args: args{
				username: "test",
				password: "test",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			mocks: myMocks{
				api: &apiMock{
					response: &generated.AuthenticateUserResponse{
						Token:         "token",
						EncryptionKey: []byte("encryption key"),
					},
				},
				storage: &storageMock{},
			},
			args: args{
				username: "test",
				password: "test",
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := mocks.NewMockApi(ctrl)
			storageMock := mocks.NewMockStorage(ctrl)

			if tt.mocks.api != nil {
				apiMock.EXPECT().AuthenticateUser(gomock.Any(), gomock.Any()).Return(tt.mocks.api.response, tt.mocks.api.err)
			}

			if tt.mocks.storage != nil {
				storageMock.EXPECT().SaveCredentials(gomock.Any(), gomock.Any()).Return(tt.mocks.storage.err)
			}

			service := &Service{
				api:     apiMock,
				storage: storageMock,
			}

			cmd := service.GetCmd()

			cmd.Flags().Set("username", tt.args.username)
			cmd.Flags().Set("password", tt.args.password)

			err := cmd.RunE(cmd, []string{})

			tt.wantErr(t, err)
		})
	}
}
