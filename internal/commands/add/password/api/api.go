//go:generate mockgen -source=api.go -destination=mocks/service.go -package=mocks
package api

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/password/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
}

// Service is an interface for interacting with the API.
type Service struct {
	encryption Encryption
}

// NewService creates a new instance of Service.
func NewService(encryption Encryption) *Service {
	return &Service{
		encryption: encryption,
	}
}

// PreparePasswordRequest prepares a request to add password data.
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
