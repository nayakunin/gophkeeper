package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/nayakunin/gophkeeper/pkg/utils"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// GenerateKey generates a new AES-256 key
func (s *Service) GenerateKey() ([]byte, error) {
	return utils.GenerateRandom(2 * aes.BlockSize)
}

// Encrypt string to base64 crypto using AES GCM
func (s *Service) Encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("unable to create new cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("unable to create new GCM: %w", err)
	}

	nonce, err := utils.GenerateRandom(aesgcm.NonceSize())
	if err != nil {
		return "", fmt.Errorf("unable to generate nonce: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(text), nil)
	return fmt.Sprintf("%x", append(nonce, ciphertext...)), nil
}

// Decrypt from base64 to decrypted string
func (s *Service) Decrypt(text string, key []byte) (string, error) {
	decoded, err := utils.DecodeHex(text)
	if err != nil {
		return "", fmt.Errorf("unable to decode hex: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("unable to create new cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("unable to create new GCM: %w", err)
	}

	nonceSize := aesgcm.NonceSize()
	if len(decoded) < nonceSize {
		return "", fmt.Errorf("invalid nonce size")
	}

	nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("unable to open cipher: %w", err)
	}

	return string(plaintext), nil
}
