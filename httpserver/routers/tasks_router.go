package routers

import (
	controllers "hacktiv8-final-project-3/httpserver/controllers"
	"hacktiv8-final-project-3/httpserver/middleware"
	"hacktiv8-final-project-3/utils"

	"github.com/gin-gonic/gin"
)

func TaskRouter(route *gin.RouterGroup, taskController controllers.TaskController, authService utils.AuthHelper) *gin.RouterGroup {
	taskRouter := route.Group("/tasks")
	{
		taskRouter.POST("", middleware.JwtGuard(authService), taskController.CreateTask)
		taskRouter.GET("", middleware.JwtGuard(authService), taskController.GetTasks)
		taskRouter.PUT("", middleware.JwtGuard(authService), taskController.UpdateTask)
		taskRouter.DELETE("", middleware.JwtGuard(authService), taskController.DeleteTasks)
	}
	return taskRouter
}
