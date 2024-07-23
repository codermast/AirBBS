package ro

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	Tel      string `json:"tel"`
	Code     string `json:"code"`
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

type UserUpdateInfoRequest struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Github    string `json:"github"`
	Mail      string `json:"mail"`
	Photo     string `json:"photo"`
	Sex       string `json:"sex"`
	Tel       string `json:"tel"`
	Introduce string `json:"introduce"`
}
