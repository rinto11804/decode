package main

import (
	"decode/room"
	"decode/task"
	"decode/user"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	addr string
	db   *mongo.Client
}

func NewAPIServer(addr string, db *mongo.Client) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	api := e.Group("/api/v1")

	api.GET("/health-check", healthCheck)

	userStore := user.NewStore(s.db)
	userService := user.NewService(userStore)
	userService.RegisterRoutes(api)

	roomStore := room.NewStore(s.db)
	roomService := room.NewService(roomStore)
	roomService.RegisterRoutes(api)

	taskStore := task.NewStore(s.db)
	taskService := task.NewService(taskStore)
	taskService.RegisterRoutes(api)

	log.Fatal(e.Start(s.addr))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Server status": "Still alive",
	})
}
