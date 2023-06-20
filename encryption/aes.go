package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)

func EnDecode(key, input []byte) ([]byte, error) {
	buff := &bytes.Buffer{}
	sw, err := EncryptWriter(buff, key)
	_, err = sw.Write(input)
	if err != nil {
		return nil, fmt.Errorf("error in enDecode while writing stream writer: %w", err)
	}

	output := buff.Bytes()

	return output, nil
}

func EncryptWriter(w io.Writer, key []byte) (io.Writer, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error in enDecode while creating Cipher: %w", err)
	}

	// initialization vector
	iv := make([]byte, aes.BlockSize)
	/*_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, fmt.Errorf("error in enDecode while generating iv: %w", err)
	}*/

	s := cipher.NewCTR(b, iv)

	return cipher.StreamWriter{
		S: s,
		W: w,
	}, nil
}
