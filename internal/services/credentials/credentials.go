package credentials

import (
	"github.com/99designs/keyring"
)

// Service is a struct of the grpc.
type Service struct {
	ring keyring.Keyring
}

// NewService returns a new Service.
func NewService() *Service {
	ring, err := keyring.Open(keyring.Config{
		ServiceName:                    "gophkeeper",
		KeychainName:                   "gophkeeper",
		KeychainTrustApplication:       true,
		KeychainAccessibleWhenUnlocked: true,
	})
	if err != nil {
		panic(err)
	}

	return &Service{
		ring: ring,
	}
}

// Set sets a value for a key.
func (s *Service) Set(key string, value []byte) error {
	return s.ring.Set(keyring.Item{
		Key:  key,
		Data: value,
	})
}

// Get returns a value for a key.
func (s *Service) Get(key string) ([]byte, error) {
	item, err := s.ring.Get(key)
	if err != nil {
		return nil, err
	}

	return item.Data, nil
}

// Delete deletes a value for a key.
func (s *Service) Delete(key string) error {
	return s.ring.Remove(key)
}
