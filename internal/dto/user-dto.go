package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	telegramId int
	username   string
}

type CreateUserResponse struct {
	id uuid.UUID
}
