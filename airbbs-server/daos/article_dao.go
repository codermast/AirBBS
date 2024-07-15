package daos

import (
	"codermast.com/airbbs/models"
	"errors"
)

// CreateArticle 文章发布
func CreateArticle(article *models.Article) error {
	result := DB.Create(article)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// GetArticle 获取所有文章
func GetArticle(status int) (*[]models.Article, error) {
	var articles []models.Article

	if status == -1 {
		result := DB.Order("id desc").Find(&articles)
		if result.Error != nil {
			return nil, result.Error
		}

	} else {
		result := DB.Order("id desc").Where("status = ?", status).Find(&articles)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return &articles, nil
}

// DeleteArticleByID 根据 ID 删除指定文章
func DeleteArticleByID(articleId string) error {
	result := DB.Table("articles").Where("id = ?", articleId).Delete(nil)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// GetArticleByID 查询指定 ID 文章
func GetArticleByID(articleID string) (*models.Article, error) {
	var article models.Article
	result := DB.Table("articles").Where("id = ?", articleID).First(&article)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &article, nil
}

// UpdateArticleByID 根据 ID 更新指定文章
func UpdateArticleByID(article *models.Article) (*models.Article, error) {
	// 1. 首先根据 ID 查文章是否存在
	articleByID, err := GetArticleByID(article.ID)

	// 2. 查询时异常，即查询失败，即文章不存在
	if err != nil {
		return nil, errors.New("文章不存在")
	}

	// 维护好作者信息
	article.Author = articleByID.Author

	// 此时文章存在，才进行更新
	result := DB.Save(article)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	return article, nil
}

// UpdateArticleListStatusById 根据 ID 批量修改文章状态
func UpdateArticleListStatusById(ids []string, status int) error {
	for _, id := range ids {
		result := DB.Table("articles").Where("id = ?", id).Update("status", status)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
