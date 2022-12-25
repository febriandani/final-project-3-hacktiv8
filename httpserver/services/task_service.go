package services

import (
	"errors"
	"fmt"
	"hacktiv8-final-project-3/httpserver/models"
	"hacktiv8-final-project-3/httpserver/repositories"
)

type TaskService interface {
	Create(params models.TaskParams) (models.TaskModel, error)
	GetAllTask(userID int64) ([]models.TaskModel, error)
	UpdateTask(ID int64, data models.TaskParamsUpdate) (models.TaskModel, error)
	DeleteTask(ID, userID int64) (models.TaskModel, error)
}

type taskService struct {
	// categoryRepo repositories.CategoryRepository
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) *taskService {
	return &taskService{taskRepo: taskRepo}
}

func (ts *taskService) Create(params models.TaskParams) (models.TaskModel, error) {
	var res models.TaskModel

	fmt.Println(params.CategoryID)

	categoryData, err := ts.taskRepo.FindByCategoryID(int64(params.CategoryID))
	if err != nil {
		return res, err
	}

	if categoryData.ID == 0 {
		return res, errors.New("Category not found")
	}

	res = models.TaskModel{
		Title:       params.Title,
		Description: params.Description,
		CategoryID:  int64(params.CategoryID),
		Status:      newBool(false),
		UserID:      uint(params.UserID),
	}

	data, err := ts.taskRepo.Save(res)
	if err != nil {
		return res, err
	}

	return data, nil
}

func (ts *taskService) GetAllTask(userID int64) ([]models.TaskModel, error) {
	data, err := ts.taskRepo.FindAll(userID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ts *taskService) UpdateTask(ID int64, data models.TaskParamsUpdate) (models.TaskModel, error) {
	var res models.TaskModel

	taskData, err := ts.taskRepo.FindByID(ID)
	if err != nil {
		return res, err
	}
	if taskData.ID == 0 {
		return res, errors.New("task not found")
	}

	task := models.TaskModel{
		Title:       data.Title,
		Description: data.Description,
	}

	newData, err := ts.taskRepo.Update(ID, task)
	if err != nil {
		return newData, err
	}

	return newData, nil
}

func (ts *taskService) DeleteTask(ID, userID int64) (models.TaskModel, error) {
	dataTask, err := ts.taskRepo.FindByID(ID)
	if err != nil {
		return dataTask, err
	}

	if dataTask.UserID != uint(userID) {
		return dataTask, errors.New("you dont have access to delete this tasks")
	}

	data, err := ts.taskRepo.Delete(ID)
	if err != nil {
		return data, err
	}

	return data, nil
}
func newBool(b bool) *bool {
	return &b
}
