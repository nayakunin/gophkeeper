package registration

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	"github.com/nayakunin/gophkeeper/pkg/utils/encryption"
	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) RegisterUser(ctx context.Context, in *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	passwordHash, err := auth.HashPassword(in.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %w", err)
	}

	encryptedMasterKey, err := encryption.Encrypt(hex.EncodeToString(in.GetEncryptionKey()), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, fmt.Errorf("unable to encrypt master key: %w", err)
	}

	userID, err := s.Storage.CreateUser(in.Username, passwordHash, encryptedMasterKey)
	if err != nil {
		return nil, fmt.Errorf("unable to create user: %w", err)
	}

	jwtToken, err := auth.GenerateJWT(userID)
	if err != nil {
		return nil, fmt.Errorf("unable to generate jwt: %w", err)
	}

	return &api.RegisterUserResponse{
		Token: jwtToken,
	}, nil
}
