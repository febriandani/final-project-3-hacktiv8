package repositories

import (
	"hacktiv8-final-project-3/httpserver/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(*models.CategoryModel) (*models.CategoryModel, error)
	GetCategories(category *models.CategoryModel) (*[]models.CategoryModel, error)
	UpdateCategoryByID(category *models.CategoryModel) (*models.CategoryModel, error)
	DeleteCategoryByID(category *models.CategoryModel) (*models.CategoryModel, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db,
	}
}

func (r *categoryRepository) CreateCategory(category *models.CategoryModel) (*models.CategoryModel, error) {

	err := r.db.Create(category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) GetCategories(category *models.CategoryModel) (*[]models.CategoryModel, error) {
	var categories []models.CategoryModel
	err := r.db.Find(&categories).Order("id desc").Error

	if err != nil {
		return &categories, err
	}
	return &categories, nil
}

func (r *categoryRepository) UpdateCategoryByID(category *models.CategoryModel) (*models.CategoryModel, error) {
	err := r.db.Model(category).Updates(category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) DeleteCategoryByID(category *models.CategoryModel) (*models.CategoryModel, error) {
	err := r.db.Model(category).Delete(category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}
