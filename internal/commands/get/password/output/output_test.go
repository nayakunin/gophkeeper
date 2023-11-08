//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/password/output/mocks"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_MakeResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type encryptionMock struct {
		data []byte
		err  error
	}

	type args struct {
		response *generated.GetLoginPasswordPairsResponse
	}
	tests := []struct {
		name           string
		encryptionMock *encryptionMock
		args           args
		want           []PasswordResult
		wantErr        assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if could not decrypt password",
			encryptionMock: &encryptionMock{
				err: assert.AnError,
			},
			args: args{
				response: &generated.GetLoginPasswordPairsResponse{
					LoginPasswordPairs: []*generated.LoginPasswordPair{
						{
							EncryptedPassword: []byte("test"),
							Id:                1,
							Login:             "test",
							Description:       "test",
							ServiceName:       "test",
						},
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return response",
			encryptionMock: &encryptionMock{
				data: []byte("test"),
			},
			args: args{
				response: &generated.GetLoginPasswordPairsResponse{
					LoginPasswordPairs: []*generated.LoginPasswordPair{
						{
							EncryptedPassword: []byte("test"),
							Id:                1,
							Login:             "test",
							Description:       "test",
							ServiceName:       "test",
						},
					},
				},
			},
			want: []PasswordResult{
				{
					Login:       "test",
					Password:    "test",
					Description: "test",
					ServiceName: "test",
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)

			if tt.encryptionMock != nil {
				e.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tt.encryptionMock.data, tt.encryptionMock.err)
			}

			s := NewService(e)

			got, err := s.MakeResponse(tt.args.response, []byte("test"))

			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
