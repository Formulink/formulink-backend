package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	TelegramID     int       `json:"telegram_id"`
	Username       string    `json:"username"`
	NeedOnboarding bool      `json:"need_onboarding"`
}
