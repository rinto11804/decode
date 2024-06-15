package main

import (
	"decode/config"
	"decode/room"
	"decode/task"
	"decode/user"
	"decode/user/auth"
	"log"
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

func (s *APIServer) Run() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	api := e.Group("/api/v1")

	api.GET("/health-check", healthCheck)

	userStore := user.NewStore(s.db)
	userService := user.NewService(userStore, s.config)
	userService.RegisterRoutes(api)

	protected := api.Group("/decode", auth.AuthMiddleware)

	roomStore := room.NewStore(s.db)
	roomService := room.NewService(roomStore)
	roomService.RegisterRoutes(protected)

	taskStore := task.NewStore(s.db)
	taskService := task.NewService(taskStore)
	taskService.RegisterRoutes(protected)

	log.Println(s.config.Port)
	log.Fatal(e.Start(s.config.Port))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Server status": "Still alive",
	})
}
