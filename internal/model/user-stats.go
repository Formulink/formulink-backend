package model

import (
	"github.com/google/uuid"
	"time"
)

type UserStats struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	Right       int       `json:"right"`
	Fail        int       `json:"fail"`
	CompletedAt time.Time `json:"completed_at"`
}
