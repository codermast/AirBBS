package controllers

import (
	"codermast.com/airblog/models"
	"codermast.com/airblog/services"
	"codermast.com/airblog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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

	// 用户解析
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	// 保存用户信息
	err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}
	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, utils.Success("注册成功！", user))
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

// UserLogin 用户登录
func (uc *UserController) UserLogin(c *gin.Context) {
	var user models.User

	// 用户解析
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error("用户数据解析失败！"))
		return
	}

	err := services.UserLogin(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	tokenString, err := utils.GetJwtToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error("Failed to generate token"))
		return
	}

	// 将 Token 返回给客户端
	c.JSON(http.StatusBadRequest, utils.Success("登录成功", tokenString))
}
