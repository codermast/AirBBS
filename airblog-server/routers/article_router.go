package router

import (
	"codermast.com/airblog/controllers"
	"codermast.com/airblog/middlewares"
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
		articleRoute.GET("/", articleController.GetArticle)
		articleRoute.DELETE("/:aid", articleController.DeleteArticleByID)
		articleRoute.PUT("/", articleController.UpdateArticleByID)
		articleRoute.GET("/:aid", articleController.GetArticleByID)
	}
}
