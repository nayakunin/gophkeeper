//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package text

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/text/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	a := mocks.NewMockApi(ctrl)
	c := mocks.NewMockCredentialsService(ctrl)
	e := mocks.NewMockEncryption(ctrl)

	ap := api.NewService(e)

	s := NewService(c, e, a)

	assert.Equal(t, &Service{
		credentialsService: c,
		encryption:         e,
		apiPreparer:        ap,
		api:                a,
	}, s)
}
