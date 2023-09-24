package registration

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	"github.com/nayakunin/gophkeeper/pkg/utils/encryption"
	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) RegisterUser(ctx context.Context, in *api.RegisterUserRequest) (*api.Empty, error) {
	passwordHash, err := auth.HashPassword(in.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %w", err)
	}

	encryptionKey := encryption.GenerateKey()
	encryptedMasterKey, err := encryption.Encrypt(string(encryptionKey), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, fmt.Errorf("unable to encrypt master key: %w", err)
	}

	if err := s.Storage.CreateUser(in.Username, passwordHash, encryptedMasterKey); err != nil {
		return nil, fmt.Errorf("unable to create user: %w", err)
	}

	return &api.Empty{}, nil
}
