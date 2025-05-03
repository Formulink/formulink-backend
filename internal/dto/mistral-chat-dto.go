package dto

import (
	"formulink-backend/internal/model"
	"github.com/google/uuid"
)

type MistralChatRequest struct {
	UserId         uuid.UUID  `json:"user_id"`
	ConversationId uuid.UUID  `json:"conversation_id"`
	Task           model.Task `json:"task"`
	Text           string     `json:"text"`
}
