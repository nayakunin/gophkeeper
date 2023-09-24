package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

// GenerateKey generates a new AES-256 key
func generateRandom(size int) ([]byte, error) {
	rand.Seed(time.Now().UnixNano())

	var result []byte
	for len(result) < size {
		// Generate a Unicode code point that is a valid character.
		r := rune(rand.Intn(0x10FFFF))
		if !utf8.ValidRune(r) {
			continue
		}

		// Convert it to UTF-8 encoding.
		var buf [4]byte
		n := utf8.EncodeRune(buf[:], r)

		// Ensure the resulting sequence doesn't exceed the desired size.
		if len(result)+n > size {
			continue
		}

		result = append(result, buf[:n]...)
	}

	return result, nil
}

// GenerateKey generates a new AES-256 key
func GenerateKey() ([]byte, error) {
	return generateRandom(2 * aes.BlockSize)
}

// Encrypt string to base64 crypto using AES GCM
func Encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("unable to create new cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("unable to create new GCM: %w", err)
	}

	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		return "", fmt.Errorf("unable to generate nonce: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(text), nil)
	return fmt.Sprintf("%x", append(nonce, ciphertext...)), nil
}

// Decrypt from base64 to decrypted string
func Decrypt(text string, key []byte) (string, error) {
	decoded, err := decodeHex(text)
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

func decodeHex(text string) ([]byte, error) {
	decoded := make([]byte, hex.DecodedLen(len(text)))
	_, err := hex.Decode(decoded, []byte(text))
	if err != nil {
		return nil, fmt.Errorf("unable to decode hex: %w", err)
	}

	return decoded, nil
}
