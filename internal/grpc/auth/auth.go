package auth

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) AuthenticateUser(ctx context.Context, in *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	user, err := s.storage.GetUser(in.Username)
	if err != nil {
		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	if err := auth.ComparePassword(user.PasswordHash, in.Password); err != nil {
		return nil, fmt.Errorf("unable to compare password: %w", err)
	}

	jwtToken, err := auth.GenerateJWT(int64(user.ID))
	if err != nil {
		return nil, fmt.Errorf("unable to generate jwt: %w", err)
	}

	decodedEncryptionKey, err := s.encryption.Decrypt(user.EncryptedMasterKey, []byte(constants.EncryptionKey))
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt master key: %w", err)
	}

	return &api.AuthenticateUserResponse{
		Token:         jwtToken,
		EncryptionKey: decodedEncryptionKey,
	}, nil
}
