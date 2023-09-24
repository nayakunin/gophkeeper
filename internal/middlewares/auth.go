package middlewares

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	auth2 "github.com/nayakunin/gophkeeper/pkg/utils/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auth is used by a middleware to authenticate requests
func Auth(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	claims, err := auth2.ParseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	userID := auth2.UserClaimFromToken(claims)

	return context.WithValue(ctx, auth2.UserIDKey, userID), nil
}
