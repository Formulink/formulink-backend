package handler

import (
	"database/sql"
	"fmt"
	"formulink-backend/internal/dto"
	"formulink-backend/internal/model"
	"formulink-backend/internal/service/handler/conversations"
	"formulink-backend/pkg/logger"
	mistral2 "formulink-backend/pkg/mistral"
	"github.com/gage-technologies/mistral-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type MistralHandler struct {
	db     *sql.DB
	redis  *redis.Client
	client *mistral.MistralClient
	repo   *conversations.ConversationRepository
}

func NewMistralHandler(_db *sql.DB, _redis *redis.Client, apiKey string) *MistralHandler {
	return &MistralHandler{
		db:     _db,
		redis:  _redis,
		client: mistral2.CreateMistralClient(apiKey),
		repo:   conversations.NewConversationRepository(_db),
	}
}

var text string = "НА РУССКОМ ЯЗЫКЕ Ты - профессиональный учитель по физике. Тебе могут дать задачу(НЕ ОБЯЗАТЕЛЬНО), ты должен ее решить, тщательно объяснив используемые методы. Однако ты ОБЯЗАТЕЛЬНО должен использовать приведенную ниже формулу для решения этой задачи. Ответ так же будет упомянут ниже для отсутствия дизинформации. Структурируй свой ответ.\n\n"

func (mh *MistralHandler) Chat(c echo.Context) error {
	var req dto.MistralChatRequest

	err := c.Bind(&req)
	if err != nil {
		logger.Lg().Infof("err: %v", err)
		return c.JSON(http.StatusBadRequest, "err")
	}

	formula, err := mh.getSingleFormula(req.Task.FormulaId)
	if err != nil {
		logger.Lg().Infof("err: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp, err := mh.client.Chat("mistral-small", []mistral.ChatMessage{
		{Role: mistral.RoleSystem, Content: fmt.Sprintf("%s | ЗАДАЧА: %s | НЕОБХОДИМАЯ ФОРМУЛА:  %s | ОТВЕТ %f", text, req.Task.TaskText, formula.Expression, req.Task.Result)},
		{Role: mistral.RoleUser, Content: req.Text},
	}, nil)
	if err != nil {
		logger.Lg().Infof("err: %v", err)
		return c.JSON(http.StatusInternalServerError, "err")
	}

	err = mh.repo.AddMessage(dto.NewMessageDto{
		UserId:         req.UserId,
		ConversationId: req.ConversationId,
		Message:        resp.Choices[0].Message.Content,
	})

	return c.JSON(http.StatusOK, resp.Choices[0].Message.Content)
}

func (mh *MistralHandler) getSingleFormula(id uuid.UUID) (model.Formula, error) {
	var formula model.Formula

	query := `SELECT * from formulas where id = $1`
	row := mh.db.QueryRow(query, id)
	if err := row.Scan(
		&formula.Id,
		&formula.SectionId,
		&formula.Name,
		&formula.Description,
		&formula.Expression,
		pq.Array(&formula.Parameters),
		&formula.Difficulty,
	); err != nil {
		logger.Lg().Infof("err: %v", err)
		return model.Formula{}, err
	}
	return formula, nil
}
