package model

import "github.com/google/uuid"

type Formula struct {
	Id          uuid.UUID `json:"id"`
	SectionId   uuid.UUID `json:"sectionId"`
	Description string    `json:"description"`
	Expression  string    `json:"expression"`
	Parameters  string    `json:"parameters"`
	Difficulty  int8      `json:"difficulty"`
}

type FormulaLike struct {
	UserID    uuid.UUID `json:"user_id" `
	FormulaID uuid.UUID `json:"formula_id"`
}
