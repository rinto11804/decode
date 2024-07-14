package user

import (
	"context"
	"decode/config"
	"decode/types"
	"decode/util"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	ErrUserAlreadyExist = errors.New("user with email already exist")
	ErrUserNotFound     = errors.New("user not found")
)

type Service struct {
	cfg   *config.Config
	store types.UserStore
}

type RegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewService(store types.UserStore, config *config.Config) *Service {
	return &Service{config, store}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/register/:role", s.handleRegister)
	api.POST("/login", s.handleLogin)
}

func (s *Service) handleRegister(c echo.Context) error {
	var userInput RegisterReq
	role := util.GetRole(c.Param("role"))

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
		Role:      role,
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

func (s *Service) handleLogin(c echo.Context) error {
	var loginInput LoginReq

	if err := c.Bind(&loginInput); err != nil {
		return err
	}

	user, err := s.store.GetUserByEmail(context.Background(), loginInput.Email)
	if err != nil {
		return ErrUserNotFound
	}

	if !util.IsValidPassword(loginInput.Password, user.Password) {
		return ErrUserNotFound
	}

	token, err := s.CreateJWTToken(user.ID.Hex(), user.Role)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: false,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  user,
	})
}
func (s Service) CreateJWTToken(id string, role types.Role) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"iss": "decode",
		"aud": role,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString([]byte(s.cfg.JwtSecret))
}
