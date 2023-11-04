package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/api/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/input"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_PrepareBinaryRequest(t *testing.T) {
	type args struct {
		result        *input.ParseBinaryResult
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
		result: &input.ParseBinaryResult{
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

			got, err := s.PrepareBinaryRequest(tt.args.result, tt.args.encryptionKey)

			if !tt.wantErr(t, err, fmt.Sprintf("prepareBinaryRequest(%v)", tt.args.result)) {
				t.Fail()
				return
			}
			assert.Equalf(t, tt.want, got, "prepareBinaryRequest(%v)", tt.args.result)
		})
	}
}
