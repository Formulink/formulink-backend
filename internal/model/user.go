package model

import "time"

type User struct {
	ID           int       `json:"id"`
	TelegramId   int       `json:"telegram_id"`
	Username     string    `json:"username"`
	RegisteredAt time.Time `json:"registered_at"`
}
