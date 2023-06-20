package HMAC

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

var key = []byte{}

func setKey() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
}

func SignMessage(msg []byte) ([]byte, error) {
	setKey()
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error while hashng message: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func CheckSign(msg, sign []byte) (bool, error) {
	newSign, err := SignMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error checking signature while gettins signature message: %w", err)
	}

	same := hmac.Equal(newSign, sign)
	return same, err
}
