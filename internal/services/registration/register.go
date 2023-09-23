package registration

import (
	"context"
	"fmt"

	api "github.com/nayakunin/gophkeeper/proto"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) RegisterUser(ctx context.Context, in *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	passwordHash, err := HashPassword(in.Password)
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
