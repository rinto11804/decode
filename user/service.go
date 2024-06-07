package user

import (
	"context"
	"decode/types"
	"decode/util"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	ErrUserAlreadyExist = errors.New("user with email already exist")
	ErrUserNotFound     = errors.New("user not found")
)

type Service struct {
	store types.UserStore
}

type RegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewService(store types.UserStore) *Service {
	return &Service{store}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/register", s.Register)
}

func (s *Service) Register(c echo.Context) error {
	var userInput RegisterReq

	if err := c.Bind(&userInput); err != nil {
		return err
	}

	u, _ := s.store.GetUserByEmail(context.Background(), userInput.Email)
	if u != nil {
		return ErrUserAlreadyExist
	}

	hash, err := util.HashPassword(userInput.Password)
	if err != nil {
		return err
	}

	user := &types.UserCreateReq{
		Name:      userInput.Name,
		Email:     userInput.Email,
		Password:  hash,
		Role:      types.USER,
		CreatedAt: time.Now(),
	}

	userId, err := s.store.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"msg":     "user created successfully",
		"user_id": userId.Hex(),
	})
}
