package controllers

import (
	"codermast.com/airbbs/models"
	"codermast.com/airbbs/services"
	"codermast.com/airbbs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct{}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// CreateArticle 文章发布 POST /article
func (ac *ArticleController) CreateArticle(c *gin.Context) {
	var article models.Article

	// 文章解析
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// author 作者信息
	article.Author = c.GetString("userID")

	// 文章发布
	err := services.CreateArticle(&article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("发布成功！", article))
}

// GetArticle 获取所有文章 GET /article
func (ac *ArticleController) GetArticle(c *gin.Context) {
	article, err := services.GetArticle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("查询成功", article))
}

// DeleteArticleByID 根据 ID 删除指定文章 DELETE /articles/:aid
func (ac *ArticleController) DeleteArticleByID(c *gin.Context) {
	articleID := c.Param("aid")

	err := services.DeleteArticleByID(articleID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("删除成功"))
}

// GetArticleByID 查询指定 ID 文章 GET /articles/:aid
func (ac *ArticleController) GetArticleByID(c *gin.Context) {
	articleID := c.Param("aid")

	if articleID == "" {
		c.JSON(http.StatusInternalServerError, utils.Error("文章 ID 不能为空"))
		return
	}

	article, err := services.GetArticleByID(articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error("查询文章失败"))
		return
	}

	c.JSON(http.StatusOK, utils.Success("查询成功", article))
}

// UpdateArticleByID 根据 ID 更新指定文章 PUT /articles
func (ac *ArticleController) UpdateArticleByID(c *gin.Context) {
	var article models.Article

	// 文章解析
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	updateArticle, err := services.UpdateArticleByID(&article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("更新成功", updateArticle))
}
