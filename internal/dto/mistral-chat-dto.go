package dto

import "formulink-backend/internal/model"

type MistralChatRequest struct {
	Task model.Task `json:"task"`
	Text string     `json:"text"`
}
