package po

type Blink struct {
	Id      string `gorm:"primaryKey;autoIncrement" json:"id"`
	Content string `gorm:"size:65535" json:"content"`
	Author  string `gorm:"size:255" json:"author"`
}
