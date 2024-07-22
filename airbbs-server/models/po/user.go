package po

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string    `gorm:"size:255" json:"username"`
	Password     string    `gorm:"size:255" json:"password"`
	Nickname     string    `gorm:"size:255" json:"nickname"`
	Photo        string    `gorm:"size:255" json:"photo"`
	Mail         string    `gorm:"size:255" json:"mail"`
	Github       string    `gorm:"size:255" json:"github"`
	Tel          string    `gorm:"size:255" json:"tel"`
	Admin        bool      `gorm:"default:false" json:"admin"`
	Introduce    string    `gorm:"size:65535" json:"introduce"`
	Sex          string    `gorm:"size:255" json:"sex"`
	RegisterTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"register_time"`
	LoginTime    time.Time `json:"login_time"`
}

func (User) TableName() string {
	return "users"
}

// 保存之前
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Sex == "male" {
		user.Sex = "1"
	} else {
		user.Sex = "0"
	}
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if user.Sex == "male" {
		user.Sex = "1"
	} else {
		user.Sex = "0"
	}
	return
}

func (user *User) AfterFind(tx *gorm.DB) (err error) {
	if user.Sex == "1" {
		user.Sex = "male"
	} else {
		user.Sex = "female"
	}
	return
}
