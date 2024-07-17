package router

import (
	"codermast.com/airbbs/controllers"
	"codermast.com/airbbs/middlewares"
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
		userGroup.POST("/register", userController.CreateUser)
		userGroup.PUT("/", userController.UpdateUser)
		userGroup.DELETE("/:uid", userController.DeleteUser)
		userGroup.POST("/login", userController.UserLogin)
		userGroup.POST("/password/reset", userController.ResetUserPassword)
		userGroup.POST("/photo", userController.UploadUserPhoto)
	}
}
