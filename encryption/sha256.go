package encryption

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func Sha() ([]byte, error) {
	f, err := os.Open("encryption/sample.txt")
	if err != nil {
		return nil, fmt.Errorf("error oping sample file: %w", err)
	}
	defer f.Close()

	h := sha256.New()

	_, err = io.Copy(h, f)
	if err != nil {
		return nil, fmt.Errorf("error copying: %w", err)
	}

	return h.Sum(nil), nil
}
