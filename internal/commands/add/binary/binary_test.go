package binary

import (
	"fmt"
	"testing"

	binaryMocks "github.com/nayakunin/gophkeeper/internal/commands/add/binary/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/add/mocks"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		filepath    string
		description string
	}

	type apiMock struct {
		err error
	}
	type apiPreparerMock struct {
		result *api.AddBinaryDataRequest
		err    error
	}
	type credentialsServiceMock struct {
		result        string
		encryptionKey []byte
		err           error
	}
	type myMocks struct {
		ap *apiPreparerMock
		a  *apiMock
		c  *credentialsServiceMock
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
				result:        "token",
				encryptionKey: []byte("key"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when prepareBinaryRequest returns error",
		args: &args{
			filepath:    "filepath",
			description: "description",
		},
		mocks: myMocks{
			c: &credentialsServiceMock{
				result:        "token",
				encryptionKey: []byte("key"),
			},
			ap: &apiPreparerMock{
				err: fmt.Errorf("error"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when api returns error",
		args: &args{
			filepath:    "filepath",
			description: "description",
		},
		mocks: myMocks{
			c: &credentialsServiceMock{
				result:        "token",
				encryptionKey: []byte("key"),
			},
			ap: &apiPreparerMock{
				result: &api.AddBinaryDataRequest{},
			},
			a: &apiMock{
				err: fmt.Errorf("error"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return no error",
		args: &args{
			filepath:    "filepath",
			description: "description",
		},
		mocks: myMocks{
			c: &credentialsServiceMock{
				result:        "token",
				encryptionKey: []byte("key"),
			},
			ap: &apiPreparerMock{
				result: &api.AddBinaryDataRequest{},
			},
			a: &apiMock{
				err: nil,
			},
		},
		want: assert.NoError,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)
			c := mocks.NewMockCredentialsService(ctrl)
			ap := binaryMocks.NewMockApiPreparer(ctrl)
			a := binaryMocks.NewMockApi(ctrl)

			if tt.mocks.c != nil {
				c.EXPECT().GetCredentials().Return(tt.mocks.c.result, tt.mocks.c.encryptionKey, tt.mocks.c.err)
			}

			if tt.mocks.ap != nil {
				ap.EXPECT().PrepareBinaryRequest(gomock.Any(), gomock.Any()).Return(tt.mocks.ap.result, tt.mocks.ap.err)
			}

			if tt.mocks.a != nil {
				a.EXPECT().AddBinaryData(gomock.Any(), gomock.Any()).Return(tt.mocks.a.err)
			}

			s := &Service{
				credentialsService: c,
				encryption:         e,
				apiPreparer:        ap,
				api:                a,
			}

			binaryCmd := s.GetCmd()

			if tt.args != nil {
				err := binaryCmd.Flags().Set("filepath", tt.args.filepath)
				assert.NoError(t, err)

				err = binaryCmd.Flags().Set("description", tt.args.description)
				assert.NoError(t, err)
			}

			tt.want(t, binaryCmd.RunE(binaryCmd, []string{}))
		})
	}
}
