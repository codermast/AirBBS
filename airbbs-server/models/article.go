package models

type Article struct {
	ID      string `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `gorm:"size:65535" json:"title"`
	Content string `gorm:"size:65535" json:"content"`
	Author  string `gorm:"size:255" json:"author"`
}
