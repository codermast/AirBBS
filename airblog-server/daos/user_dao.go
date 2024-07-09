package daos

import (
	"codermast.com/airblog/models"
	"errors"
)

// GetAllUsers 查询所有用户
func GetAllUsers() []models.User {
	var users []models.User
	DB.Find(&users)
	return users
}

// GetUserByID 根据 ID 查询指定用户信息
func GetUserByID(userID string) interface{} {
	var user models.User
	// 根据 ID 查询用户
	result := DB.First(&user, userID)

	if result != nil && result.RowsAffected == 0 {
		return nil
	}
	return user
}

// CreateUser 保存用户
func CreateUser(user *models.User) error {
	// 保存用户，Gorm 会自动创建 UserID
	result := DB.Create(user)

	return result.Error
}

// GetUserByUserName 根据 UserName 查用户
func GetUserByUserName(username string) (models.User, error) {
	var userQuery models.User
	result := DB.First(&userQuery, "username = ?", username)

	// 查询为空
	if result != nil && result.RowsAffected == 0 {
		return userQuery, result.Error
	}
	return userQuery, nil
}

// UpdateUser 更新用户
func UpdateUser(user *models.User) error {
	userRet := GetUserByID(user.ID)

	// 用户不存在
	if userRet == nil {
		return errors.New("用户不存在")
	}
	// 更新操作
	result := DB.Save(user)

	// 更新失败
	if !(result.Error == nil && result.RowsAffected == 1) {
		return errors.New("更新失败")
	}
	return result.Error
}

func DeleteUserByID(userID string) error {
	result := DB.Delete(&models.User{}, userID)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return result.Error
}

func UserLogin(user *models.User) error {
	result := DB.Where("username = ? AND password = ?", user.Username, user.Password).First(user)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("登录失败")
	}

	return nil
}
