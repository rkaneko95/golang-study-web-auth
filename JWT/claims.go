package JWT

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("tokens has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid sesion ID")
	}

	return nil
}
