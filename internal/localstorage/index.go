package localstorage

import (
	"fmt"
)

type CredentialsService interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}

type Storage struct {
	credentialsService CredentialsService
}

func NewStorage(credentialsService CredentialsService) *Storage {
	return &Storage{
		credentialsService: credentialsService,
	}
}

func (s *Storage) SaveCredentials(token string, encryptionKey []byte) error {
	if err := s.credentialsService.Set("token", []byte(token)); err != nil {
		return fmt.Errorf("unable to save token: %w", err)
	}

	if err := s.credentialsService.Set("encryptionKey", encryptionKey); err != nil {
		return fmt.Errorf("unable to save encryption key: %w", err)
	}

	return nil
}

func (s *Storage) GetCredentials() (token string, encryptionKey []byte, err error) {
	tokenBytes, err := s.credentialsService.Get("token")
	if err != nil {
		return "", nil, fmt.Errorf("unable to get token: %w", err)
	}

	encryptionKey, err = s.credentialsService.Get("encryptionKey")
	if err != nil {
		return "", nil, fmt.Errorf("unable to get encryption key: %w", err)
	}

	return string(tokenBytes), encryptionKey, nil
}

func (s *Storage) DeleteCredentials() error {
	if err := s.credentialsService.Delete("token"); err != nil {
		return fmt.Errorf("unable to delete token: %w", err)
	}

	if err := s.credentialsService.Delete("encryptionKey"); err != nil {
		return fmt.Errorf("unable to delete encryption key: %w", err)
	}

	return nil
}
