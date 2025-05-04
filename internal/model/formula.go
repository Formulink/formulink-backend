package model

import (
	"github.com/google/uuid"
	"time"
)

type Formula struct {
	Id          uuid.UUID `json:"id"`
	SectionId   int       `json:"sectionId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Expression  string    `json:"expression"`
	Parameters  []string  `json:"parameters"`
	Difficulty  int8      `json:"difficulty"`
}

type FormulaLike struct {
	UserID    uuid.UUID `json:"user_id" `
	FormulaID uuid.UUID `json:"formula_id"`
	CreatedAt time.Time `json:"created_at"`
}
