package main

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

var (
	ErrEventCreate = errors.New("creating event failed")
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterRoutes(group *echo.Group) {
	group.POST("/event", s.handleCreateEvent)

}

func (s *Service) handleCreateEvent(c echo.Context) error {
	var eventInput EventInput
	if err := c.Bind(&eventInput); err != nil {
		return err
	}
	fmt.Println(eventInput)
	return nil
}
