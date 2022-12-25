package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"hacktiv8-final-project-3/httpserver/models"
	"hacktiv8-final-project-3/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	VerifyToken(token string) (bool, interface{}, error)
	GenerateToken(user *models.UserModel) (string, string, error)
}

type authService struct {
	JWT_SECRET_KEY string
}

func NewAuthService() *authService {
	constants := &utils.Constants
	print("constants.JWT_SECRET_KEY", constants.JWT_SECRET_KEY)
	return &authService{
		JWT_SECRET_KEY: constants.JWT_SECRET_KEY,
	}
}

func (s *authService) VerifyToken(accessToken string) (bool, interface{}, error) {
	jwtToken, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		method, isRsa := t.Method.(*jwt.SigningMethodHMAC)
		if !isRsa {
			return nil, errors.New("invalid algorithm")
		}
		if method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid algorithm")
		}

		return []byte(s.JWT_SECRET_KEY), nil
	})

	if jwtToken == nil {
		return false, nil, errors.New("invalid token")
	}

	if err != nil {
		validation, _ := err.(*jwt.ValidationError)
		if validation.Errors == jwt.ValidationErrorExpired {
			return false, nil, errors.New("token expired")
		}
	}

	_, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok || !jwtToken.Valid {
		return false, nil, errors.New("invalid token")
	}

	return true, jwtToken.Claims.(jwt.MapClaims), nil

}

func (s *authService) GenerateToken(user *models.UserModel) (string, string, error) {
	const ttlAccessToken = 24 * time.Hour
	const ttlRefreshToken = (24 * 7) * time.Hour

	var userMap map[string]interface{}
	data, err := json.Marshal(user)
	if err != nil {
		return "", "", err
	}

	json.Unmarshal(data, &userMap)

	accessClaims, refreshClaims := jwt.MapClaims{
		"data": userMap,
		"exp":  time.Now().UTC().Add(ttlAccessToken).Unix(),
	}, jwt.MapClaims{
		"data": userMap,
		"exp":  time.Now().UTC().Add(ttlRefreshToken).Unix(),
	}
	fmt.Println("SECRET KEY", s.JWT_SECRET_KEY)
	var secretKeyByte = []byte(s.JWT_SECRET_KEY)

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKeyByte)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKeyByte)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
