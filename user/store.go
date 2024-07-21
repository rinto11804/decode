package user

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collName = "user"

type Store struct {
	db *mongo.Database
}

func NewStore(client *mongo.Client) *Store {
	client.Database(types.DBName).Collection(collName).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.E{Key: "email", Value: 1},
			Options: options.Index().SetUnique(true),
		},
	)

	return &Store{db: client.Database(types.DBName)}
}

func (s *Store) CreateUser(ctx context.Context, user *types.UserCreateReq) (primitive.ObjectID, error) {
	coll := s.db.Collection(collName)
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*types.UserModel, error) {
	var user *types.UserModel
	coll := s.db.Collection(collName)
	res := coll.FindOne(ctx, bson.M{"email": email})
	if err := res.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByID(ctx context.Context, id string) (*types.UserModel, error) {
	var user *types.UserModel
	coll := s.db.Collection(collName)
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := coll.FindOne(ctx, bson.M{"_id": ID})
	if err := res.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// func (s *Store) GetUserByIDs(ctx context.Context, ids []primitive.ObjectID) ([]types.UserModel, error) {
// 	var users []types.UserModel
// 	coll := s.db.Collection(collName)
// 	cursor, err := coll.Find(ctx, bson.M{"_id": ids}, &options.FindOptions{
// 		Projection: bson.M{"_id": 1, "name": 1},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := cursor.All(ctx, &users); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
