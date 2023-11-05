//go:generate mockgen -source=api.go -destination=mocks/service.go -package=mocks
package api

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/password/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
}

type Service struct {
	encryption Encryption
}

func NewService(encryption Encryption) *Service {
	return &Service{
		encryption: encryption,
	}
}

func (s *Service) PreparePasswordRequest(result *input.ParsePasswordResult, encryptionKey []byte) (*generated.AddLoginPasswordPairRequest, error) {
	encryptedPassword, err := s.encryption.Encrypt([]byte(result.Password), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt password: %w", err)
	}

	return &generated.AddLoginPasswordPairRequest{
		ServiceName:       result.ServiceName,
		Login:             result.Login,
		EncryptedPassword: encryptedPassword,
		Description:       result.Description,
	}, nil
}
