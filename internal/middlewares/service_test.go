package middlewares

import (
	"context"
	"errors"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/middlewares/mocks"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	"github.com/nayakunin/gophkeeper/pkg/utils/authcommon"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		a AuthClient
	}
	a := mocks.NewMockAuthClient(ctrl)

	tests := []struct {
		name string
		args args
		want *Service
	}{{
		name: "TestNewService",
		args: args{
			a: a,
		},
		want: &Service{
			a: a,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewService(tt.args.a)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_Auth(t *testing.T) {
	t.Skip();
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type mockParseToken struct {
		claims *authcommon.CustomClaims
		err    error
	}

	type mockUserClaimFromToken struct {
		userID int64
	}

	token := "token"
	md := utils.GetRequestMetadata(token)

	tests := []struct {
		name                   string
		in                     context.Context
		out                    context.Context
		mockParseToken         *mockParseToken
		mockUserClaimFromToken *mockUserClaimFromToken
		hasError               bool
	}{
		{
			name:     "invalid token",
			in:       context.Background(),
			hasError: true,
		},
		{
			name: "invalid token parsing",
			in:   metadata.NewOutgoingContext(context.Background(), md),
			mockParseToken: &mockParseToken{
				err: errors.New("invalid token"),
			},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := mocks.NewMockAuthClient(ctrl)
			if tt.mockUserClaimFromToken != nil {
				a.EXPECT().UserClaimFromToken(gomock.Any()).Return(tt.mockUserClaimFromToken.userID)
			}
			if tt.mockParseToken != nil {
				a.EXPECT().ParseToken(gomock.Any()).Return(tt.mockParseToken.claims, tt.mockParseToken.err)
			}

			s := NewService(a)
			got, err := s.Auth(tt.in)

			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.out, got)
		})
	}
}
