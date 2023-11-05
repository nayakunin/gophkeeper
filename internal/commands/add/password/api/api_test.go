package api

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/password/api/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/add/password/input"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_PreparePasswordRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		result        *input.ParsePasswordResult
		encryptionKey []byte
	}

	type encryptionMock struct {
		encryptedPassword []byte
		err               error
	}

	tests := []struct {
		name    string
		args    args
		mock    encryptionMock
		want    *api.AddLoginPasswordPairRequest
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if encryption.Encrypt returns error",
			args: args{
				result: &input.ParsePasswordResult{
					ServiceName: "service",
					Login:       "login",
					Password:    "password",
					Description: "description",
				},
				encryptionKey: []byte("encryptionKey"),
			},
			mock: encryptionMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			args: args{
				result: &input.ParsePasswordResult{
					ServiceName: "service",
					Login:       "login",
					Password:    "password",
					Description: "description",
				},
				encryptionKey: []byte("encryptionKey"),
			},
			mock: encryptionMock{
				encryptedPassword: []byte("encryptedPassword"),
			},
			want: &api.AddLoginPasswordPairRequest{
				ServiceName:       "service",
				Login:             "login",
				EncryptedPassword: []byte("encryptedPassword"),
				Description:       "description",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)
			e.EXPECT().Encrypt([]byte(tt.args.result.Password), tt.args.encryptionKey).Return(tt.mock.encryptedPassword, tt.mock.err)

			s := &Service{
				encryption: e,
			}

			got, err := s.PreparePasswordRequest(tt.args.result, tt.args.encryptionKey)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := mocks.NewMockEncryption(ctrl)

	s := NewService(e)

	assert.Equal(t, &Service{
		encryption: e,
	}, s)
}
