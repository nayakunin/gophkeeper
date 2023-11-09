package text

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/text/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/get/text/output"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type apiMock struct {
		response *generated.GetTextDataResponse
		err      error
	}

	type credentialsMock struct {
		key []byte
		err error
	}

	type outputMock struct {
		response []output.TextResult
		err      error
	}

	tests := []struct {
		name            string
		apiMock         *apiMock
		credentialsMock *credentialsMock
		outputMock      *outputMock
		wantErr         assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if could not get credentials",
			credentialsMock: &credentialsMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not get text data",
			credentialsMock: &credentialsMock{
				key: []byte("test"),
			},
			apiMock: &apiMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not make text response",
			credentialsMock: &credentialsMock{
				key: []byte("test"),
			},
			apiMock: &apiMock{
				response: &generated.GetTextDataResponse{
					TextData: []*generated.GetTextDataResponseItem{
						{
							EncryptedText: []byte("test"),
							Id:            1,
							Description:   "test",
						},
					},
				},
			},
			outputMock: &outputMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return response",
			apiMock: &apiMock{
				response: &generated.GetTextDataResponse{
					TextData: []*generated.GetTextDataResponseItem{
						{
							EncryptedText: []byte("test"),
							Id:            1,
							Description:   "test",
						},
					},
				},
			},
			credentialsMock: &credentialsMock{
				key: []byte("test"),
			},
			outputMock: &outputMock{
				response: []output.TextResult{
					{
						Description: "test",
						Text:        "test",
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := mocks.NewMockCredentialsService(ctrl)
			e := mocks.NewMockEncryption(ctrl)
			api := mocks.NewMockApi(ctrl)
			out := mocks.NewMockOutput(ctrl)

			if tt.apiMock != nil {
				api.EXPECT().GetTextData(gomock.Any()).Return(tt.apiMock.response, tt.apiMock.err)
				api.EXPECT().SetToken(gomock.Any()).Return()
			}

			if tt.credentialsMock != nil {
				cs.EXPECT().GetCredentials().Return("", tt.credentialsMock.key, tt.credentialsMock.err)
			}

			if tt.outputMock != nil {
				out.EXPECT().MakeResponse(gomock.Any(), gomock.Any()).Return(tt.outputMock.response, tt.outputMock.err)
			}

			s := &Service{
				api:                api,
				credentialsService: cs,
				output:             out,
				encryption:         e,
			}

			cmd := s.GetCmd()

			err := cmd.RunE(cmd, nil)

			tt.wantErr(t, err)
		})
	}
}
