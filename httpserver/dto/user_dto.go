package dto

type UpdateTodoDto struct {
	Status bool `json:"status" binding:"required"`
}

type RegisterUserDto struct {
	Full_name string `json:"full_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

type UpsertUserDto struct {
	Full_name string `json:"full_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
