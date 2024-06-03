package room

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collName = "room"

type Store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db}
}

func (s *Store) CreateRoom(ctx context.Context, room *types.RoomCreateReq) (primitive.ObjectID, error) {
	res, err := s.db.Database(types.DBName).Collection(collName).InsertOne(ctx, room)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
