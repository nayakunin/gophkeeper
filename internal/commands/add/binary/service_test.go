package binary

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_NewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mocks.NewMockCredentialsService(ctrl)
	e := mocks.NewMockEncryption(ctrl)
	a := mocks.NewMockApi(ctrl)

	ap := api.NewService(e)

	assert.Equal(t, &Service{
		credentialsService: c,
		encryption:         e,
		api:                a,
		apiPreparer:        ap,
	}, NewService(c, e, a))
}
