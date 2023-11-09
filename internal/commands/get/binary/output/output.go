//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"fmt"

	generated "github.com/nayakunin/gophkeeper/proto"
)

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Decrypt(data []byte, key []byte) ([]byte, error)
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

// BinaryResult is a result of getting binary data.
type BinaryResult struct {
	Data        []byte `json:"data"`
	Description string `json:"description"`
}

// MakeResponse prepares a response to get binary data.
func (s *Service) MakeResponse(response *generated.GetBinaryDataResponse, encryptionKey []byte) ([]BinaryResult, error) {
	results := make([]BinaryResult, len(response.GetBinaryData()))
	for i, pair := range response.GetBinaryData() {
		data, err := s.encryption.Decrypt(pair.GetEncryptedData(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt data: %w", err)
		}
		results[i] = BinaryResult{
			Data:        data,
			Description: pair.GetDescription(),
		}
	}

	return results, nil
}
