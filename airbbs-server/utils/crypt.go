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
// newPassword 是明文密码，oldPassword 是加密后的密码
func ComparePassword(newPassword string, oldPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(newPassword), []byte(oldPassword))
	return err == nil
}
