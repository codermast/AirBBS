package daos

import "codermast.com/airblog/models"

// GetAllUsers 查询所有用户
func GetAllUsers() []models.User {
	var users []models.User
	DB.Find(&users)
	return users
}
