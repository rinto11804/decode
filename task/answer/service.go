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
	store     types.AnswerStore
	taskStore types.TaskStore
}

type AnswerCreateBody struct {
	Body   string `json:"body"`
	TaskID string `json:"task_id"`
}

func NewService(store types.AnswerStore, taskStore types.TaskStore) *Service {
	return &Service{store, taskStore}
}

func (s *Service) RegisterRoutes(api *echo.Group) {
	api.POST("/answer", s.handleCreateAnswer)
}

// @Summary			Create Answer
// @Description		create the answer for a task
// @Tags			Answer
// @ID				create-answer
// @Accept			json
// @Produce			json
// @Param			answerInput	body	AnswerCreateBody	true	"create answer request input"
// @Success			200		{object}	types.Response[string] 	"answerId"
// @Router			/decode/answer [post]
// @Security		Bearer
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

	task, err := s.taskStore.GetTaskByID(context.Background(), answerInput.TaskID)
	if err != nil {
		return err
	}

	answer := &types.AnswerCreateReq{
		Body:      answerInput.Body,
		TaskID:    task.ID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	answerID, err := s.store.CreateAnswer(context.Background(), answer)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, types.Response[string]{
		Msg:  "answer created",
		Data: answerID.Hex(),
	})
}
