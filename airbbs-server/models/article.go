package models

type Article struct {
	ID      string `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `gorm:"size:65535" json:"title"`
	Content string `gorm:"size:65535" json:"content"`
	Status  int    `gorm:"default:1" json:"status"`
	Author  string `gorm:"size:255" json:"author"`
}

func (Article) TableName() string {
	return "article"
}

type ArticleUpdateStatusRequest struct {
	IDs    []string `json:"ids"`
	Status int      `json:"status"`
}

type ArticleListPageRequest struct {
	PageNumber int `form:"pageNumber"`
	PageSize   int `form:"pageSize"`
}

type ArticleListPage struct {
	PageNumber int       `json:"pageNumber"`
	PageSize   int       `json:"pageSize"`
	PageCount  int       `json:"pageCount"`
	TotalCount int       `json:"totalCount"`
	Articles   []Article `json:"articles"`
}

type Author struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
