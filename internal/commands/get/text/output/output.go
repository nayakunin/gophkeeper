//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"fmt"

	generated "github.com/nayakunin/gophkeeper/proto"
)

type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

type Service struct {
	encryption Encryption
}

func NewService(encryption Encryption) *Service {
	return &Service{
		encryption: encryption,
	}
}

type TextResult struct {
	Description string `json:"description"`
	Text        string `json:"text"`
}

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
