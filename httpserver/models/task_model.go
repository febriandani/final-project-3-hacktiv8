package models

import (
	"time"
)

type TaskModel struct {
	BaseModel
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      *bool     `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  int64     `json:"category_id" gorm:"foreignKey:CategoryID;references:ID"`
	User        UserModel `gorm:"Foreignkey:user_id;association_foreignkey:id;" json:"user,omitempty"`
}

type User struct {
	BaseModel
	Full_name string `json:"full_name"`
}

func (TaskModel) TableName() string {
	return "public.Task"
}

type TaskParams struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
	UserID      int64
}

type TaskParamsUpdate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      int64
}

type TaskCreateResponse struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type TasksModel struct {
	BaseModel
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      *bool     `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  int64     `json:"category_id" gorm:"foreignKey:CategoryID;references:ID"`
	User        UserModel `gorm:"Foreignkey:user_id;association_foreignkey:id;" json:"user"`
}
