package models

type UserModel struct {
	BaseModel
	Full_name     string          `json:"full_name"`
	Email         string          `gorm:"uniqueIndex" json:"email"`
	Password      string          `json:"password"`
	Role          string          `json:"role"`
	TaskModel     []TaskModel     `json:"task" gorm:"foreignKey:UserID;references:ID"`
	CategoryModel []CategoryModel `json:"category" gorm:"foreignKey:UserID;references:ID"`
}

func (UserModel) TableName() string {
	return "public.User"
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
