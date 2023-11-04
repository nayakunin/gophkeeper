package add

import (
	"fmt"
	"os"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/mocks"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

//func TestService_binaryCmd(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	binaryCmd := NewService(c, e).binaryCmd()
//	buf := new(bytes.Buffer)
//	binaryCmd.SetOut(buf)
//
//	err := binaryCmd.Flags().Set("name", "test")
//	if err != nil {
//		return
//	}
//
//	type
//
//	type myMocks struct {
//
//	}
//
//
//	tests := []struct {
//		name string
//	}{{}}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			e := mocks.NewMockEncryption(ctrl)
//			c := mocks.NewMockCredentialsService(ctrl)
//
//			s := &Service{
//				credentialsService: c,
//				encryption:         e,
//			}
//			if got := s.binaryCmd(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("binaryCmd() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestService_parseBinaryRequest(t *testing.T) {
	type args struct {
		cmd *cobra.Command
	}

	cmd0 := &cobra.Command{}

	cmd1 := &cobra.Command{}
	cmd1.Flags().String("filepath", "", "")

	cmd2 := &cobra.Command{}
	cmd2.Flags().String("description", "", "")

	cmd3 := &cobra.Command{}
	cmd3.Flags().String("filepath", "", "")
	cmd3.Flags().String("description", "", "")

	tests := []struct {
		name    string
		args    args
		want    *parseBinaryResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error when no flags are set",
			args: args{
				cmd: cmd0,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should return error when only filepath flag is set",
			args: args{
				cmd: cmd1,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should return error when only description flag is set",
			args: args{
				cmd: cmd2,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should return no error when both flags are set",
			args: args{
				cmd: cmd3,
			},
			want: &parseBinaryResult{
				Filepath:    "",
				Description: "",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got, err := s.parseBinaryRequest(tt.args.cmd)
			if !tt.wantErr(t, err, fmt.Sprintf("parseBinaryRequest(%v)", tt.args.cmd)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseBinaryRequest(%v)", tt.args.cmd)
		})
	}
}

func TestService_prepareBinaryRequest(t *testing.T) {
	type args struct {
		result        *parseBinaryResult
		encryptionKey []byte
	}
	type encryptionMock struct {
		result []byte
		err    error
	}
	type osMock struct {
		result []byte
		err    error
	}
	type myMocks struct {
		e  *encryptionMock
		os *osMock
	}

	defaultArgs := args{
		result: &parseBinaryResult{
			Filepath:    "",
			Description: "",
		},
		encryptionKey: []byte("key"),
	}

	tests := []struct {
		name    string
		args    args
		mocks   myMocks
		want    *api.AddBinaryDataRequest
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return file error",
			mocks: myMocks{
				os: &osMock{
					err: fmt.Errorf("error"),
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return encryption error",
			mocks: myMocks{
				os: &osMock{
					result: []byte("test"),
				},
				e: &encryptionMock{
					err: fmt.Errorf("error"),
				},
			},
			args:    defaultArgs,
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			mocks: myMocks{
				os: &osMock{
					result: []byte("test"),
				},
				e: &encryptionMock{
					result: []byte("test"),
				},
			},
			args: defaultArgs,
			want: &api.AddBinaryDataRequest{
				EncryptedData: []byte("test"),
				Description:   defaultArgs.result.Description,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := mocks.NewMockEncryption(ctrl)
			if tt.mocks.e != nil {
				e.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tt.mocks.e.result, tt.mocks.e.err)
			}

			if tt.mocks.os != nil {
				osReadFile = func(filename string) ([]byte, error) {
					return tt.mocks.os.result, tt.mocks.os.err
				}
				defer func() { osReadFile = os.ReadFile }()
			}

			s := &Service{
				encryption: e,
			}

			got, err := s.prepareBinaryRequest(tt.args.result, tt.args.encryptionKey)

			if !tt.wantErr(t, err, fmt.Sprintf("prepareBinaryRequest(%v)", tt.args.result)) {
				t.Fail()
				return
			}
			assert.Equalf(t, tt.want, got, "prepareBinaryRequest(%v)", tt.args.result)
		})
	}
}
