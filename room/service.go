package room

import "github.com/labstack/echo/v4"

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/room", s.handleCreateRoom)
}

func (s *Service) handleCreateRoom(c echo.Context) error {
	return nil
}
