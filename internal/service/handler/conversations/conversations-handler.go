package conversations

import (
	"database/sql"
	"formulink-backend/internal/dto"
	"formulink-backend/pkg/logger"
	"formulink-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ConversationHandler struct {
	db   *sql.DB
	repo *ConversationRepository
}

func NewConversationHandler(db *sql.DB) *ConversationHandler {
	return &ConversationHandler{
		db:   db,
		repo: NewConversationRepository(db),
	}
}

func (ch *ConversationHandler) CreateNewConversation(c echo.Context) error {
	if utils.ValidateRequest(c) {
		return c.JSON(http.StatusBadRequest, "ur request body is null")
	}
	var req dto.CreateConversationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := ch.repo.createNewConversation(req.UserId)
	if err != nil {
		logger.Lg().Logf(0, "can't create new conversation: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, id)
}

func (ch *ConversationHandler) GetConversation(c echo.Context) error {
	conversationId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	conversation, err := ch.repo.getConversation(conversationId)
	if err != nil {
		logger.Lg().Logf(0, "can't get conversation: %v", err)
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, conversation)
}

func (ch *ConversationHandler) GetAllConversations(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	conversations, err := ch.repo.getAllConversations(userId)
	if err != nil {
		logger.Lg().Logf(0, "can't get conversations: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, conversations)
}

func (ch *ConversationHandler) AddMessage(c echo.Context) error {
	var req dto.NewMessageDto
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ch.repo.AddMessage(req); err != nil {
		logger.Lg().Logf(0, "can't add mesage: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "")
}

func (ch *ConversationHandler) DeleteConversation(c echo.Context) error {
	conversationId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err = ch.repo.deleteConversation(conversationId); err != nil {
		logger.Lg().Logf(0, "can't delete conversation: %v", err)
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, "")
}
