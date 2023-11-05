//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package logout

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/auth/logout/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockStorage(ctrl)

	service := NewService(s)

	assert.Equal(t, &Service{
		storage: s,
	}, service)
}
