package binary

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/output"
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
		response *generated.GetBinaryDataResponse
		err      error
	}

	type outputMock struct {
		response []output.BinaryResult
		err      error
	}

	tests := []struct {
		name    string
		csm     *credentialsServiceMock
		api     *apiMock
		output  *outputMock
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if could not get credentials",
			csm: &credentialsServiceMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not get binary data",
			csm: &credentialsServiceMock{
				key: []byte("key"),
			},
			api: &apiMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not make binary response",
			csm: &credentialsServiceMock{
				key: []byte("key"),
			},
			api: &apiMock{
				response: &generated.GetBinaryDataResponse{
					BinaryData: []*generated.GetBinaryDataResponseItem{{
						EncryptedData: []byte("encrypted data"),
					}},
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
			csm := mocks.NewMockCredentialsService(ctrl)
			e := mocks.NewMockEncryption(ctrl)
			api := mocks.NewMockApi(ctrl)
			output := mocks.NewMockOutput(ctrl)

			if tt.csm != nil {
				csm.EXPECT().GetCredentials().Return("token", tt.csm.key, tt.csm.err)
			}

			if tt.api != nil {
				api.EXPECT().GetBinaryData(gomock.Any()).Return(tt.api.response, tt.api.err)
			}

			if tt.output != nil {
				output.EXPECT().MakeResponse(tt.api.response, tt.csm.key).Return(tt.output.response, tt.output.err)
			}

			api.EXPECT().SetToken(gomock.Any()).Return()

			s := &Service{
				credentialsService: csm,
				api:                api,
				output:             output,
				encryption:         e,
			}

			cmd := s.GetCmd()

			err := cmd.RunE(cmd, []string{})

			tt.wantErr(t, err)
		})
	}
}
