//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package card

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/card/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	api := mocks.NewMockApi(ctrl)
	credentialsService := mocks.NewMockCredentialsService(ctrl)
	encryption := mocks.NewMockEncryption(ctrl)
	out := mocks.NewMockOutput(ctrl)

	s := NewService(encryption, credentialsService, api)

	assert.Equal(t, &Service{
		output:             out,
		credentialsService: credentialsService,
		encryption:         encryption,
		api:                api,
	}, s)
}
