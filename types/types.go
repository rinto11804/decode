package types

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DBName = "decode"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Role = string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)

type User struct {
	ID   string `json:"id"`
	Role Role   `json:"role"`
}

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"_"`
	Role      Role               `bson:"role" json:"role"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type RoomModel struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type TaskModel struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	RoomID      primitive.ObjectID `bson:"room_id" json:"room_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type TaskStore interface {
	CreateTask(context.Context, *TaskCreateReq) (primitive.ObjectID, error)
}

type RoomStore interface {
	CreateRoom(context.Context, *RoomCreateReq) (primitive.ObjectID, error)
}

type RoomCreateReq struct {
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type TaskCreateReq struct {
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	RoomID      primitive.ObjectID `bson:"room_id" json:"room_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
