//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/output/mocks"
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
		response *generated.GetBankCardDetailsResponse
	}

	defaultArgs := args{
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
	}

	tests := []struct {
		name    string
		args    args
		em      []encryptionMock
		want    []CardResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if could not decrypt card number",
			args: defaultArgs,
			em: []encryptionMock{
				{
					err: assert.AnError,
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not decrypt card expiration date",
			args: defaultArgs,
			em: []encryptionMock{
				{
					data: []byte("test"),
				},
				{
					err: assert.AnError,
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if could not decrypt card CVC",
			args: defaultArgs,
			em: []encryptionMock{
				{
					data: []byte("test"),
				},
				{
					data: []byte("test2"),
				},
				{
					err: assert.AnError,
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return card result",
			args: defaultArgs,
			em: []encryptionMock{
				{
					data: []byte("test"),
				},
				{
					data: []byte("test2"),
				},
				{
					data: []byte("test3"),
				},
			},
			want: []CardResult{
				{
					Name:        "test",
					Number:      "test",
					Expiration:  "test2",
					Cvc:         "test3",
					Description: "test",
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)

			for i, em := range tt.em {
				e.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(em.data, em.err).Times(i + 1)
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
