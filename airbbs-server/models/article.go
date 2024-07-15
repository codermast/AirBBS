package models

type Article struct {
	ID      string `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `gorm:"size:65535" json:"title"`
	Content string `gorm:"size:65535" json:"content"`
	Status  int    `gorm:"default:1" json:"status"`
	Author  string `gorm:"size:255" json:"author"`
}

type ArticleUpdateStatusRequest struct {
	IDs    []string `json:"ids"`
	Status int      `json:"status"`
}
