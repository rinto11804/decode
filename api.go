package main

import (
	"decode/config"
	"decode/room"
	"decode/task"
	"decode/task/answer"
	"decode/user"
	"decode/user/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	config *config.Config
	db     *mongo.Client
}

func NewAPIServer(config *config.Config, db *mongo.Client) *APIServer {
	return &APIServer{
		config: config,
		db:     db,
	}
}

func (s *APIServer) Run() error {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	api := e.Group("/api/v1")

	api.GET("/healthcheck", healthCheck)

	userStore := user.NewStore(s.db)
	userService := user.NewService(userStore, s.config)
	userService.RegisterRoutes(api)

	authService := auth.NewService(s.config, userStore)

	protected := api.Group("/decode", authService.AuthMiddleware)
	protected.GET("/authstatus", authCheck)

	roomStore := room.NewStore(s.db)
	roomService := room.NewService(roomStore)
	roomService.RegisterRoutes(protected)

	taskStore := task.NewStore(s.db)
	taskService := task.NewService(taskStore, roomStore)
	taskService.RegisterRoutes(protected)

	answerStore := answer.NewStore(s.db)
	answerService := answer.NewService(answerStore)
	answerService.RegisterRoutes(protected)

	if err := e.Start(s.config.Port); err != nil {
		return err
	}
	return nil
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Server status": "Still alive",
	})
}

func authCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"auth": "user auth success",
	})
}
