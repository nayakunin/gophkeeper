package auth

import (
	"context"

	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) AuthenticateUser(ctx context.Context, in *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	baseResponse := &api.AuthenticateUserResponse{}

	user, err := s.Storage.GetUser(in.Username)
	if err != nil {
		return baseResponse, err
	}

	if err := auth.ComparePassword(in.Password, user.PasswordHash); err != nil {
		return baseResponse, err
	}

	jwtToken, err := auth.GenerateJWT(int64(user.ID))
	if err != nil {
		return baseResponse, err
	}

	return &api.AuthenticateUserResponse{
		Token:   jwtToken,
		Success: true,
	}, nil
}
