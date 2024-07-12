package utils

// GetCode 获取验证码
func GetCode() string {
	return RandomStrLower(6)
}
