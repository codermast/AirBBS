package router

import (
	"github.com/gin-gonic/gin"
)

// AirBlogRouterConfig 路由配置
func AirBlogRouterConfig(r *gin.Engine) {
	// 设置用户模块路由
	SetupUserRoutes(r)
}
