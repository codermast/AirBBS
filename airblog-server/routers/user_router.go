package router

import (
	"codermast.com/airblog/controllers"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes 设置用户路由
func SetupUserRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:uid", userController.GetUserByID)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/", userController.UpdateUser)
		userGroup.DELETE("/:uid", userController.DeleteUser)
	}
}
