package utils_test

import (
	"hacktiv8-final-project-3/httpserver/models"
	"hacktiv8-final-project-3/utils"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

const mockSecretKey = "Rebahaners"

func TestNewAuthHelper(t *testing.T) {
	helper := utils.NewAuthHelper(mockSecretKey)
	assert.NotEmpty(t, helper.JWT_SECRET_KEY)
	assert.Equal(t, helper.JWT_SECRET_KEY, mockSecretKey)
}

func TestGenerateToken(t *testing.T) {
	helper := utils.NewAuthHelper(mockSecretKey)

	user := models.UserModel{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Full_name: "test",
		Email:     "test@email.com",
		Password:  "123qweasd",
		Role:      "admin",
	}

	accessToken, refreshToken, err := helper.GenerateToken(&user)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)

	isValid, payload, err := helper.VerifyToken(accessToken)

	assert.NoError(t, err)
	assert.NotEmpty(t, payload)
	assert.True(t, isValid)

	userModel := helper.JwtClaimsToUserModel(payload.(jwt.MapClaims))
	assert.NotZero(t, userModel.ID)
	assert.Equal(t, userModel, user)
}
