package register

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/auth/register/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockStorage(ctrl)
	a := mocks.NewMockApi(ctrl)
	e := mocks.NewMockEncryption(ctrl)

	service := NewService(e, s, a)

	assert.Equal(t, &Service{
		encryption: e,
		storage:    s,
		api:        a,
	}, service)
}
