package dto

import "github.com/google/uuid"

type NewMessageDto struct {
	UserId         uuid.UUID `json:"user_id"`
	ConversationId uuid.UUID `json:"id"`
	Message        string    `json:"message"`
}
