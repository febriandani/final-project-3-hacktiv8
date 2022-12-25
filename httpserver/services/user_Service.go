package services

import (
	"hacktiv8-final-project-3/httpserver/dto"
	"hacktiv8-final-project-3/httpserver/models"
	"hacktiv8-final-project-3/httpserver/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(dto *dto.RegisterUserDto) (*models.UserModel, error)
	RegisterAdmin(dto *dto.RegisterUserDto) (*models.UserModel, error)
	Login(dto *dto.LoginDto) (*models.UserModel, error)
	UpdateUser(dto *dto.UpsertUserDto) (*models.UserModel, error)
	DeleteUser(user *models.UserModel) (bool, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) *userService {
	return &userService{r}
}

func (s *userService) RegisterUser(dto *dto.RegisterUserDto) (*models.UserModel, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dto.Password = string(hashedPassword)

	user := models.UserModel{
		Full_name: dto.Full_name,
		Email:     dto.Email,
		Password:  dto.Password,
		Role:      "member",
	}

	_, err = s.userRepository.Register(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (s *userService) RegisterAdmin(dto *dto.RegisterUserDto) (*models.UserModel, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dto.Password = string(hashedPassword)

	user := models.UserModel{
		Full_name: dto.Full_name,
		Email:     dto.Email,
		Password:  dto.Password,
		Role:      "admin",
	}

	_, err = s.userRepository.Register(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (s *userService) Login(dto *dto.LoginDto) (*models.UserModel, error) {
	user := models.UserModel{
		Email:    dto.Email,
		Password: dto.Password,
	}

	result, err := s.userRepository.Login(&user)
	if err != nil {
		return &user, err
	}

	return result, nil
}

func (s *userService) UpdateUser(dto *dto.UpsertUserDto) (*models.UserModel, error) {

	userModel := models.UserModel{
		Full_name: dto.Full_name,
		Email:     dto.Email,
	}

	user, err := s.userRepository.UpdateUser(&userModel)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) DeleteUser(user *models.UserModel) (bool, error) {
	_, err := s.userRepository.DeleteUser(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
