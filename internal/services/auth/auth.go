package auth

import (
	"context"

	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) AuthenticateUser(ctx context.Context, in *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	// ...
	return nil, nil
}
