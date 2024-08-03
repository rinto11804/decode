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
	api.GET("/room", s.handleGetAllRoomsByUserID)
}

// @Summary			Create Room
// @Description		create the room
// @Tags			Room
// @ID				create-room
// @Accept			json
// @Produce			json
// @Param			roomInput	body	RoomCreateBody	true	"room create request body"
// @Success			200		{object}	types.Response[string] 	"roomId"
// @Router			/decode/room [post]
// @Security		Bearer
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

	return c.JSON(http.StatusCreated, types.Response[string]{
		Msg:  "room created successfully",
		Data: roomID.Hex(),
	})
}

// @Summary			Get all room of current user
// @Description		get rooms joined by current login user
// @Tags			Room
// @ID				get-all-rooms-by-login-user
// @Accept			json
// @Produce			json
// @Success			200		{object}	types.Response[[]types.RoomModel] 	"rooms"
// @Router			/decode/room [get]
// @Security		Bearer
func (s *Service) handleGetAllRoomsByUserID(c echo.Context) error {
	user := c.Get("user").(types.User)
	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	rooms, err := s.store.GetAllRoomByUserID(context.Background(), userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.Response[[]types.RoomModel]{
		Msg:  "room found",
		Data: rooms,
	})
}
