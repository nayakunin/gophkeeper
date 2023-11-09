package logout

import (
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/auth/logout/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_LogoutCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type storageMock struct {
		err error
	}

	tests := []struct {
		name    string
		mock    storageMock
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if storage returns error",
			mock: storageMock{
				err: assert.AnError,
			},
			wantErr: assert.Error,
		},
		{
			name:    "should return no error",
			mock:    storageMock{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := mocks.NewMockStorage(ctrl)
			s.EXPECT().DeleteCredentials().Return(tt.mock.err)

			service := &Service{
				storage: s,
			}

			cmd := service.GetCmd()

			got := cmd.RunE(cmd, nil)

			if !tt.wantErr(t, got) {
				return
			}
		})
	}
}
