package controllers

import (
	"codermast.com/airbbs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// SendAuthCodeToMail 发送验证码
func (ac *AuthController) SendAuthCodeToMail(c *gin.Context) {
	mail := c.PostForm("mail")

	if mail == "" {
		c.JSON(http.StatusConflict, utils.Error("请输入邮箱地址！"))
		return
	}

	// 1. 获取验证码
	code := utils.GetCode()

	// 2. 发送验证码
	// TODO: 这里先模拟发送
	fmt.Println(fmt.Sprintf("验证码为：%s", code))

	// 3. 存入 Redis
	err := utils.Set(mail, code, 5*time.Minute)

	if err != nil {
		log.Println("Redis 异常")
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("验证码发送成功！"))
}
