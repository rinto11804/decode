package room

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Store) GetRoomByID(ctx context.Context, id string) (*types.RoomModel, error) {
	var room types.RoomModel
	roomID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := s.db.Database(types.DBName).Collection(collName).FindOne(ctx, bson.M{"_id": roomID})
	if err := res.Decode(&room); err != nil {
		return nil, err
	}
	return &room, nil
}
