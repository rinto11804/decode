package room

import (
	"context"
	"decode/types"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	store types.RoomStore
}

type RoomCreateBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RoomCreateRes struct {
	Msg    string `json:"msg"`
	RoomID string `json:"room_id"`
}

func NewService(store types.RoomStore) *Service {
	return &Service{store}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/room", s.handleCreateRoom)
}

func (s *Service) handleCreateRoom(c echo.Context) error {
	user := c.Get("user").(types.User)
	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	var roomInput RoomCreateBody
	if err := c.Bind(&roomInput); err != nil {
		return err
	}

	room := &types.RoomCreateReq{
		Title:       roomInput.Title,
		Description: roomInput.Description,
		UserID:      userID,
		CreatedAt:   time.Now(),
	}

	roomID, err := s.store.CreateRoom(context.Background(), room)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"msg":     "room created successfully",
		"room_id": roomID.Hex(),
	})
}
