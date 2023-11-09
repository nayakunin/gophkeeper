package password

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/password/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/get/password/output"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type credentialsServiceMock struct {
		key []byte
		err error
	}

	type apiMock struct {
		response *generated.GetLoginPasswordPairsResponse
		err      error
	}

	type outputMock struct {
		response []output.PasswordResult
		err      error
	}

	type args struct {
		label string
	}

	tests := []struct {
		name    string
		cs      *credentialsServiceMock
		api     *apiMock
		output  *outputMock
		args    *args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return errror if could not get credentials",
			cs: &credentialsServiceMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not parse request",
			cs: &credentialsServiceMock{
				key: []byte("test"),
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not get password",
			cs: &credentialsServiceMock{
				key: []byte("test"),
			},
			api: &apiMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not make response",
			cs: &credentialsServiceMock{
				key: []byte("test"),
			},
			api: &apiMock{
				response: &generated.GetLoginPasswordPairsResponse{
					LoginPasswordPairs: []*generated.LoginPasswordPair{
						{
							EncryptedPassword: []byte("test"),
							Id:                1,
							Login:             "test",
							Description:       "test",
							ServiceName:       "test",
						},
					},
				},
			},
			output: &outputMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := mocks.NewMockCredentialsService(ctrl)
			e := mocks.NewMockEncryption(ctrl)
			api := mocks.NewMockApi(ctrl)
			output := mocks.NewMockOutput(ctrl)

			if tt.cs != nil {
				cs.EXPECT().GetCredentials().Return("", tt.cs.key, tt.cs.err)
			}

			if tt.api != nil {
				api.EXPECT().GetLoginPasswordPairs(gomock.Any(), gomock.Any()).Return(tt.api.response, tt.api.err)
			}

			if tt.output != nil {
				output.EXPECT().MakeResponse(gomock.Any(), gomock.Any()).Return(tt.output.response, tt.output.err)
			}

			api.EXPECT().SetToken(gomock.Any()).Return()

			s := &Service{
				credentialsService: cs,
				encryption:         e,
				api:                api,
				output:             output,
			}

			cmd := s.GetCmd()

			if tt.args != nil {
				cmd.Flags().Set("service", tt.args.label)
			}

			err := cmd.RunE(cmd, []string{})

			tt.wantErr(t, err)
		})
	}
}
