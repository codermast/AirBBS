package services

import (
	"codermast.com/airblog/daos"
	"codermast.com/airblog/models"
)

// GetAllUsers 查询所有用户
func GetAllUsers() []models.User {
	users := daos.GetAllUsers()
	return users
}
