package task

import (
	"context"
	"decode/types"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	ErrTaskCreateBodyParse = errors.New("error in parsing create task body")
	ErrInvalidRoomID       = errors.New("invalid room id provided")
)

type Service struct {
	store     types.TaskStore
	roomStore types.RoomStore
}

type taskCreateBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Handler     string `json:"handler"`
	Body        string `json:"body"`
	RoomID      string `json:"room_id"`
}

func NewService(store types.TaskStore, roomSore types.RoomStore) *Service {
	return &Service{store, roomSore}
}

func (s *Service) RegisterRoutes(group *echo.Group) {
	group.POST("/task", s.handleCreateTask)
	group.GET("/task/:roomId", s.handleGetAllTaskByRoomID)
	group.GET("/task/:taskId", s.GetTaskDetails)
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
		Handler:     taskInput.Handler,
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

func (s *Service) handleGetAllTaskByRoomID(c echo.Context) error {
	roomID := c.Param("roomId")
	if roomID == "" {
		return ErrInvalidRoomID
	}

	_, err := s.roomStore.GetRoomByID(context.Background(), roomID)
	if err != nil {
		return err
	}

	tasks, err := s.store.GetAllTaskByRoomID(context.Background(), roomID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg":   "task found",
		"tasks": tasks,
	})
}

// @Summary			Get Task Details
// @Description		get task details by taskId
// @Tags			Task
// @ID				get-task-details-by-id
// @Accept			json
// @Produce			json
// @Param			taskId	path		string	true	"Task ID"
// @Success			200		{object}	types.Response[types.TaskModel]
// @Router			/decode/task/{taskId} [get]
// @Security		Bearer
func (s *Service) GetTaskDetails(c echo.Context) error {
	taskID := c.Param("taskId")
	fmt.Println(taskID)
	if taskID == "" {
		return ErrInvalidRoomID
	}

	task, err := s.store.GetTaskByID(context.Background(), taskID)
	fmt.Println(task)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "task found",
		"task": task,
	})
}
