package joinlist

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collName = "room_joinlist"

type Store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db: db.Database(types.DBName)}
}

func (s *Store) CreateJoinlist(ctx context.Context, join types.RoomJoinlistReq) (primitive.ObjectID, error) {
	res, err := s.db.Collection(collName).InsertOne(ctx, join)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
