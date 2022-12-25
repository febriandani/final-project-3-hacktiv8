package dto

type UpsertCategoryDto struct {
	Type string `json:"type" binding:"required"`
}
