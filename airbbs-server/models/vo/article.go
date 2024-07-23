package vo

import "codermast.com/airbbs/models/po"

type ArticleListPageVo struct {
	PageNumber int          `json:"pageNumber"`
	PageSize   int          `json:"pageSize"`
	PageCount  int          `json:"pageCount"`
	TotalCount int          `json:"totalCount"`
	Articles   []po.Article `json:"articles"`
}
