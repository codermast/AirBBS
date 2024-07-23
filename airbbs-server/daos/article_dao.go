package daos

import (
	"codermast.com/airbbs/models/po"
	"codermast.com/airbbs/models/ro"
	"codermast.com/airbbs/models/vo"
	"errors"
	"math"
)

// CreateArticle 文章发布
func CreateArticle(article *po.Article) error {
	result := DB.Table("articles").Create(article)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// GetArticle 获取所有文章
func GetArticle(status int) (*[]po.Article, error) {
	var articles []po.Article

	if status == -1 {
		result := DB.Table("articles").Order("id desc").Find(&articles)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		result := DB.Table("articles").Order("id desc").Where("status = ?", status).Find(&articles)
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
func GetArticleByID(articleID string) (*po.Article, error) {
	var article po.Article
	result := DB.Table("articles").Where("id = ?", articleID).First(&article)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &article, nil
}

// UpdateArticleByID 根据 ID 更新指定文章
func UpdateArticleByID(article *po.Article) (*po.Article, error) {
	// 1. 首先根据 ID 查文章是否存在
	_, err := GetArticleByID(article.ID)

	// 2. 查询时异常，即查询失败，即文章不存在
	if err != nil {
		return nil, errors.New("文章不存在")
	}

	// 此时文章存在，才进行更新
	result := DB.Table("articles").Updates(article)

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

func GetArticleListPage(articleListPageRequest *ro.ArticleListPageRequest) (vo.ArticleListPageVo, error) {
	// 页号
	pageNumber := articleListPageRequest.PageNumber

	// 页面大小
	pageSize := articleListPageRequest.PageSize

	// 偏移量
	offset := (pageNumber - 1) * pageSize

	var articleListPage vo.ArticleListPageVo

	var totalCount int64

	DB.Table("articles").Where("status = ?", 1).Count(&totalCount)
	// 分页查询
	result := DB.Table("articles").Where("status = ?", 1).Limit(pageSize).Offset(offset).Find(&articleListPage.Articles)
	if result.Error != nil {
		return vo.ArticleListPageVo{}, result.Error
	}

	articleListPage.PageNumber = pageNumber
	articleListPage.PageSize = pageSize
	articleListPage.TotalCount = int(totalCount)
	articleListPage.PageCount = int(math.Ceil(float64(int(totalCount) / pageSize)))

	return articleListPage, nil
}

// 根据 作者 id 获取名称
func getAuthorNameListByIds(articles []po.Article) {
	for i := range articles {
		var author vo.AuthorVo
		DB.Table("users").Where("id = ?", articles[i].Author).First(&author)

		if author.Nickname != "" {
			articles[i].Author = author.Nickname
		} else {
			articles[i].Author = author.Username
		}
	}
}

// GetAuthorInfoById 根据作者 id 获取作者信息
func GetAuthorInfoById(authorId string) vo.AuthorVo {
	var author vo.AuthorVo
	DB.Table("users").Where("id = ?", authorId).First(&author)

	// 文章总数
	var articleTotal int64
	DB.Table("articles").Where("status = ? AND author = ?", 1, authorId).Count(&articleTotal)
	author.ArticleTotal = int(articleTotal)

	// 浏览量总数
	var viewTotal int64
	DB.Table("articles").Where("status = ? AND author = ?", 1, authorId).Select("SUM(views) as total_views").Scan(&viewTotal)
	author.ViewTotal = int(viewTotal)

	return author
}

func UpdateArticleViewsById(articleId string, views int) error {
	result := DB.Table("articles").Where("id = ?", articleId).Update("views", views)

	return result.Error
}
