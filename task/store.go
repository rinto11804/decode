package task

import (
	"context"
	"decode/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collName = "task"

type Store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Client) *Store {
	return &Store{db: db.Database(types.DBName)}
}

func (s Store) CreateTask(ctx context.Context, task *types.TaskCreateReq) (primitive.ObjectID, error) {
	res, err := s.db.Collection(collName).InsertOne(context.Background(), task)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (s Store) GetTaskByID(ctx context.Context, id string) (*types.TaskModel, error) {
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

func (s Store) GetAllTaskByRoomID(ctx context.Context, roomID string) ([]types.ProjectedTask, error) {
	var tasks []types.ProjectedTask
	id, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, err
	}

	opts := options.Find().SetProjection(bson.M{"_id": 1, "title": 1, "handler": 1, "created_at": 1})
	couser, err := s.db.Collection(collName).Find(ctx, bson.M{"room_id": id}, opts)
	if err != nil {
		return nil, err
	}
	defer couser.Close(ctx)
	if err := couser.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
