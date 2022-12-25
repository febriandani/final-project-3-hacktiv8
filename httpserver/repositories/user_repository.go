package repositories

import (
	"hacktiv8-final-project-3/httpserver/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Register(user *models.UserModel) (*models.UserModel, error)
	Login(user *models.UserModel) (*models.UserModel, error)
	UpdateUser(user *models.UserModel) (*models.UserModel, error)
	DeleteUser(user *models.UserModel) (*models.UserModel, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUser(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Find(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Login(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Where("email = ?", user.Email).First(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Model(user).Where("email = ?", user.Email).Updates(user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *userRepository) DeleteUser(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Preload(clause.Associations).Delete(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
