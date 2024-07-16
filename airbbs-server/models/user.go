package models

type User struct {
	ID        string `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"size:255" json:"username"`
	Password  string `gorm:"size:255" json:"password"`
	Nickname  string `gorm:"size:255" json:"nickname"`
	Mail      string `gorm:"size:255" json:"mail"`
	Github    string `gorm:"size:255" json:"github"`
	Tel       string `gorm:"size:255" json:"tel"`
	Admin     bool   `gorm:"default:false" json:"admin"`
	Introduce string `gorm:"size:65535" json:"introduce"`
	Sex       bool   `gorm:"default:false" json:"sex"`
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Mail     string `json:"mail"`
	Tel      string `json:"tel"`
}

type UserVO struct {
	ID        string `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"size:255" json:"username"`
	Nickname  string `gorm:"size:255" json:"nickname"`
	Mail      string `gorm:"size:255" json:"mail"`
	Github    string `gorm:"size:255" json:"github"`
	Tel       string `gorm:"size:255" json:"tel"`
	Admin     bool   `gorm:"default:false" json:"admin"`
	Introduce string `gorm:"size:65535" json:"introduce"`
	Sex       bool   `gorm:"default:false" json:"sex"`
}

// TableName 返回 UserVO 结构体对应的表名
func (UserVO) TableName() string {
	return "users"
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

type UserResetPasswordDto struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type UserLoginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
