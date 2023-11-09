//go:generate mockgen -source=api.go -destination=mocks/service.go -package=mocks
package api

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/card/input"
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

// PrepareCardRequest prepares a request to add card data.
func (s *Service) PrepareCardRequest(data *input.ParseCardResult, encryptionKey []byte) (*api.AddBankCardDetailRequest, error) {
	encryptedNumber, err := s.encryption.Encrypt([]byte(data.Number), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt card number: %w", err)
	}
	encryptedExpiration, err := s.encryption.Encrypt([]byte(data.Expiration), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt card expiration date: %w", err)
	}
	encryptedCVC, err := s.encryption.Encrypt([]byte(data.Cvc), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt card CVC: %w", err)
	}

	return &api.AddBankCardDetailRequest{
		CardName:            data.Name,
		EncryptedCardNumber: encryptedNumber,
		EncryptedExpiryDate: encryptedExpiration,
		EncryptedCvc:        encryptedCVC,
		Description:         data.Description,
	}, nil
}
