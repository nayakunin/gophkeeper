package credentials

import (
	"github.com/99designs/keyring"
)

type Service struct {
	ring keyring.Keyring
}

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

func (s *Service) Set(key, value string) error {
	return s.ring.Set(keyring.Item{
		Key:  key,
		Data: []byte(value),
	})
}

func (s *Service) Get(key string) (string, error) {
	item, err := s.ring.Get(key)
	if err != nil {
		return "", err
	}

	return string(item.Data), nil
}

func (s *Service) Delete(key string) error {
	return s.ring.Remove(key)
}
