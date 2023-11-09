//go:generate mockgen -source=api.go -destination=mocks/service.go -package=mocks
package api

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	api "github.com/nayakunin/gophkeeper/proto"
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

// PrepareTextRequest prepares a request to add text data.
func (s *Service) PrepareTextRequest(result *input.ParseTextResult, encryptionKey []byte) (*api.AddTextDataRequest, error) {
	encryptedText, err := s.encryption.Encrypt([]byte(result.Text), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt text: %w", err)
	}

	return &api.AddTextDataRequest{
		EncryptedText: encryptedText,
		Description:   result.Description,
	}, nil
}
