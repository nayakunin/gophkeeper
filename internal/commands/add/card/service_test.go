//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package card

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/card/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/card/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
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
