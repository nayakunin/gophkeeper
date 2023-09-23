package middlewares

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var tokenInfoKey struct{}

func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}

func userClaimFromToken(struct{}) string {
	return "foobar"
}

// Auth is used by a middleware to authenticate requests
func Auth(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid middlewares token: %v", err)
	}

	ctx = logging.InjectFields(ctx, logging.Fields{"middlewares.sub", userClaimFromToken(tokenInfo)})

	// WARNING: In production define your own type to avoid context collisions.
	return context.WithValue(ctx, tokenInfoKey, tokenInfo), nil
}
