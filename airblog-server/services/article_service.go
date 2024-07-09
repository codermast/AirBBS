package services

import (
	"codermast.com/airblog/daos"
	"codermast.com/airblog/models"
	"errors"
)

// CreateArticle 文章发布
func CreateArticle(article *models.Article) error {
	if article.Title == "" {
		return errors.New("article title can not be empty")
	}

	return daos.CreateArticle(article)
}

// GetArticle 获取所有文章
func GetArticle() (*[]models.Article, error) {
	article, err := daos.GetArticle()
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
