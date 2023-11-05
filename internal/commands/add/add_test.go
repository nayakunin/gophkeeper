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
	a := mocks.NewMockApi(ctrl)

	type args struct {
		credentialsService CredentialsService
		encryption         Encryption
		api                Api
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
			api:                a,
		},
		want: &Service{
			credentialsService: c,
			encryption:         e,
			api:                a,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.credentialsService, tt.args.encryption, tt.args.api)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := mocks.NewMockEncryption(ctrl)
	c := mocks.NewMockCredentialsService(ctrl)
	a := mocks.NewMockApi(ctrl)

	got := NewService(c, e, a).Handle()

	assert.True(t, got.HasSubCommands())
}
