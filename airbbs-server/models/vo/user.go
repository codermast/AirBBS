package vo

import "gorm.io/gorm"

type UserVO struct {
	ID        string `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"size:255" json:"username"`
	Nickname  string `gorm:"size:255" json:"nickname"`
	Photo     string `gorm:"size:255" json:"photo"`
	Mail      string `gorm:"size:255" json:"mail"`
	Github    string `gorm:"size:255" json:"github"`
	Tel       string `gorm:"size:255" json:"tel"`
	Admin     bool   `gorm:"default:false" json:"admin"`
	Introduce string `gorm:"size:65535" json:"introduce"`
	Sex       string `gorm:"size:255" json:"sex"`
}

func (userVo *UserVO) AfterFind(tx *gorm.DB) (err error) {
	if userVo.Sex == "1" {
		userVo.Sex = "male"
	} else {
		userVo.Sex = "female"
	}
	return
}

// TableName 返回 UserVO 结构体对应的表名
func (UserVO) TableName() string {
	return "users"
}

type UserLoginVo struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}