package models

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"size:255" json:"username"`
	Password string `gorm:"size:255" json:"password"`
	Nickname string `gorm:"size:255" json:"nickname"`
	Mail     string `gorm:"size:255" json:"mail"`
	Tel      string `gorm:"size:255" json:"tel"`
	Admin    bool   `gorm:"default:false" json:"admin"`
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Mail     string `json:"mail"`
	Tel      string `json:"tel"`
}

type UserVO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Mail     string `json:"mail"`
	Tel      string `json:"tel"`
}

type UserRegisterDto struct {
	ID              string `gorm:"primaryKey;autoIncrement" json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Nickname        string `json:"nickname"`
	Mail            string `json:"mail"`
	Tel             string `json:"tel"`
	Admin           bool   `json:"admin"`
	Code            string `json:"code"`
}

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
