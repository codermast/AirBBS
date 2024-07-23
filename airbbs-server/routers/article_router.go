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

	articleRoute.Use(middlewares.UserAdminAuthMiddleware())

	{
		// 文章发布
		articleRoute.POST("/", articleController.CreateArticle)

		// 文章查询
		articleRoute.GET("/:aid", articleController.GetArticleByID)
		articleRoute.GET("/all", articleController.GetArticleAllList)
		articleRoute.GET("/publish", articleController.GetArticlePublishList)
		articleRoute.GET("/page", articleController.GetArticleListPage)

		// 文章删除
		articleRoute.DELETE("/:aid", articleController.DeleteArticleByID)

		// 文章修改
		articleRoute.PUT("/", articleController.UpdateArticleListStatusById)
		articleRoute.PUT("/:aid", articleController.UpdateArticleByID)

		// 作者查询
		articleRoute.GET("/author/:id", articleController.GetAuthorInfoById)
	}
}
