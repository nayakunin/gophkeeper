package registration

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) RegisterUser(ctx context.Context, in *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	passwordHash, err := auth.HashPassword(in.Password)
	if err != nil {
		return &api.RegisterUserResponse{
			Message: "Unable to hash password",
			Success: false,
		}, fmt.Errorf("unable to hash password: %w", err)
	}

	if err := s.Storage.CreateUser(in.Username, passwordHash); err != nil {
		return &api.RegisterUserResponse{
			Message: "Unable to create user",
			Success: false,
		}, fmt.Errorf("unable to create user: %w", err)
	}

	return &api.RegisterUserResponse{
		Message: "User created",
		Success: true,
	}, nil
}
