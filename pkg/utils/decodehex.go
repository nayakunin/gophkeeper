package utils

import (
	"encoding/hex"
	"fmt"
)

func DecodeHex(text string) ([]byte, error) {
	decoded := make([]byte, hex.DecodedLen(len(text)))
	_, err := hex.Decode(decoded, []byte(text))
	if err != nil {
		return nil, fmt.Errorf("unable to decode hex: %w", err)
	}

	return decoded, nil
}
