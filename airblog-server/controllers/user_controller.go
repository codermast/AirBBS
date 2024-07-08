package controllers

import (
	"codermast.com/airblog/services"
	"codermast.com/airblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// GetAllUsers Get /users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()

	// 处理获取所有用户的逻辑
	c.JSON(http.StatusOK, utils.Success("查询成功", users))
}

// GetUserByID Get /users/:id
func (uc *UserController) GetUserByID(c *gin.Context) {
	// 获取路径参数
	userID := c.Param("id")

	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID",
		"userID":  userID,
	})
}

// CreateUser POST /users
func (uc *UserController) CreateUser(c *gin.Context) {
	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Save User Info Success",
	})
}

// UpdateUser PUT /users/:id
func (uc *UserController) UpdateUser(c *gin.Context) {
	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Update User Info Success",
	})
}

// DeleteUser DELETE /users/:id
func (uc *UserController) DeleteUser(c *gin.Context) {
	// 处理根据用户ID获取用户信息的逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User Success",
	})
}
