package router

import (
	"codermast.com/airblog/controllers"
	"codermast.com/airblog/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes 设置用户路由
func SetupUserRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()

	userGroup := r.Group("/users")
	// 登录校验中间件
	userGroup.Use(middlewares.UserLoginAuthMiddleware())
	{
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:uid", userController.GetUserByID)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/", userController.UpdateUser)
		userGroup.DELETE("/:uid", userController.DeleteUser)
		userGroup.POST("/login", userController.UserLogin)
	}
}
