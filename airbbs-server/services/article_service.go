package services

import (
	"codermast.com/airbbs/daos"
	"codermast.com/airbbs/models/po"
	"codermast.com/airbbs/models/ro"
	"codermast.com/airbbs/models/vo"
	"errors"
)

// CreateArticle 文章发布
func CreateArticle(article *po.Article) error {
	if article.Title == "" {
		return errors.New("文章标题不能为空！")
	}

	return daos.CreateArticle(article)
}

// GetArticle 获取所有文章
func GetArticle(status int) (*[]po.Article, error) {
	article, err := daos.GetArticle(status)
	return article, err
}

// DeleteArticleByID 根据 ID 删除指定文章
func DeleteArticleByID(articleID string) error {
	return daos.DeleteArticleByID(articleID)
}

// GetArticleByID 查询指定 ID 文章
func GetArticleByID(articleID string) (*po.Article, error) { // 根据作者id获取作者名称
	article, err := daos.GetArticleByID(articleID)

	if err != nil {
		return nil, err
	}
	return article, nil
}

// UpdateArticleByID 根据 ID 更新指定文章
func UpdateArticleByID(article *po.Article) (*po.Article, error) {
	return daos.UpdateArticleByID(article)
}

// UpdateArticleListStatusById 根据 ID 批量修改文章状态
func UpdateArticleListStatusById(ids []string, status int) error {
	return daos.UpdateArticleListStatusById(ids, status)
}

func GetArticleListPage(articleListPageRequest *ro.ArticleListPageRequest) (vo.ArticleListPageVo, error) {
	return daos.GetArticleListPage(articleListPageRequest)
}

// GetAuthorInfoById 根据作者 id 获取作者信息
func GetAuthorInfoById(authorId string) vo.AuthorVo {
	return daos.GetAuthorInfoById(authorId)
}

func UpdateArticleViewsById(articleId string, views int) error {
	return daos.UpdateArticleViewsById(articleId, views)
}
