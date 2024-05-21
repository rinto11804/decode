package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	ErrCourseCreateBodyParse = errors.New("error in parsing create course body")
)

type Service struct {
	store Storage
}

func NewService(store Storage) *Service {
	return &Service{store: store}
}

func (s *Service) RegisterRoutes(group *echo.Group) {
	group.POST("/course", s.handleCreateCourse)
}

func (s *Service) handleCreateCourse(c echo.Context) error {
	user, ok := c.Get("user").(UserModel)
	if !ok {
		return ErrUserNotFound
	}
	var courseInput CourseCreateBody
	if err := c.Bind(&courseInput); err != nil {
		return err
	}
	course := &CourseModel{
		Title:       courseInput.Title,
		Description: courseInput.Description,
		Author:      courseInput.Author,
		UserID:      user.ID,
		CreatedAt:   time.Time{},
	}
	if err := s.store.CreateCourse(course); err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"msg":    "course created successfully",
		"course": courseInput,
	})
}
