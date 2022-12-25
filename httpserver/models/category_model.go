package models

type CategoryModel struct {
	BaseModel
	Type   string `json:"type"`
	UserID uint   `json:"user_id"`
}

func (CategoryModel) TableName() string {
	return "public.Category"
}
