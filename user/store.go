package user

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collName = "user"

type Store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db}
}

func (s *Store) CreateUser(ctx context.Context, user *types.UserCreateReq) (primitive.ObjectID, error) {
	return primitive.NilObjectID, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*types.UserModel, error) {
	return nil, nil
}

func (s *Store) GetUserByID(ctx context.Context, id primitive.ObjectID) (*types.UserModel, error) {
	return nil, nil
}
