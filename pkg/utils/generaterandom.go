package utils

import (
	"math/rand"
	"time"
	"unicode/utf8"
)

// GenerateRandom generates a new AES-256 key
func GenerateRandom(size int) ([]byte, error) {
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
