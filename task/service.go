package task

import (
	"context"
	"decode/types"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	ErrTaskCreateBodyParse = errors.New("error in parsing create task body")
)

type Service struct {
	store     types.TaskStore
	roomStore types.RoomStore
}

type taskCreateBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	RoomID      string `json:"room_id"`
}

func NewService(store types.TaskStore) *Service {
	return &Service{store: store}
}

func (s *Service) RegisterRoutes(group *echo.Group) {
	group.POST("/task", s.handleCreateTask)
}

func (s *Service) handleCreateTask(c echo.Context) error {
	var taskInput taskCreateBody
	if err := c.Bind(&taskInput); err != nil {
		return err
	}

	room, err := s.roomStore.GetRoomByID(context.Background(), taskInput.RoomID)
	if err != nil {
		return err
	}

	task := &types.TaskCreateReq{
		Title:       taskInput.Title,
		Description: taskInput.Description,
		Body:        taskInput.Body,
		RoomID:      room.ID,
		CreatedAt:   time.Now(),
	}

	id, err := s.store.CreateTask(context.Background(), task)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"msg":     "task created successfully",
		"task_id": id.Hex(),
	})
}
