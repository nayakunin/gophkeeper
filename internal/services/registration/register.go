package registration

import (
	"context"
	"fmt"

	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) RegisterUser(ctx context.Context, in *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	if err := s.Storage.CreateUser(in); err != nil {
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
