package controllers

import (
	"codermast.com/airbbs/daos"
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

// SendRegisterAuthCodeToMail 发送注册验证码
func (ac *AuthController) SendRegisterAuthCodeToMail(c *gin.Context) {
	mail := c.PostForm("mail")

	if mail == "" {
		c.JSON(http.StatusConflict, utils.Error("账户不能为空！"))
		return
	}

	// 1. 获取验证码
	code := utils.GetCode()

	// 2. 发送验证码
	// TODO: 这里先模拟发送
	fmt.Println(fmt.Sprintf("验证码为：%s", code))

	// 3. 存入 Redis
	err := utils.Set(fmt.Sprintf("%s:%s", "register", mail), code, 5*time.Minute)

	if err != nil {
		log.Println("Redis 异常")
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("验证码发送成功！"))
}

// SendResetAuthCodeToMail 发送重置密码验证码
func (ac *AuthController) SendResetAuthCodeToMail(c *gin.Context) {
	account := c.PostForm("account")

	// 账户为空返回
	if account == "" {
		c.JSON(http.StatusConflict, utils.Error("账户不能为空！"))
		return
	}

	// 根据账户获取邮箱账号
	mail, err := daos.GetMailByAccount(account)

	if err != nil {
		c.JSON(http.StatusConflict, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// 发送验证码到邮箱
	// 1. 获取验证码
	code := utils.GetCode()

	// 2. 发送验证码
	// TODO: 这里先模拟发送
	fmt.Println(fmt.Sprintf("验证码为：%s", code))

	// 3. 存入 Redis
	err = utils.Set(fmt.Sprintf("%s:%s", "reset", mail), code, 5*time.Minute)

	if err != nil {
		log.Println("Redis 异常")
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("验证码发送成功！"))
}

// MatchResetAuthCodeToMail 匹配发送重置密码验证码
func (ac *AuthController) MatchResetAuthCodeToMail(c *gin.Context) {
	account := c.PostForm("account")
	code := c.PostForm("code")

	// 账户为空返回
	if account == "" {
		c.JSON(http.StatusConflict, utils.Error("账户不能为空！"))
		return
	}

	// 根据账户获取邮箱账号
	mail, err := daos.GetMailByAccount(account)

	if err != nil {
		c.JSON(http.StatusConflict, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// 获取验证码
	redisCode, err := utils.Get(fmt.Sprintf("%s:%s", "reset", mail))

	if err != nil {
		c.JSON(http.StatusConflict, utils.Error("请先发送验证码！"))
		return
	}

	// 判断验证码是否正确
	if redisCode != code {
		c.JSON(http.StatusConflict, utils.Error("验证码错误！"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("验证码正确！"))
}
