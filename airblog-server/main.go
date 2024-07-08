package main

import (
	"codermast.com/airblog/config"
	"codermast.com/airblog/daos"
	router "codermast.com/airblog/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	// 0. 加载项目配置信息
	config.AirBlogSystemConfig()

	// 1. 初始化数据库
	daos.InitDatabase()

	// 2. 初始化 Gin
	r := gin.Default()

	// 3. 加载路由配置器
	router.AirBlogRouterConfig(r)

	// 4. 启动运行
	err := r.Run(":8080")

	if err != nil {
		return
	}
}
