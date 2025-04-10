package model

type Section struct {
	Id          int    `json:"id"`
	SubjectId   int    `json:"subject_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
