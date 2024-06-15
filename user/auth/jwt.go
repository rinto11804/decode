package auth

import (
	"decode/types"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrSigningMethodNotFound = errors.New("signed method not found")

func CreateJWTToken(secret string, id string, role types.Role) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"iss": "decode",
		"aud": role,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ValidateToken(token string, secret string) {
	jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrSigningMethodNotFound
		}
		return []byte(secret), nil
	})
}
