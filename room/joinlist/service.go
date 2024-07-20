package joinlist

import (
	"context"
	"decode/types"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	store     types.JoinlistStore
	roomStore types.RoomStore
}

func NewService(store types.JoinlistStore, roomStore types.RoomStore) *Service {
	return &Service{store, roomStore}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/join/:roomId", s.handleJoinRoom)
}

// @Summary			Join Room
// @Description		join the room with roomId
// @Tags			Room
// @ID				join-room
// @Accept			json
// @Produce			json
// @Param			roomId	path		string	true	"Room ID"
// @Success			200		{object}	types.Response[string]
// @Router			/decode/join/{roomId} [post]
// @Security		Bearer
func (s *Service) handleJoinRoom(c echo.Context) error {
	roomID := c.Param("roomId")

	user := c.Get("user").(types.User)
	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	room, err := s.roomStore.GetRoomByID(context.Background(), roomID)
	if err != nil {
		return err
	}
	joinlist := types.RoomJoinlistReq{
		UserID:    userID,
		RoomID:    room.ID,
		CreatedAt: time.Now(),
	}

	joinId, err := s.store.CreateJoinlist(context.Background(), joinlist)
	return c.JSON(http.StatusCreated, types.Response[string]{
		Msg:  "joined room",
		Data: joinId.Hex(),
	})
}
