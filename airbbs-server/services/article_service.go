package services

import (
	"codermast.com/airbbs/daos"
	"codermast.com/airbbs/models"
	"errors"
)

// CreateArticle 文章发布
func CreateArticle(article *models.Article) error {
	if article.Title == "" {
		return errors.New("文章标题不能为空！")
	}

	return daos.CreateArticle(article)
}

// GetArticle 获取所有文章
func GetArticle(status int) (*[]models.Article, error) {
	article, err := daos.GetArticle(status)
	return article, err
}

// DeleteArticleByID 根据 ID 删除指定文章
func DeleteArticleByID(articleID string) error {
	return daos.DeleteArticleByID(articleID)
}

// GetArticleByID 查询指定 ID 文章
func GetArticleByID(articleID string) (*models.Article, error) {
	return daos.GetArticleByID(articleID)
}

// UpdateArticleByID 根据 ID 更新指定文章
func UpdateArticleByID(article *models.Article) (*models.Article, error) {
	return daos.UpdateArticleByID(article)
}

// UpdateArticleListStatusById 根据 ID 批量修改文章状态
func UpdateArticleListStatusById(ids []string, status int) error {
	return daos.UpdateArticleListStatusById(ids, status)
}

func GetArticleListPage(articleListPageRequest *models.ArticleListPageRequest) (models.ArticleListPage, error) {
	return daos.GetArticleListPage(articleListPageRequest)
}
