package dto

import (
	"github.com/google/uuid"
	"time"
)

type NewRecordRequest struct {
	UserId uuid.UUID `json:"user_id"`
	Right  int       `json:"right"`
	Fail   int       `json:"fail"`
}

type NewRecordResponse struct {
	Id          uuid.UUID `json:"id"`
	CompletedAt time.Time `json:"completed_at"`
}
