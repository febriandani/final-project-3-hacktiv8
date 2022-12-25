package dto

type TaskModel struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      *bool     `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  int64     `json:"category_id" gorm:"foreignKey:CategoryID;references:ID"`
	User        UserModel `gorm:"Foreignkey:user_id;association_foreignkey:id;" json:"user,omitempty"`
}

type UserModel struct {
	Full_name string `json:"full_name"`
	Email     string `gorm:"uniqueIndex" json:"email"`
}
