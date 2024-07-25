package router

import (
	"codermast.com/airbbs/controllers"
	"codermast.com/airbbs/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupBlinkRoutes 设置文章路由
func SetupBlinkRoutes(r *gin.Engine) {
	// 动态控制器
	blinkController := controllers.NewBlinkController()

	// 动态路由器
	blinkRoute := r.Group("/blink")

	blinkRoute.Use(middlewares.UserLoginAuthMiddleware())

	blinkRoute.Use(middlewares.UserAdminAuthMiddleware())

	{
		// 动态发布
		blinkRoute.POST("/", blinkController.CreateBlink)

		// 查看动态列表
		blinkRoute.GET("/list", blinkController.GetBlinkList)
	}
}
