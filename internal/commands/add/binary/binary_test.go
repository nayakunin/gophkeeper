package binary

import (
	"fmt"
	"testing"

	binaryMocks "github.com/nayakunin/gophkeeper/internal/commands/add/binary/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/add/mocks"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_binaryCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		cmd *cobra.Command
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
		a *apiPreparerMock
		c *credentialsServiceMock
	}

	defaultCmd := &cobra.Command{}
	cmdWithFlags := &cobra.Command{}
	err := cmdWithFlags.Flags().Set("filepath", "filepath")
	assert.NoError(t, err)
	err = cmdWithFlags.Flags().Set("description", "description")
	assert.NoError(t, err)

	defaultArgs := args{
		cmd: defaultCmd,
	}
	cmdWithFlagsArgs := args{
		cmd: cmdWithFlags,
	}

	tests := []struct {
		name  string
		args  args
		mocks myMocks
		want  assert.ErrorAssertionFunc
	}{{
		name: "should return error when getCredentials returns error",
		args: defaultArgs,
		mocks: myMocks{
			c: &credentialsServiceMock{
				err: fmt.Errorf("error"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when parseBinaryRequest returns error",
		args: defaultArgs,
		mocks: myMocks{
			c: &credentialsServiceMock{
				result:        "token",
				encryptionKey: []byte("key"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when prepareBinaryRequest returns error",
		args: cmdWithFlagsArgs,
		mocks: myMocks{
			c: &credentialsServiceMock{
				result:        "token",
				encryptionKey: []byte("key"),
			},
			a: &apiPreparerMock{
				err: fmt.Errorf("error"),
			},
		},
		want: assert.Error,
	}, {
		name: "should return error when grpc.Dial returns error",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)
			c := mocks.NewMockCredentialsService(ctrl)
			a := binaryMocks.NewMockApi(ctrl)

			if tt.mocks.c != nil {
				c.EXPECT().GetCredentials().Return(tt.mocks.c.result, tt.mocks.c.encryptionKey, tt.mocks.c.err)
			}

			if tt.mocks.a != nil {
				a.EXPECT().PrepareBinaryRequest(gomock.Any(), gomock.Any()).Return(tt.mocks.a.result, tt.mocks.a.err)
			}

			s := &Service{
				credentialsService: c,
				encryption:         e,
				apiPreparer:        a,
			}

			binaryCmd := NewService(c, e).binaryCmd()
			assert.NoError(t, err)

			if !tt.want(t, binaryCmd.RunE(tt.args.cmd, []string{}), fmt.Sprintf("binaryCmd(%v)", tt.args.cmd)) {
				return
			}
		})
	}
}
