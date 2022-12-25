package repositories

import (
	"hacktiv8-final-project-3/httpserver/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Save(data models.TaskModel) (models.TaskModel, error)
	FindAll(userID int64) ([]models.TaskModel, error)
	FindByID(ID int64) (models.TaskModel, error)
	FindByCategoryID(categoryID int64) (models.CategoryModel, error)
	Update(ID int64, data models.TaskModel) (models.TaskModel, error)
	Delete(ID int64) (models.TaskModel, error)
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepo {
	return &taskRepo{db}
}

func (r *taskRepo) Save(data models.TaskModel) (models.TaskModel, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *taskRepo) FindByCategoryID(categoryID int64) (models.CategoryModel, error) {
	var task models.CategoryModel

	err := r.db.Where("id = ?", categoryID).Find(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *taskRepo) FindByID(ID int64) (models.TaskModel, error) {
	var task models.TaskModel

	err := r.db.Where("id = ?", ID).Find(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *taskRepo) FindAll(userID int64) ([]models.TaskModel, error) {
	var task []models.TaskModel

	err := r.db.Preload("User").Where("user_id", userID).Find(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepo) Update(ID int64, data models.TaskModel) (models.TaskModel, error) {

	err := r.db.Where("id = ?", ID).Updates(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *taskRepo) Delete(ID int64) (models.TaskModel, error) {
	var task models.TaskModel

	err := r.db.Where("id = ?", ID).Delete(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}
