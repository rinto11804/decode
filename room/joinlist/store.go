package joinlist

import (
	"context"
	"decode/types"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Store) GrantPoint(ctx context.Context, id string, point int) error {
	s.db.Collection(collName).UpdateByID(ctx, id, nil)
	return types.ErrNotImplemented
}

func (s *Store) GenerateLeaderBoard(ctx context.Context, roomID string) ([]types.LeaderBoardItems, error) {
	var leaderboardItems []types.LeaderBoardItems
	ID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, err
	}
	fmt.Println(ID)
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"room_id": ID}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "user",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user",
		}}},
		{{Key: "$unwind", Value: "$user"}},
		{{Key: "$group", Value: bson.M{
			"_id":      "$_id",
			"user_id":  bson.M{"$first": "$user_id"},
			"username": bson.M{"$first": "$user.name"},
			"points":   bson.M{"$sum": "$points"},
		}}},
		{{Key: "$sort", Value: bson.M{"points": -1}}},
	}

	cursor, err := s.db.Collection(collName).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &leaderboardItems); err != nil {
		return nil, err
	}
	fmt.Println(leaderboardItems)
	return leaderboardItems, nil
}
