package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	TelegramID   int       `json:"telegram_id"`
	Username     string    `json:"username"`
	RegisteredAt time.Time `json:"registered_at"`
}
