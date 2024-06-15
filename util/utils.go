package util

import (
	"decode/types"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetRole(role string) types.Role {
	if role == "admin" {
		return types.ADMIN
	}
	return types.USER
}

func SentErrResponse(c echo.Context, statusCode int, err error) error {
	return c.JSON(statusCode, echo.Map{
		"msg": err,
	})
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func IsValidPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return !(err != nil)
}
