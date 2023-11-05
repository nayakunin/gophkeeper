package api

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/api/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_prepareTextRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type encryptionMock struct {
		enctyptedText []byte
		err           error
	}
	type myMocks struct {
		encryption encryptionMock
	}
	type args struct {
		result        *input.ParseTextResult
		encryptionKey []byte
	}
	tests := []struct {
		name    string
		mocks   myMocks
		args    args
		want    *api.AddTextDataRequest
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if encryption failed",
			mocks: myMocks{
				encryption: encryptionMock{
					err: assert.AnError,
				},
			},
			args: args{
				result: &input.ParseTextResult{
					Text:        "text",
					Description: "description",
				},
				encryptionKey: []byte("test"),
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			mocks: myMocks{
				encryption: encryptionMock{
					enctyptedText: []byte("encrypted text"),
				},
			},
			args: args{
				result: &input.ParseTextResult{
					Text:        "text",
					Description: "description",
				},
				encryptionKey: []byte("test"),
			},
			want: &api.AddTextDataRequest{
				EncryptedText: []byte("encrypted text"),
				Description:   "description",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := mocks.NewMockEncryption(ctrl)

			e.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tt.mocks.encryption.enctyptedText, tt.mocks.encryption.err)

			s := &Service{
				encryption: e,
			}

			got, err := s.PrepareTextRequest(tt.args.result, tt.args.encryptionKey)

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
