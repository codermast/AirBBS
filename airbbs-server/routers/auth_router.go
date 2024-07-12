package router

import (
	"codermast.com/airbbs/controllers"
	"github.com/gin-gonic/gin"
)

// SetupAuthCodeRoutes 设置用户路由
func SetupAuthCodeRoutes(r *gin.Engine) {
	authController := controllers.NewAuthController()

	authCodeGroup := r.Group("/auths")

	{
		// 发送验证码
		authCodeGroup.POST("/", authController.SendAuthCodeToMail)
	}
}
