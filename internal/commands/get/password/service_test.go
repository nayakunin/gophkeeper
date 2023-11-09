//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package password

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/password/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/get/password/output"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	api := mocks.NewMockApi(ctrl)
	cs := mocks.NewMockCredentialsService(ctrl)
	encryption := mocks.NewMockEncryption(ctrl)

	out := output.NewService(encryption)

	s := NewService(encryption, cs, api)

	assert.Equal(t, &Service{
		output:             out,
		credentialsService: cs,
		encryption:         encryption,
		api:                api,
	}, s)
}
