package router

import (
	"codermast.com/airbbs/controllers"
	"codermast.com/airbbs/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupFollowRoutes 设置关注路由
func SetupFollowRoutes(r *gin.Engine) {
	followController := controllers.NewFollowController()

	followGroup := r.Group("/follow")

	// 登录校验中间件
	followGroup.Use(middlewares.UserLoginAuthMiddleware())

	{
		// 关注指定用户
		followGroup.POST("/:id", followController.FollowUser)

		// 取消关注指定用户
		followGroup.DELETE("/:id", followController.UnfollowUser)

		// 查看指定用户的粉丝列表
		followGroup.GET("/:id", followController.GetUserFans)

	}
}
