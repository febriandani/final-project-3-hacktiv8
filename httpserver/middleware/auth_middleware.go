package middleware

import (
	"hacktiv8-final-project-3/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtGuard(s utils.AuthHelper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Authorization")
		hasPrefix := strings.HasPrefix(accessToken, "Bearer")

		if !hasPrefix {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", "No Bearer Found"))
			return
		}

		splitToken := strings.Split(accessToken, " ")

		if len(splitToken) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", "Invalid Token"))
			return
		}

		jwtToken := splitToken[1]

		isVerified, jwtDecoded, err := s.VerifyToken(string(jwtToken))

		if !isVerified {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", err.Error()))
			return
		}

		userModel := s.JwtClaimsToUserModel(jwtDecoded.(jwt.MapClaims))

		ctx.Set("user", userModel)
		ctx.Next()
	}
}
