//go:generate mockgen -source=api.go -destination=mocks/service.go -package=mocks
package api

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	api "github.com/nayakunin/gophkeeper/proto"
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
