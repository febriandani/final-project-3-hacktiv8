package controllers

import (
	"errors"
	"fmt"
	"hacktiv8-final-project-3/httpserver/dto"
	"hacktiv8-final-project-3/httpserver/models"
	"hacktiv8-final-project-3/httpserver/services"
	"hacktiv8-final-project-3/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	// GetCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *categoryController {
	return &categoryController{
		categoryService,
	}
}

// GetCategories godoc
//	@Tags		Categories
//	@Summary	get all categories based on user
//	@Success	200	{object}	utils.HttpSuccess[[]models.CategoryModel]
//	@Failure	400	{object}	utils.HttpError
//	@Failure	500	{object}	utils.HttpError
//	@Router		/categories [get]
//	@Security	BearerAuth
func (c *categoryController) GetCategories(ctx *gin.Context) {

	userCredential, isExist := ctx.Get("user")

	if !isExist {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", errors.New("invalid credential")))
		return
	}

	userModel := userCredential.(models.UserModel)

	category, err := c.categoryService.GetCategories(userModel.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Success Get Categories", category))
}

// CreateCategory godoc
//	@Tags		Categories
//	@Summary	create a category
//	@Param		Category	body		dto.UpsertCategoryDto	true	"Create Category DTO"
//	@Success	201			{object}	utils.HttpSuccess[models.CategoryModel]
//	@Failure	400			{object}	utils.HttpError
//	@Failure	500			{object}	utils.HttpError
//	@Router		/categories [post]
//	@Security	BearerAuth
func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var dto dto.UpsertCategoryDto
	err := ctx.BindJSON(&dto)

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

	category, err := c.categoryService.CreateCategory(&dto, userModel.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewHttpSuccess("Category Created", category))
}

// UpdateCategoryByID godoc
//	@Tags		Categories
//	@Summary	Update a category
//	@Param		id			path		int						true	"Category ID"
//	@Param		SocialMedia	body		dto.UpsertCategoryDto	true	"Update Category On User"
//	@Success	200			{object}	utils.HttpSuccess[models.CategoryModel]
//	@Failure	400			{object}	utils.HttpError
//	@Failure	500			{object}	utils.HttpError
//	@Router		/categories/{id} [patch]
//	@Security	BearerAuth
func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	var dto dto.UpsertCategoryDto
	err := ctx.BindJSON(&dto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	idParam := ctx.Param("id")
	categoryID, err := strconv.Atoi(idParam)

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

	category, err := c.categoryService.UpdateCategory(&dto, uint(categoryID), userModel.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Update Category Success", category))
}

// DeleteCategoryByID godoc
//	@Tags		Categories
//	@Summary	Delete a category
//	@Param		id			path		int						true	"Category ID"
//	@Success  200 {object} utils.HttpSuccess[any]
//	@Failure	400			{object}	utils.HttpError
//	@Failure	500			{object}	utils.HttpError
//	@Router		/categories/{id} [delete]
//	@Security	BearerAuth
func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	categoryID, err := strconv.Atoi(idParam)

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

	_, err = c.categoryService.DeleteCategory(uint(categoryID), userModel.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}
	message := fmt.Sprintf("Social Media ID %d has been deleted", categoryID)
	ctx.JSON(http.StatusOK, utils.NewHttpSuccess(message, struct{}{}))
}
