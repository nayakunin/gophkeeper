package commands

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := mocks.NewMockLocalStorage(ctrl)
	e := mocks.NewMockEncryption(ctrl)
	a := mocks.NewMockApi(ctrl)

	root := NewRoot(l, e, a)

	assert.Equal(t, &Root{
		cmd: &cobra.Command{},
	}, root)
}
