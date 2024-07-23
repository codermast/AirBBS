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
		// 用户注册
		userGroup.POST("/register", userController.CreateUser)

		// 用户查询
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:uid", userController.GetUserByID)
		userGroup.PUT("/", userController.UpdateUser)

		// 用户删除
		userGroup.DELETE("/:uid", userController.DeleteUser)

		// 用户登录
		userGroup.POST("/login", userController.UserLogin)

		// 用户修改
		userGroup.POST("/password/reset", userController.ResetUserPassword)
		userGroup.POST("/photo", userController.UploadUserPhoto)

	}
}
