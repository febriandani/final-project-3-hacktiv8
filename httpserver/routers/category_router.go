package routers

import (
	controllers "hacktiv8-final-project-3/httpserver/controllers"
	"hacktiv8-final-project-3/httpserver/middleware"
	"hacktiv8-final-project-3/utils"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(route *gin.RouterGroup, categoryController controllers.CategoryController, authService utils.AuthHelper) *gin.RouterGroup {
	categoryRouter := route.Group("/categories")
	{
		categoryRouter.POST("", middleware.JwtGuard(authService), categoryController.CreateCategory)
		categoryRouter.GET("", middleware.JwtGuard(authService), categoryController.GetCategories)
		categoryRouter.PATCH(":id", middleware.JwtGuard(authService), categoryController.UpdateCategory)
		categoryRouter.DELETE(":id", middleware.JwtGuard(authService), categoryController.DeleteCategory)
	}
	return categoryRouter
}
