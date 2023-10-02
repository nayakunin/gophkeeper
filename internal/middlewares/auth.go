package middlewares

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/nayakunin/gophkeeper/pkg/utils/authcommon"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auth is used by a middleware to authenticate requests
func (s *Service) Auth(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	claims, err := s.a.ParseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid authcommon token: %v", err)
	}

	userID := s.a.UserClaimFromToken(claims)

	return context.WithValue(ctx, authcommon.UserIDKey, userID), nil
}
