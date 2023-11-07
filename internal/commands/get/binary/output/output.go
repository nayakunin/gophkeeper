//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"fmt"

	generated "github.com/nayakunin/gophkeeper/proto"
)

type Encryption interface {
	Decrypt(data []byte, key []byte) ([]byte, error)
}

type Service struct {
	encryption Encryption
}

func NewService(encryption Encryption) *Service {
	return &Service{
		encryption: encryption,
	}
}

type BinaryResult struct {
	Data        []byte `json:"data"`
	Description string `json:"description"`
}

func (s *Service) MakeBinaryResponse(response *generated.GetBinaryDataResponse, encryptionKey []byte) ([]BinaryResult, error) {
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
