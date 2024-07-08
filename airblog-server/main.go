package main

import (
	router "codermast.com/airblog/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化 Gin
	r := gin.Default()

	// 2. 加载路由配置器
	router.AirBlogRouterConfig(r)

	// 3. 启动运行
	err := r.Run(":8080")

	if err != nil {
		return
	}
}
