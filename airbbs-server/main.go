package main

import (
	"codermast.com/airbbs/config"
	"codermast.com/airbbs/daos"
	router "codermast.com/airbbs/routers"
	"codermast.com/airbbs/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// 0. 加载项目配置信息
	config.AirBbsSystemConfig()

	// 1. 初始化数据库
	daos.DatabaseInit()

	// 初始化 Redis
	utils.RedisInit()

	// 2. 初始化 Gin
	r := gin.Default()

	// 打印每个请求的日志
	r.Use(gin.Logger())

	// 使用默认配置的 CORS 中间件，允许所有来源
	r.Use(cors.Default())

	//3. 加载路由配置器
	router.AirBlogRouterConfig(r)

	// 4. 启动运行
	err := r.Run(":" + config.GetServerConfig().Port)

	if err != nil {
		return
	}
}
