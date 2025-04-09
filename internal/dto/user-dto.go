package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	TelegramId int    `json:"telegram_id"`
	Username   string `json:"username"`
}

type CreateUserResponse struct {
	Id       uuid.UUID `json:"id"`
	HaveData bool      `json:"have_data"`
}
