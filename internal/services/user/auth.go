package user

import (
	"context"

	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) Auth(ctx context.Context, in *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	// ...
	return nil, nil
}
