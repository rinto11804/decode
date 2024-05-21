package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	api := e.Group("/api/v1")

	api.GET("/health-check", healthCheck)

	cfg := LoadConfig()
	store, err := NewPostgresStore(cfg)
	if err != nil {
		log.Fatal(err)
	}

	service := NewService(store)
	service.RegisterRoutes(api)

	log.Fatal(e.Start(s.addr))
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Server status": "Still alive",
	})
}
