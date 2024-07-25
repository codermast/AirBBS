package ro

// ArticleCreateRequest 文章发布请求体
type ArticleCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	Author  string `json:"author"`
}

// ArticleUpdateRequest 文章更新请求体
type ArticleUpdateRequest struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	Author  string `json:"author"`
}

type ArticleUpdateStatusRequest struct {
	IDs    []string `json:"ids"`
	Status int      `json:"status"`
}

type ArticleListPageRequest struct {
	PageNumber int `form:"pageNumber"`
	PageSize   int `form:"pageSize"`
}
