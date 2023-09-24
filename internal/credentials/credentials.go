package credentials

import "github.com/99designs/keyring"

type store struct {
	ring keyring.Keyring
}

func newStore() *store {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: "gophkeeper",
	})
	if err != nil {
		panic(err)
	}

	return &store{
		ring: ring,
	}
}

func (s *store) Set(key, value string) error {
	return s.ring.Set(keyring.Item{
		Key:  key,
		Data: []byte(value),
	})
}

func (s *store) Get(key string) (string, error) {
	item, err := s.ring.Get(key)
	if err != nil {
		return "", err
	}

	return string(item.Data), nil
}

func (s *store) Delete(key string) error {
	return s.ring.Remove(key)
}

var Store = newStore()
