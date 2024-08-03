package answer

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "answer"

type Store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db: db.Database(types.DBName)}
}

func (s *Store) CreateAnswer(ctx context.Context, answer *types.AnswerCreateReq) (primitive.ObjectID, error) {
	res, err := s.db.Collection(collName).InsertOne(ctx, answer)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (s *Store) GetTaskByID(ctx context.Context, id string) (*types.TaskModel, error) {
	var task types.TaskModel
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := s.db.Collection(collName).FindOne(ctx, bson.M{"_id": taskID})

	if err := res.Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *Store) MarkAsCorrect(ctx context.Context, id string) error {
	return types.ErrNotImplemented
}
