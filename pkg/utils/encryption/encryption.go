package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Encrypt string to base64 crypto using AES GCM
func Encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(text), nil)
	return fmt.Sprintf("%x", append(nonce, ciphertext...)), nil
}

// Decrypt from base64 to decrypted string
func Decrypt(text string, key []byte) (string, error) {
	decoded, err := decodeHex(text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesgcm.NonceSize()
	if len(decoded) < nonceSize {
		return "", fmt.Errorf("invalid nonce size")
	}

	nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]
	plaintext, err := aesgcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func decodeHex(text string) (string, error) {
	decoded := make([]byte, len(text)/2)
	_, err := fmt.Sscanf(text, "%x", &decoded)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
