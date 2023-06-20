package JWT

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"io"
	"time"
)

type key struct {
	key     []byte
	created time.Time
}

var currentKid = ""
var keys = map[string]key{}

// Periodically trigger
func GenerateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("error in generateNewKey while generating key: %w", err)
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error in generateNewKey while generating kid: %w", err)
	}

	keys[uid.String()] = key{
		key:     newKey,
		created: time.Now(),
	}
	currentKid = uid.String()

	return nil
}

func CreateToken(c *UserClaims) (string, error) {
	//{JWT standard fields}.{Your fields}.Signature
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(keys[currentKid].key)
	if err != nil {
		return "", fmt.Errorf("error in CreateToken when signing token: %w", err)
	}

	return signedToken, nil
}

func ValidateToken(token string) (*UserClaims, error) {
	return parseToken(token)
}

func parseToken(token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Check witch key you are using
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}

		// Rotating => change periodically
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key ID")
		}

		k, ok := keys[kid]
		if !ok {
			return nil, fmt.Errorf("invalid key ID")
		}

		return k.key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("error in parseToken, token is not valid")
	}

	return t.Claims.(*UserClaims), nil
}
