package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword 加密算法
func EncryptPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ComparePassword 验证密码
func ComparePassword(password1 string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	return err == nil
}
