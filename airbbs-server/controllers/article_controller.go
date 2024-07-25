package controllers

import (
	"codermast.com/airbbs/constant"
	"codermast.com/airbbs/models/ro"
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
	var articleCreateRequest ro.ArticleCreateRequest

	// 文章解析
	if err := c.BindJSON(&articleCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// author 作者信息
	articleCreateRequest.Author = c.GetString(constant.USERID)

	// 文章发布
	err := services.CreateArticle(&articleCreateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("发布成功！", articleCreateRequest))
}

// GetArticleAllList 获取所有文章 GET /article/all
func (ac *ArticleController) GetArticleAllList(c *gin.Context) {

	article, err := services.GetArticleByStatus(-1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("查询成功", article))
}

// GetArticlePublishList 获取所有公开发布文章 GET /article/publish
func (ac *ArticleController) GetArticlePublishList(c *gin.Context) {
	article, err := services.GetArticleByStatus(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("查询成功", article))
}

// DeleteArticleByID 根据 ID 删除指定文章 DELETE /article/:aid
func (ac *ArticleController) DeleteArticleByID(c *gin.Context) {
	articleID := c.Param("aid")

	err := services.DeleteArticleByID(articleID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("删除成功"))
}

// GetArticleByID 查询指定 ID 文章 GET /article/:aid
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

	// 浏览量 + 1
	_ = services.UpdateArticleViewsById(article.ID, article.Views+1)

	c.JSON(http.StatusOK, utils.Success("查询成功", article))
}

// UpdateArticleByID 根据 ID 更新指定文章 PUT /article
func (ac *ArticleController) UpdateArticleByID(c *gin.Context) {
	var article ro.ArticleUpdateRequest

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

// UpdateArticleListStatusById 根据 ID 批量修改文章状态
func (ac *ArticleController) UpdateArticleListStatusById(c *gin.Context) {
	var req ro.ArticleUpdateStatusRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	err := services.UpdateArticleListStatusById(req.IDs, req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("更新成功！"))
}

// GetArticleListPage 分页查询文章
func (ac *ArticleController) GetArticleListPage(c *gin.Context) {
	var ArticleListPageRequest ro.ArticleListPageRequest

	// 分页请求解析
	if err := c.ShouldBindQuery(&ArticleListPageRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	ArticleListPage, err := services.GetArticleListPage(&ArticleListPageRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("查询成功！", ArticleListPage))
}

// GetAuthorInfoById 根据作者 ID 获取作者信息
func (ac *ArticleController) GetAuthorInfoById(c *gin.Context) {
	authorId := c.Param("id")

	authorInfo := services.GetAuthorInfoById(authorId)
	c.JSON(http.StatusOK, utils.Success("查询成功！", authorInfo))
}
