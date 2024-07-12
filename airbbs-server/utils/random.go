package utils

import (
	"math/rand"
	"strings"
)

// 定义包含 a-z, A-Z 和 0-9 的所有字符的字符串
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomNickname 随机昵称
func RandomNickname(length int) string {
	preStr := "user_"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return preStr + string(b)
}

// RandomStr 随机字符串
func RandomStr(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

// RandomStrLower 随机字符串
func RandomStrLower(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return strings.ToLower(string(b))
}
