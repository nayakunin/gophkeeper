package password

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/password/mocks"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_passwordCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		service     string
		login       string
		password    string
		description string
	}

	defaultArgs := &args{
		service:     "service",
		login:       "login",
		password:    "password",
		description: "description",
	}

	type credentialsServiceMock struct {
		token string
		key   []byte
		err   error
	}
	type apiPreparerMock struct {
		request *generated.AddLoginPasswordPairRequest
		err     error
	}
	type apiMock struct {
		err error
	}
	type myMock struct {
		credentialsServiceMock *credentialsServiceMock
		apiPreparerMock        *apiPreparerMock
		apiMock                *apiMock
	}
	tests := []struct {
		name    string
		mock    myMock
		args    *args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if credentialsService.GetCredentials returns error",
			mock: myMock{
				credentialsServiceMock: &credentialsServiceMock{
					err: assert.AnError,
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if input parser returns error",
			mock: myMock{
				credentialsServiceMock: &credentialsServiceMock{
					token: "token",
					key:   []byte("key"),
				},
			},
			args:    &args{},
			wantErr: assert.Error,
		},
		{
			name: "should return error if apiPreparer.PreparePasswordRequest returns error",
			mock: myMock{
				credentialsServiceMock: &credentialsServiceMock{
					token: "token",
					key:   []byte("key"),
				},
				apiPreparerMock: &apiPreparerMock{
					err: assert.AnError,
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return error if api.AddPasswordData returns error",
			mock: myMock{
				credentialsServiceMock: &credentialsServiceMock{
					token: "token",
					key:   []byte("key"),
				},
				apiPreparerMock: &apiPreparerMock{
					request: &generated.AddLoginPasswordPairRequest{},
				},
				apiMock: &apiMock{
					err: assert.AnError,
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return nil if everything is ok",
			mock: myMock{
				credentialsServiceMock: &credentialsServiceMock{
					token: "token",
					key:   []byte("key"),
				},
				apiPreparerMock: &apiPreparerMock{
					request: &generated.AddLoginPasswordPairRequest{},
				},
				apiMock: &apiMock{},
			},
			args:    defaultArgs,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mocks.NewMockCredentialsService(ctrl)
			e := mocks.NewMockEncryption(ctrl)
			ap := mocks.NewMockApiPreparer(ctrl)
			a := mocks.NewMockApi(ctrl)

			if tt.mock.credentialsServiceMock != nil {
				c.EXPECT().GetCredentials().Return(tt.mock.credentialsServiceMock.token, tt.mock.credentialsServiceMock.key, tt.mock.credentialsServiceMock.err)
			}

			if tt.mock.apiPreparerMock != nil {
				ap.EXPECT().PreparePasswordRequest(gomock.Any(), gomock.Any()).Return(tt.mock.apiPreparerMock.request, tt.mock.apiPreparerMock.err)
			}

			if tt.mock.apiMock != nil {
				a.EXPECT().AddPasswordData(gomock.Any(), gomock.Any()).Return(tt.mock.apiMock.err)
				a.EXPECT().SetToken(gomock.Any()).Return()
			}

			s := &Service{
				credentialsService: c,
				encryption:         e,
				apiPreparer:        ap,
				api:                a,
			}

			cmd := s.GetCmd()

			if tt.args != nil {
				cmd.Flags().Set("service", tt.args.service)
				cmd.Flags().Set("login", tt.args.login)
				cmd.Flags().Set("password", tt.args.password)
				cmd.Flags().Set("description", tt.args.description)
			}

			err := cmd.RunE(cmd, []string{})

			if !tt.wantErr(t, err) {
				t.Fail()
				return
			}
		})
	}
}
