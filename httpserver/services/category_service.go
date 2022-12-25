package services

import (
	"hacktiv8-final-project-3/httpserver/dto"
	"hacktiv8-final-project-3/httpserver/models"
	"hacktiv8-final-project-3/httpserver/repositories"
)

type CategoryService interface {
	CreateCategory(dto *dto.UpsertCategoryDto, userID uint) (*models.CategoryModel, error)
	GetCategories(userID uint) (*[]models.CategoryModel, error)
	UpdateCategory(dto *dto.UpsertCategoryDto, categoryID, userID uint) (*models.CategoryModel, error)
	DeleteCategory(categoryID uint, userID uint) (*models.CategoryModel, error)
	// GetCategory()
}

type categoryService struct {
	repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *categoryService {
	return &categoryService{
		categoryRepository,
	}
}

func (s *categoryService) CreateCategory(dto *dto.UpsertCategoryDto, userID uint) (*models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
		Type:   dto.Type,
	}

	result, err := s.CategoryRepository.CreateCategory(&category)

	if err != nil {
		return &category, err
	}

	return result, nil
}

func (s *categoryService) GetCategories(userID uint) (*[]models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
	}

	result, err := s.CategoryRepository.GetCategories(&category)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *categoryService) UpdateCategory(dto *dto.UpsertCategoryDto, categoryID, userID uint) (*models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
		BaseModel: models.BaseModel{
			ID: categoryID,
		},
		Type: dto.Type,
	}

	result, err := s.CategoryRepository.UpdateCategoryByID(&category)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *categoryService) DeleteCategory(categoryID, userID uint) (*models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
		BaseModel: models.BaseModel{
			ID: categoryID,
		},
	}

	result, err := s.CategoryRepository.DeleteCategoryByID(&category)

	if err != nil {
		return result, err
	}

	return result, nil
}
