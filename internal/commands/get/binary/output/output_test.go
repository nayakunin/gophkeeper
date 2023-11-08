package output

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/output/mocks"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_MakeResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		input *generated.GetBinaryDataResponse
	}

	type encryptionMock struct {
		data []byte
		err  error
	}

	tests := []struct {
		name           string
		args           args
		encryptionMock encryptionMock
		want           []BinaryResult
		wantErr        assert.ErrorAssertionFunc
	}{{
		name: "should return error if could not decrypt data",
		args: args{
			input: &generated.GetBinaryDataResponse{
				BinaryData: []*generated.GetBinaryDataResponseItem{{
					EncryptedData: []byte("encrypted data"),
					Description:   "description",
				}},
			},
		},
		encryptionMock: encryptionMock{
			err: assert.AnError,
		},
		wantErr: assert.Error,
	}, {
		name: "should return decrypted data",
		args: args{
			input: &generated.GetBinaryDataResponse{
				BinaryData: []*generated.GetBinaryDataResponseItem{{
					EncryptedData: []byte("encrypted data"),
					Description:   "description",
				}},
			},
		},
		encryptionMock: encryptionMock{
			data: []byte("data"),
		},
		want: []BinaryResult{{
			Data:        []byte("data"),
			Description: "description",
		}},
		wantErr: assert.NoError,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encryption := mocks.NewMockEncryption(ctrl)
			encryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tt.encryptionMock.data, tt.encryptionMock.err)

			s := NewService(encryption)
			got, err := s.MakeResponse(tt.args.input, []byte("key"))

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
