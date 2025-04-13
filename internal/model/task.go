package model

import "github.com/google/uuid"

type Task struct {
	Id         uuid.UUID `json:"id"`
	FormulaId  uuid.UUID `json:"formula_id"`
	Difficulty int       `json:"level"`
	TaskText   string    `json:"text"`
	Result     float64   `json:"result"`
}
