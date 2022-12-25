package routers

import (
	controllers "hacktiv8-final-project-3/httpserver/controllers"
	"hacktiv8-final-project-3/httpserver/middleware"
	"hacktiv8-final-project-3/utils"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup, userController controllers.UserController, authService utils.AuthHelper) *gin.RouterGroup {
	userRouter := route.Group("/users")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("login", userController.Login)
		userRouter.PUT("update-account", middleware.JwtGuard(authService), userController.UpdateUser)
		userRouter.DELETE("delete-account", middleware.JwtGuard(authService), userController.DeleteUser)
	}
	return userRouter
}
