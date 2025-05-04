package dto

type LikeRequest struct {
	UserId    string `json:"user_id"`
	FormulaId string `json:"formula_id"`
}
