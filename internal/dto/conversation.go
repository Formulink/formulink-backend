package dto

import "github.com/google/uuid"

type CreateConversationRequest struct {
	UserId uuid.UUID `json:"user_id"`
}
