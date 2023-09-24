package auth

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	"github.com/nayakunin/gophkeeper/pkg/utils/encryption"
	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) AuthenticateUser(ctx context.Context, in *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	user, err := s.Storage.GetUser(in.Username)
	if err != nil {
		return nil, err
	}

	if err := auth.ComparePassword(in.Password, user.PasswordHash); err != nil {
		return nil, err
	}

	jwtToken, err := auth.GenerateJWT(int64(user.ID))
	if err != nil {
		return nil, err
	}

	decodedEncryptionKey, err := encryption.Decrypt(user.EncryptedMasterKey, []byte(constants.EncryptionKey))
	if err != nil {
		return nil, err
	}

	return &api.AuthenticateUserResponse{
		Token:         jwtToken,
		EncryptionKey: decodedEncryptionKey,
	}, nil
}
