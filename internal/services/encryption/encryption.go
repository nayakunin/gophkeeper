package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/nayakunin/gophkeeper/pkg/utils"
)

// Service is a struct of the encryption service.
type Service struct{}

// NewService returns a new Service.
func NewService() *Service {
	return &Service{}
}

// GenerateKey generates a new AES-256 key
func (s *Service) GenerateKey() ([]byte, error) {
	return utils.GenerateRandom(2 * aes.BlockSize)
}

// Encrypt string to base64 crypto using AES GCM
func (s *Service) Encrypt(text, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("unable to create new cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("unable to create new GCM: %w", err)
	}

	nonce, err := utils.GenerateRandom(aesgcm.NonceSize())
	if err != nil {
		return nil, fmt.Errorf("unable to generate nonce: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, text, nil)

	return append(nonce, ciphertext...), nil
}

// Decrypt from base64 to decrypted string
func (s *Service) Decrypt(text, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("unable to create new cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("unable to create new GCM: %w", err)
	}

	nonceSize := aesgcm.NonceSize()
	if len(text) < nonceSize {
		return nil, fmt.Errorf("invalid nonce size")
	}

	nonce, ciphertext := text[:nonceSize], text[nonceSize:]
	textBytes, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to open cipher: %w", err)
	}

	return textBytes, nil
}
