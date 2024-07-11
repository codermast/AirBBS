package main

import (
	"codermast.com/airbbs/config"
	"codermast.com/airbbs/daos"
	router "codermast.com/airbbs/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// 0. 加载项目配置信息
	config.AirBbsSystemConfig()

	// 1. 初始化数据库
	daos.InitDatabase()

	// 2. 初始化 Gin
	r := gin.Default()

	//// 使用默认配置的 CORS 中间件，允许所有来源
	r.Use(cors.Default())

	//3. 加载路由配置器
	router.AirBlogRouterConfig(r)

	// 4. 启动运行
	err := r.Run(":" + config.GetServerConfig().Port)

	if err != nil {
		return
	}
}
