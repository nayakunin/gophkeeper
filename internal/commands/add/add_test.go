package add

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := mocks.NewMockEncryption(ctrl)
	c := mocks.NewMockCredentialsService(ctrl)
	type args struct {
		credentialsService CredentialsService
		encryption         Encryption
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{{
		name: "TestNewService",
		args: args{
			credentialsService: c,
			encryption:         e,
		},
		want: &Service{
			credentialsService: c,
			encryption:         e,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.credentialsService, tt.args.encryption)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := mocks.NewMockEncryption(ctrl)
	c := mocks.NewMockCredentialsService(ctrl)

	got := NewService(c, e).Handle()

	assert.True(t, got.HasSubCommands())
}
