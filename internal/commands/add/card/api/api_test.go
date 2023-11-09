package api

import (
	"fmt"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/api/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/add/card/input"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_prepareCardRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		data          *input.ParseCardResult
		encryptionKey []byte
	}
	type EncryptionMock struct {
		result []byte
		err    error
	}
	type myMocks struct {
		e []EncryptionMock
	}

	defaultArgs := args{
		data: &input.ParseCardResult{
			Name:        "name",
			Number:      "number",
			Expiration:  "expiration",
			Cvc:         "cvc",
			Description: "description",
		},
		encryptionKey: []byte("key"),
	}

	tests := []struct {
		name    string
		args    args
		mocks   myMocks
		want    *api.AddBankCardDetailRequest
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error when encrypt returns error for number",
			args: defaultArgs,
			mocks: myMocks{
				e: []EncryptionMock{{
					err: fmt.Errorf("error"),
				}},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error when encrypt returns error for expiration",
			args: defaultArgs,
			mocks: myMocks{
				e: []EncryptionMock{{
					result: []byte("encrypted number"),
				}, {
					err: fmt.Errorf("error"),
				}},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error when encrypt returns error for cvc",
			args: defaultArgs,
			mocks: myMocks{
				e: []EncryptionMock{{
					result: []byte("encrypted number"),
				}, {
					result: []byte("encrypted expiration"),
				}, {
					err: fmt.Errorf("error"),
				}},
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			args: defaultArgs,
			mocks: myMocks{
				e: []EncryptionMock{{
					result: []byte("encrypted number"),
				}, {
					result: []byte("encrypted expiration"),
				}, {
					result: []byte("encrypted cvc"),
				}},
			},
			want: &api.AddBankCardDetailRequest{
				CardName:            "name",
				EncryptedCardNumber: []byte("encrypted number"),
				EncryptedExpiryDate: []byte("encrypted expiration"),
				EncryptedCvc:        []byte("encrypted cvc"),
				Description:         "description",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)

			for _, m := range tt.mocks.e {
				e.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(m.result, m.err)
			}

			s := &Service{
				encryption: e,
			}

			got, err := s.PrepareCardRequest(tt.args.data, tt.args.encryptionKey)
			if !tt.wantErr(t, err) {
				t.Fail()
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_NewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := mocks.NewMockEncryption(ctrl)

	s := NewService(e)

	assert.Equal(t, s, &Service{
		encryption: e,
	})
}
