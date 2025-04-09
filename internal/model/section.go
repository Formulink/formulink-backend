package model

import "github.com/google/uuid"

type Section struct {
	Id          uuid.UUID `json:"id"`
	SubjectId   uuid.UUID `json:"subject_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
