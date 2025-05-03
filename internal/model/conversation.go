package model

import (
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Messages  []string  `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
