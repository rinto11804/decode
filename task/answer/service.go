package answer

import (
	"context"
	"decode/types"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	store types.AnswerStore
}

type AnswerCreateBody struct {
	Body   string `json:"body"`
	TaskID string `json:"task_id"`
}

func NewService(store types.AnswerStore) *Service {
	return &Service{store}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/answer", s.handleCreateAnswer)
}

func (s *Service) handleCreateAnswer(c echo.Context) error {
	user := c.Get("user").(types.User)
	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	var answerInput AnswerCreateBody
	if err := c.Bind(&answerInput); err != nil {
		return err
	}

	taskID, err := primitive.ObjectIDFromHex(answerInput.TaskID)
	if err != nil {
		return err
	}

	answer := &types.AnswerCreateReq{
		Body:      answerInput.Body,
		TaskID:    taskID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	answerID, err := s.store.CreateAnswer(context.Background(), answer)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"msg":       "answer created successfully",
		"answer_id": answerID.Hex(),
	})
}
