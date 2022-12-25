package main

import (
	"hacktiv8-final-project-3/config"
	"hacktiv8-final-project-3/docs"
	"hacktiv8-final-project-3/httpserver/controllers"
	"hacktiv8-final-project-3/httpserver/repositories"
	"hacktiv8-final-project-3/httpserver/routers"
	"hacktiv8-final-project-3/httpserver/services"
	"hacktiv8-final-project-3/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

/// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       localhost:3030
// @BasePath                   /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Environment Variables not found")
	}
	app := gin.Default()
	appRoute := app.Group("/api")
	db, _ := config.Connect()

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	routers.UserRouter(appRoute, userController, authService)
	routers.CategoryRouter(appRoute, categoryController, authService)
	routers.TaskRouter(appRoute, taskController, authService)

	docs.SwaggerInfo.Title = "Hacktiv8 final-project-2 API"
	docs.SwaggerInfo.Description = "This is just a simple TODO List"
	docs.SwaggerInfo.Host = "https://hacktiv8-final-project-3-production-c895.up.railway.app/"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run()
}
