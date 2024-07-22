package ro

type ArticleUpdateStatusRequest struct {
	IDs    []string `json:"ids"`
	Status int      `json:"status"`
}

type ArticleListPageRequest struct {
	PageNumber int `form:"pageNumber"`
	PageSize   int `form:"pageSize"`
}
