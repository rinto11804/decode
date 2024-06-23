package answer

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "answer"

type Store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db}
}

func (s *Store) CreateAnswer(ctx context.Context, answer *types.AnswerCreateReq) (primitive.ObjectID, error) {
	res, err := s.db.Database(types.DBName).Collection(collName).InsertOne(ctx, answer)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}
