package router

import (
	"codermast.com/airbbs/controllers"
	"codermast.com/airbbs/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupArticleRoutes 设置文章路由
func SetupArticleRoutes(r *gin.Engine) {
	// 文章控制器
	articleController := controllers.NewArticleController()

	// 文章路由器
	articleRoute := r.Group("/articles")

	articleRoute.Use(middlewares.UserLoginAuthMiddleware())

	{
		articleRoute.POST("/", articleController.CreateArticle)
		articleRoute.GET("/all", articleController.GetArticleAllList)
		articleRoute.GET("/publish", articleController.GetArticlePublishList)
		articleRoute.DELETE("/:aid", articleController.DeleteArticleByID)
		articleRoute.PUT("/:aid", articleController.UpdateArticleByID)
		articleRoute.GET("/:aid", articleController.GetArticleByID)
		articleRoute.PUT("/", articleController.UpdateArticleListStatusById)
	}
}
