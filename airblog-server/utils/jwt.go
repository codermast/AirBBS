package utils

import (
	"codermast.com/airblog/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	jwtKey = []byte(config.GetJWTConfig().SigningKey) // 需要设置一个安全的密钥来签名和验证 JWT
)

// Claims 结构定义了 JWT 中包含的内容
type Claims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

// GetJwtToken 生成 JWT
func GetJwtToken(userID string) (string, error) {
	// 过期时间
	hours := config.GetJWTConfig().ExpireHours

	// 创建 JWT 的声明（Claims）
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(hours * time.Hour).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			Issuer:    config.GetJWTConfig().Issuer,
			Subject:   config.GetJWTConfig().Subject,
		},
	}

	// 使用声明创建一个 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名 Token 并获取完整的编码后的字符串
	tokenString, err := token.SignedString(jwtKey)

	// 编码异常时
	if err != nil {
		return "", err
	}
	// 返回 JWT Token 字符串
	return tokenString, nil
}

// VerifyJWTToken 校验客户端传递的 JWT
func VerifyJWTToken(jwtToken string) (*Claims, error) {

	// 解析 Token
	token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 校验 Token 是否有效
	if !token.Valid {
		return nil, errors.New("jwt token invalid")
	}

	// 提取声明（Claims）中的信息
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("claim Extract Error")
	}

	// 返回校验通过的信息
	return claims, nil
}
