package main

import (
	"decode/config"
	"decode/room"
	"decode/room/joinlist"
	"decode/task"
	"decode/task/answer"
	"decode/user"
	"decode/user/auth"
	"net/http"

	_ "decode/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	config   *config.Config
	dbClient *mongo.Client
}

func NewAPIServer(config *config.Config, client *mongo.Client) *APIServer {
	return &APIServer{
		config:   config,
		dbClient: client,
	}
}

func (s *APIServer) Run() error {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	api := e.Group("/api/v1")
	api.GET("/swagger/*", echoSwagger.WrapHandler)

	api.GET("/healthcheck", healthCheck)

	userStore := user.NewStore(s.dbClient)
	userService := user.NewService(userStore, s.config)
	userService.RegisterRoutes(api)

	authService := auth.NewService(s.config, userStore)

	protected := api.Group("/decode", authService.AuthMiddleware)
	protected.GET("/authstatus", authCheck)

	roomStore := room.NewStore(s.dbClient)
	roomService := room.NewService(roomStore)
	roomService.RegisterRoutes(protected)

	joinListStore := joinlist.NewStore(s.dbClient)
	joinListService := joinlist.NewService(joinListStore, roomStore)
	joinListService.RegisterRoutes(protected)

	taskStore := task.NewStore(s.dbClient)
	taskService := task.NewService(taskStore, roomStore)
	taskService.RegisterRoutes(protected)

	answerStore := answer.NewStore(s.dbClient)
	answerService := answer.NewService(answerStore, taskStore)
	answerService.RegisterRoutes(protected)

	if err := e.Start(s.config.Port); err != nil {
		return err
	}
	return nil
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"msg": "Still alive",
	})
}

func authCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"auth": "user auth success",
	})
}
