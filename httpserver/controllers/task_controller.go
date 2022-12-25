package controllers

import (
	"errors"
	"hacktiv8-final-project-3/httpserver/models"
	"hacktiv8-final-project-3/httpserver/services"
	"hacktiv8-final-project-3/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	CreateTask(ctx *gin.Context)
	GetTasks(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	DeleteTasks(ctx *gin.Context)
}

type taskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *taskController {
	return &taskController{
		taskService: taskService,
	}
}

func (tc *taskController) CreateTask(ctx *gin.Context) {

	var params models.TaskParams
	err := ctx.BindJSON(&params)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	params.UserID = int64(userModel.ID)

	taskData, err := tc.taskService.Create(params)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	// ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Category Created", taskData))

	taskResponse := models.TaskCreateResponse{
		ID:          int(taskData.BaseModel.ID),
		Title:       taskData.Title,
		Status:      *taskData.Status,
		Description: taskData.Description,
		UserID:      int(taskData.UserID),
		CategoryID:  int(taskData.CategoryID),
		CreatedAt:   taskData.CreatedAt,
	}

	ctx.JSON(
		http.StatusCreated,
		models.Response{
			StatusCode: http.StatusCreated,
			Message:    "Created",
			Data:       taskResponse,
		},
	)

}

func (tc *taskController) GetTasks(ctx *gin.Context) {

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	userID := int64(userModel.ID)

	taskData, err := tc.taskService.GetAllTask(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	// ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Category Created", taskData))

	ctx.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "Successfully",
			Data:       taskData,
		},
	)

}

func (tc *taskController) UpdateTask(ctx *gin.Context) {

	var params models.TaskParamsUpdate
	err := ctx.BindJSON(&params)

	id := ctx.Request.URL.Query().Get("taskId")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	params.UserID = int64(userModel.ID)

	idConv, _ := strconv.Atoi(id)

	taskData, err := tc.taskService.UpdateTask(int64(idConv), params)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	// ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Category Created", taskData))

	ctx.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "Updated",
			Data:       taskData,
		},
	)

}

func (tc *taskController) DeleteTasks(ctx *gin.Context) {

	id := ctx.Request.URL.Query().Get("taskId")
	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	userID := int64(userModel.ID)

	idConv, _ := strconv.Atoi(id)

	_, err := tc.taskService.DeleteTask(int64(idConv), userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	// ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Category Created", taskData))
	ctx.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "task has been successfully deleted",
		},
	)

}
