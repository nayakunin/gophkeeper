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

// PasswordResult is a result of getting a password.
type PasswordResult struct {
	ServiceName string `json:"service_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

// MakeResponse prepares a response to get password data.
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
