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
		// 注册验证码
		authCodeGroup.POST("/register", authController.SendRegisterAuthCodeToMail)

		// 重置密码验证码
		authCodeGroup.POST("/reset", authController.SendResetAuthCodeToMail)

		// 匹配重置密码验证码
		authCodeGroup.POST("/reset/match", authController.MatchResetAuthCodeToMail)
	}
}
