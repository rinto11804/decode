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
	ErrInvalidRoomID       = errors.New("invalid room id provided")
)

type Service struct {
	store     types.TaskStore
	roomStore types.RoomStore
}

type TaskCreateBody struct {
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
	group.GET("/task/room/:roomId", s.handleGetAllTaskByRoomID)
	group.GET("/task/:taskId", s.GetTaskDetails)
}

// @Summary			Create Task
// @Description		create task for a room
// @Tags			Task
// @ID				create-task
// @Accept			json
// @Produce			json
// @Param			taskInput	body	TaskCreateBody	true	"create task request body"
// @Success			200		{object}	types.Response[string] 	"taskId"
// @Router			/decode/task [post]
// @Security		Bearer
func (s *Service) handleCreateTask(c echo.Context) error {
	var taskInput TaskCreateBody
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

	return c.JSON(http.StatusCreated, types.Response[string]{
		Msg:  "task created",
		Data: id.Hex(),
	})
}

// @Summary			Get all task in a room
// @Description		get all task by roomId
// @Tags			Task
// @ID				get-all-task-by-roomId
// @Accept			json
// @Produce			json
// @Param			roomId	path	string	true	"roomId"
// @Success			200		{object}	types.Response[[]types.ProjectedTask] 	"tasks"
// @Router			/decode/task/room/{roomId} [get]
// @Security		Bearer
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

	return c.JSON(http.StatusOK, types.Response[[]types.ProjectedTask]{
		Msg:  "tasks found",
		Data: tasks,
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
	if taskID == "" {
		return ErrInvalidRoomID
	}

	task, err := s.store.GetTaskByID(context.Background(), taskID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response[types.TaskModel]{
		Msg:  "task found",
		Data: *task,
	})
}
