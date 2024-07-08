package router

import (
	"codermast.com/airblog/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUserByID)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
