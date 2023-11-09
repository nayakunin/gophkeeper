//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package password

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/password/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/password/mocks"
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

	s := NewService(c, e, a)

	assert.Equal(t, &Service{
		credentialsService: c,
		encryption:         e,
		apiPreparer:        ap,
		api:                a,
	}, s)
}
