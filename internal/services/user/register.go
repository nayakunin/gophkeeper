package user

import (
	"context"

	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) Register(ctx context.Context, in *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	// ...
	return nil, nil
}
