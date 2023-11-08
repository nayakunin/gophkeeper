//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package text

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/text/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	api := mocks.NewMockApi(ctrl)
	cs := mocks.NewMockCredentialsService(ctrl)
	e := mocks.NewMockEncryption(ctrl)
	out := mocks.NewMockOutput(ctrl)

	s := NewService(e, cs, api)

	assert.Equal(t, &Service{
		output:             out,
		credentialsService: cs,
		encryption:         e,
		api:                api,
	}, s)
}
