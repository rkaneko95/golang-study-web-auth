package encryption

import (
	"encoding/base64"
	"fmt"
)

func Base64Encoded(msg string) string {
	encoded := base64.URLEncoding.EncodeToString([]byte(msg))
	return encoded
}

func Base64Decoded(encoded string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("error in base64Decoded while decoding string: %w", err)
	}

	return string(decoded), nil
}
