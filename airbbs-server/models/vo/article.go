package vo

import "codermast.com/airbbs/models/po"

type ArticleListPageVo struct {
	PageNumber int          `json:"pageNumber"`
	PageSize   int          `json:"pageSize"`
	PageCount  int          `json:"pageCount"`
	TotalCount int          `json:"totalCount"`
	Articles   []po.Article `json:"articles"`
}

type AuthorVo struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	ArticleTotal int    `json:"articleTotal"` // 文章总数
	ViewTotal    int    `json:"viewTotal"`    // 总访问量
	FansTotal    int    `json:"fansTotal"`    // 粉丝量
}
