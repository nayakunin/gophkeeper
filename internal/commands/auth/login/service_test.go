//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package login

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/auth/login/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockStorage(ctrl)
	a := mocks.NewMockApi(ctrl)

	service := NewService(s, a)

	assert.Equal(t, &Service{
		storage: s,
		api:     a,
	}, service)
}
