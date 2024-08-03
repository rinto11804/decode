package types

import (
	"context"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DBName = "decode"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrNotImplemented = errors.New("todo:service not implemented")
)

type Role string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
	GUEST Role = "GUEST"
)

type ResponseErr struct {
	Msg   string `json:"message"`
	Error string `json:"error"`
}

type Response[T any] struct {
	Msg  string `json:"message"`
	Data T      `json:"data"`
}

type User struct {
	ID   string `json:"id"`
	Role Role   `json:"role"`
}

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	Role      Role               `bson:"role" json:"role"`
	Points    int                `bson:"points" json:"points"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type RoomModel struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type RoomJoinlistModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	RoomID    primitive.ObjectID `bson:"room_id" json:"room_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Points    int                `bson:"points" json:"points"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type TaskModel struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Handler     string             `bson:"handler" json:"handler,omitempty"`
	Body        string             `bson:"body" json:"body,omitempty"`
	Point       int                `bson:"point" json:"point,omitempty"`
	RoomID      primitive.ObjectID `bson:"room_id" json:"room_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at,omitempty"`
}

type AnswerModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Body      string             `bson:"body" json:"body"`
	TaskID    primitive.ObjectID `bson:"task_id" json:"task_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	IsCorrect bool               `bson:"is_correct" json:"is_correct"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type UserStore interface {
	CreateUser(context.Context, *UserCreateReq) (primitive.ObjectID, error)
	GetUserByID(ctx context.Context, id string) (*UserModel, error)
	GetUserByEmail(ctx context.Context, email string) (*UserModel, error)
}

type TaskStore interface {
	CreateTask(context.Context, *TaskCreateReq) (primitive.ObjectID, error)
	GetTaskByID(ctx context.Context, id string) (*TaskModel, error)
	GetAllTaskByRoomID(ctx context.Context, roomID string) ([]ProjectedTask, error)
}

type RoomStore interface {
	CreateRoom(context.Context, *RoomCreateReq) (primitive.ObjectID, error)
	GetRoomByID(ctx context.Context, id string) (*RoomModel, error)
	GetAllRoomByUserID(ctx context.Context, userID primitive.ObjectID) ([]RoomModel, error)
}

type JoinlistStore interface {
	CreateJoinlist(context.Context, RoomJoinlistReq) (primitive.ObjectID, error)
	GenerateLeaderBoard(ctx context.Context, roomID string) ([]LeaderBoardItems, error)
}

type AnswerStore interface {
	CreateAnswer(context.Context, *AnswerCreateReq) (primitive.ObjectID, error)
	MarkAsCorrect(ctx context.Context, id string) error
}

type SubRoute = func(api *echo.Group)

type UserCreateReq struct {
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"_"`
	Role      Role      `bson:"role" json:"role"`
	Points    int       `bson:"points" json:"points"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type RoomCreateReq struct {
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type RoomJoinlistReq struct {
	RoomID    primitive.ObjectID `bson:"room_id" json:"room_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Points    int                `bson:"points" json:"points"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type TaskCreateReq struct {
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Handler     string             `bson:"handler" json:"handler"`
	Body        string             `bson:"body" json:"body"`
	RoomID      primitive.ObjectID `bson:"room_id" json:"room_id"`
	Point       int                `bson:"point" json:"point,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type AnswerCreateReq struct {
	Body      string             `bson:"body" json:"body"`
	TaskID    primitive.ObjectID `bson:"task_id" json:"task_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type ProjectedTask struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Handler   string             `bson:"handler" json:"handler"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type LeaderBoardItems struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	UserName string             `bson:"username" json:"username"`
	Points   int                `bson:"points" json:"points"`
}
