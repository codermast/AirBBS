package controllers

import (
	"codermast.com/airbbs/models"
	"codermast.com/airbbs/services"
	"codermast.com/airbbs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// GetAllUsers 获取所有用户 Get /users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()

	// 处理获取所有用户的逻辑
	c.JSON(http.StatusOK, utils.Success("查询成功", users))
}

// GetUserByID 根据 ID 获取指定用户 Get /users/:uid
func (uc *UserController) GetUserByID(c *gin.Context) {
	// 获取路径参数
	userID := c.Param("uid")

	user := services.GetUserByID(userID)
	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, utils.SuccessData(user))
}

// CreateUser 创建用户 POST /users
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	var userRegisterDto models.UserRegisterDto

	// 用户解析
	if err := c.BindJSON(&userRegisterDto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// 获取验证码
	code := userRegisterDto.Code

	if code == "" {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("验证码为空！")))
		return
	}

	user.Username = userRegisterDto.Username
	user.Password = userRegisterDto.Password
	user.Mail = userRegisterDto.Mail
	user.Tel = userRegisterDto.Tel
	user.Nickname = userRegisterDto.Nickname
	user.Admin = false

	// 账号或密码为空
	if (user.Username == "") || (user.Password == "") {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("username or password is empty")))
		return
	}

	// 校验验证码
	redisCode, err := utils.Get(user.Mail)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("请先获取验证码")))
		return
	}

	if redisCode != code {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("验证码错误！")))
		return
	}

	// 保存用户信息
	err = services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusConflict, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// 此时验证码正确且用户注册成功，则通过校验，删除 Redis 中验证码，减少资源占用
	err = utils.Del(userRegisterDto.Mail)
	if err != nil {
		log.Println("验证码删除失败")
	}

	var userVO models.UserVO
	userVO.ID = user.ID
	userVO.Username = user.Username
	userVO.Nickname = user.Nickname
	userVO.Mail = user.Mail
	userVO.Tel = user.Tel

	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, utils.Success("注册成功！", userVO))
}

// UpdateUser 更新指定 userID 的用户信息 PUT /users
func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User

	// 用户解析
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error("用户数据解析失败！"))
		return
	}

	err := services.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}
	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, utils.SuccessMsg("更新成功！"))
}

// DeleteUser 删除指定 userID 的用户 DELETE /users/:uid
func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("uid")

	err := services.DeleteUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, utils.SuccessMsg("删除成功！"))
}

// UserLogin 用户登录 POST /users/login
func (uc *UserController) UserLogin(c *gin.Context) {
	var userLoginDto models.UserLoginDto

	// 用户解析
	if err := c.BindJSON(&userLoginDto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error("用户数据解析失败！"))
		return
	}

	if userLoginDto.Username == "" || userLoginDto.Password == "" {
		c.JSON(http.StatusBadRequest, utils.Error("用户名或账号密码为空！"))
		return
	}

	var user models.User

	user.Username = userLoginDto.Username
	user.Password = userLoginDto.Password

	err := services.UserLogin(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// 登录成功则设置 JWT token
	tokenString, err := utils.GetJwtToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error("Failed to generate token"))
		return
	}

	// 将 Token 返回给客户端
	c.JSON(http.StatusOK, utils.Success("登录成功", tokenString))
}
