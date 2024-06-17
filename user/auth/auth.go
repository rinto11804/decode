package auth

import (
	"context"
	"decode/config"
	"decode/types"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	ErrSigningMethodNotFound = errors.New("signed method not found")
	ErrTokenNotValid         = errors.New("token is not valid")
)

type Service struct {
	cfg       *config.Config
	userStore types.UserStore
}

func NewService(cfg *config.Config, userStore types.UserStore) *Service {
	return &Service{cfg, userStore}
}

func (s Service) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := strings.Split(c.Request().Header.Get("Authorization"), " ")
		token, err := s.validateToken(header[1])
		if err != nil {
			return err
		}
		if !token.Valid {
			return ErrTokenNotValid
		}

		claims := (token.Claims).(jwt.MapClaims)
		id := claims["sub"].(string)

		user, err := s.userStore.GetUserByID(context.Background(), id)
		if err != nil {
			return echo.ErrUnauthorized
		}
		c.Set("user", user)
		return next(c)
	}
}

func (s Service) validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrSigningMethodNotFound
		}
		return []byte(s.cfg.JwtSecret), nil
	})
}
