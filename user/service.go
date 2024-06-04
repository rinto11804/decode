package user

import (
	"decode/types"

	"github.com/labstack/echo/v4"
)

type Service struct {
	store types.UserStore
}

func NewService(store types.UserStore) *Service {
	return &Service{store}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/user", s.handleCreateUser)
}

func (s *Service) handleCreateUser(c echo.Context) error {
	return nil
}
