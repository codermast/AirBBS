package po

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID          string    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:65535" json:"title"`
	Content     string    `gorm:"size:65535" json:"content"`
	PublishTime time.Time `gorm:"autoCreateTime" json:"publish_time"`
	UpdateTime  time.Time `gorm:"autoCreateTime" json:"update_time"`
	Status      int       `gorm:"default:1" json:"status"`
	Author      string    `gorm:"size:255" json:"author"`
}

func (a Article) TableName() string {
	return "article"
}

// BeforeUpdate 在更新之前
func (a Article) BeforeUpdate(tx *gorm.DB) error {
	// 设置更新时间为当前时间
	a.UpdateTime = time.Now()
	return nil
}
