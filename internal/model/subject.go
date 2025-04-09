package model

import "github.com/google/uuid"

type Subject struct {
	Id   uuid.UUID `json:"id"`
	Name uuid.UUID `json:"name"`
}
