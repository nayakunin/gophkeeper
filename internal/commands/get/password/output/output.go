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

type PasswordResult struct {
	ServiceName string `json:"service_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

func (s *Service) MakeResponse(response *generated.GetLoginPasswordPairsResponse, encryptionKey []byte) ([]PasswordResult, error) {
	results := make([]PasswordResult, len(response.GetLoginPasswordPairs()))
	for i, pair := range response.GetLoginPasswordPairs() {
		password, err := s.encryption.Decrypt(pair.GetEncryptedPassword(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt password: %w", err)
		}
		results[i] = PasswordResult{
			ServiceName: pair.GetServiceName(),
			Login:       pair.GetLogin(),
			Password:    string(password),
			Description: pair.GetDescription(),
		}
	}

	return results, nil
}
