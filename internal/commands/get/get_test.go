package get

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := mocks.NewMockEncryption(ctrl)
	c := mocks.NewMockCredentialsService(ctrl)
	a := mocks.NewMockApi(ctrl)

	got := NewService(c, e, a)

	assert.Equal(t, &Service{
		credentialsService: c,
		encryption:         e,
		api:                a,
	}, got)
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
