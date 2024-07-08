package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	// 处理获取所有用户的逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all users",
	})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// 获取路径参数
	userID := c.Param("id")

	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID",
		"userID":  userID,
	})
}
