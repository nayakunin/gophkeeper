package card

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/card/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/get/card/output"
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
		response *generated.GetBankCardDetailsResponse
		err      error
	}

	type outputMock struct {
		response []output.CardResult
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
			name: "should return error if could not get credentials",
			cs: &credentialsServiceMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not parse card request",
			cs: &credentialsServiceMock{
				key: []byte("test"),
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not get card",
			cs: &credentialsServiceMock{
				key: []byte("test"),
			},
			api: &apiMock{
				err: assert.AnError,
			},
			args: &args{
				label: "test",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not make card response",
			cs: &credentialsServiceMock{
				key: []byte("test"),
			},
			api: &apiMock{
				response: &generated.GetBankCardDetailsResponse{
					BankCardDetails: []*generated.BankCardDetail{
						{
							EncryptedCardNumber: []byte("test"),
							EncryptedExpiryDate: []byte("test"),
							Id:                  1,
							CardName:            "test",
							EncryptedCvc:        []byte("test"),
							Description:         "test",
						},
					},
				},
			},
			output: &outputMock{
				err: assert.AnError,
			},
			args: &args{
				label: "test",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			cs: &credentialsServiceMock{
				key: []byte("test"),
			},
			api: &apiMock{
				response: &generated.GetBankCardDetailsResponse{
					BankCardDetails: []*generated.BankCardDetail{
						{
							EncryptedCardNumber: []byte("test"),
							EncryptedExpiryDate: []byte("test"),
							Id:                  1,
							CardName:            "test",
							EncryptedCvc:        []byte("test"),
							Description:         "test",
						},
					},
				},
			},
			output: &outputMock{
				response: []output.CardResult{
					{
						Name:        "test",
						Number:      "test",
						Expiration:  "test",
						Cvc:         "test",
						Description: "test",
					},
				},
			},
			args: &args{
				label: "test",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := mocks.NewMockCredentialsService(ctrl)
			api := mocks.NewMockApi(ctrl)
			out := mocks.NewMockOutput(ctrl)
			e := mocks.NewMockEncryption(ctrl)

			if tt.cs != nil {
				cs.EXPECT().GetCredentials().Return("", tt.cs.key, tt.cs.err)
			}

			if tt.api != nil {
				api.EXPECT().GetCardDetails(gomock.Any(), gomock.Any()).Return(tt.api.response, tt.api.err)
				api.EXPECT().SetToken(gomock.Any()).Return()
			}

			if tt.output != nil {
				out.EXPECT().MakeResponse(gomock.Any(), gomock.Any()).Return(tt.output.response, tt.output.err)
			}

			s := &Service{
				credentialsService: cs,
				api:                api,
				output:             out,
				encryption:         e,
			}

			cmd := s.GetCmd()

			if tt.args != nil {
				cmd.Flags().Set("label", tt.args.label)
			}

			err := cmd.RunE(cmd, []string{})

			tt.wantErr(t, err)
		})
	}
}
