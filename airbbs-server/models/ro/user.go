package ro

type UserRegisterRequest struct {
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

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResetPasswordRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
