//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"fmt"

	generated "github.com/nayakunin/gophkeeper/proto"
)

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

// Service is an interface for preparing output.
type Service struct {
	encryption Encryption
}

// NewService creates a new instance of Service.
func NewService(encryption Encryption) *Service {
	return &Service{
		encryption: encryption,
	}
}

// CardResult is a result of getting card details.
type TextResult struct {
	Description string `json:"description"`
	Text        string `json:"text"`
}

// MakeResponse prepares a response to get card details.
func (s *Service) MakeResponse(response *generated.GetTextDataResponse, encryptionKey []byte) ([]TextResult, error) {
	results := make([]TextResult, len(response.GetTextData()))
	for i, textData := range response.GetTextData() {
		decryptedText, err := s.encryption.Decrypt(textData.GetEncryptedText(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt text: %w", err)
		}

		results[i] = TextResult{
			Description: textData.GetDescription(),
			Text:        string(decryptedText),
		}
	}

	return results, nil
}
