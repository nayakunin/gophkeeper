package card

import (
	"fmt"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/card/mocks"
	addMocks "github.com/nayakunin/gophkeeper/internal/commands/add/mocks"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		label       string
		number      string
		expiration  string
		cvc         string
		description string
	}

	type apiMock struct {
		err error
	}
	type apiPreparerMock struct {
		result *generated.AddBankCardDetailRequest
		err    error
	}
	type credentialsServiceMock struct {
		token         string
		encryptionKey []byte
		err           error
	}
	type myMocks struct {
		ap *apiPreparerMock
		a  *apiMock
		c  *credentialsServiceMock
	}

	defaultArgs := args{
		label:       "label",
		number:      "number",
		expiration:  "expiration",
		cvc:         "cvc",
		description: "description",
	}

	tests := []struct {
		name  string
		args  *args
		mocks myMocks
		want  assert.ErrorAssertionFunc
	}{{
		name: "should return error when getCredentials returns error",
		mocks: myMocks{
			c: &credentialsServiceMock{
				err: fmt.Errorf("error"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when parseBinaryRequest returns error",
		mocks: myMocks{
			c: &credentialsServiceMock{
				token:         "token",
				encryptionKey: []byte("key"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when prepareBinaryRequest returns error",
		args: &defaultArgs,
		mocks: myMocks{
			c: &credentialsServiceMock{
				token:         "token",
				encryptionKey: []byte("key"),
			},
			ap: &apiPreparerMock{
				err: fmt.Errorf("error"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when api returns error",
		args: &defaultArgs,
		mocks: myMocks{
			c: &credentialsServiceMock{
				token:         "token",
				encryptionKey: []byte("key"),
			},
			ap: &apiPreparerMock{
				result: &generated.AddBankCardDetailRequest{},
			},
			a: &apiMock{
				err: fmt.Errorf("error"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return no error",
		args: &defaultArgs,
		mocks: myMocks{
			c: &credentialsServiceMock{
				token:         "token",
				encryptionKey: []byte("key"),
			},
			ap: &apiPreparerMock{
				result: &generated.AddBankCardDetailRequest{},
			},
			a: &apiMock{
				err: nil,
			},
		},
		want: assert.NoError,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := addMocks.NewMockEncryption(ctrl)
			c := addMocks.NewMockCredentialsService(ctrl)
			ap := mocks.NewMockApiPreparer(ctrl)
			a := mocks.NewMockApi(ctrl)

			if tt.mocks.c != nil {
				c.EXPECT().GetCredentials().Return(tt.mocks.c.token, tt.mocks.c.encryptionKey, tt.mocks.c.err)
			}

			if tt.mocks.ap != nil {
				ap.EXPECT().PrepareCardRequest(gomock.Any(), gomock.Any()).Return(tt.mocks.ap.result, tt.mocks.ap.err)
			}

			if tt.mocks.a != nil {
				a.EXPECT().AddCardData(gomock.Any(), gomock.Any()).Return(tt.mocks.a.err)
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
				cmd.Flags().Set("label", tt.args.label)
				cmd.Flags().Set("number", tt.args.number)
				cmd.Flags().Set("expiration", tt.args.expiration)
				cmd.Flags().Set("cvc", tt.args.cvc)
				cmd.Flags().Set("description", tt.args.description)
			}

			tt.want(t, cmd.RunE(cmd, []string{}))
		})
	}
}
