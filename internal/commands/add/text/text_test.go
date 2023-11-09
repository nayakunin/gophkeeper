package text

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/mocks"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_textCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type credentialsServiceMock struct {
		token string
		key   []byte
		err   error
	}
	type apiPreparerMock struct {
		request *generated.AddTextDataRequest
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

	type args struct {
		text        string
		description string
	}

	defaultArgs := &args{
		text:        "text",
		description: "description",
	}

	tests := []struct {
		name    string
		args    *args
		mock    myMock
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
			args: &args{
				text: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if apiPreparer returns error",
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
			name: "should return error if api returns error",
			mock: myMock{
				credentialsServiceMock: &credentialsServiceMock{
					token: "token",
					key:   []byte("key"),
				},
				apiPreparerMock: &apiPreparerMock{
					request: &generated.AddTextDataRequest{},
				},
				apiMock: &apiMock{
					err: assert.AnError,
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			mock: myMock{
				credentialsServiceMock: &credentialsServiceMock{
					token: "token",
					key:   []byte("key"),
				},
				apiPreparerMock: &apiPreparerMock{
					request: &generated.AddTextDataRequest{},
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
			a := mocks.NewMockApi(ctrl)
			ap := mocks.NewMockApiPreparer(ctrl)

			if tt.mock.credentialsServiceMock != nil {
				c.EXPECT().GetCredentials().Return(tt.mock.credentialsServiceMock.token, tt.mock.credentialsServiceMock.key, tt.mock.credentialsServiceMock.err)
			}

			if tt.mock.apiPreparerMock != nil {
				ap.EXPECT().PrepareTextRequest(gomock.Any(), gomock.Any()).Return(tt.mock.apiPreparerMock.request, tt.mock.apiPreparerMock.err)
			}

			if tt.mock.apiMock != nil {
				a.EXPECT().AddTextData(gomock.Any(), gomock.Any()).Return(tt.mock.apiMock.err)
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
				cmd.Flags().Set("text", tt.args.text)
				cmd.Flags().Set("description", tt.args.description)
			}

			err := cmd.RunE(cmd, []string{})

			tt.wantErr(t, err)
		})
	}
}
