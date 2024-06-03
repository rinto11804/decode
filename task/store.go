package task

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collName = "task"

type Store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db}
}

func (s Store) CreateTask(ctx context.Context, task *types.TaskCreateReq) (primitive.ObjectID, error) {
	res, err := s.db.Database(types.DBName).Collection(collName).InsertOne(context.Background(), task)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}
