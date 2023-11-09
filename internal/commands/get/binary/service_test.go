package binary

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/mocks"
	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/output"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cs := mocks.NewMockCredentialsService(ctrl)
	es := mocks.NewMockEncryption(ctrl)
	api := mocks.NewMockApi(ctrl)

	out := output.NewService(es)

	s := NewService(es, cs, api)

	assert.Equal(t, &Service{
		encryption:         es,
		credentialsService: cs,
		output:             out,
		api:                api,
	}, s)
}
