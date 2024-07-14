package daos

import (
	"codermast.com/airbbs/models"
	"codermast.com/airbbs/utils"
	"errors"
)

// GetAllUsers 查询所有用户
func GetAllUsers() []models.User {
	var users []models.User
	DB.Find(&users)
	return users
}

// GetUserByID 根据 ID 查询指定用户信息
func GetUserByID(userID string) (models.UserVO, error) {
	var user models.User
	// 根据 ID 查询用户
	result := DB.First(&user, userID)

	if result != nil && result.RowsAffected == 0 {
		return models.UserVO{}, errors.New("用户不存在！")
	}

	// 拼接 UserVo
	var userVo models.UserVO

	userVo.ID = user.ID
	userVo.Username = user.Username
	userVo.Nickname = user.Nickname
	userVo.Mail = user.Mail
	userVo.Github = user.Github
	userVo.Tel = user.Tel
	userVo.Admin = user.Admin
	userVo.Introduce = user.Introduce
	userVo.Sex = user.Sex

	return userVo, nil
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

// UpdateUser 更新用户信息
func UpdateUser(userVo *models.UserVO) error {
	_, err := GetUserByID(userVo.ID)

	// 用户不存在
	if err != nil {
		return errors.New("用户不存在")
	}

	// 更新操作
	result := DB.Save(userVo)

	// 更新失败
	if result.Error != nil {
		return errors.New("更新失败！")
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

	var dbUser models.User
	result := DB.Where("username = ?", user.Username).First(&dbUser)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("用户名不存在！")
	}

	if !utils.ComparePassword(dbUser.Password, user.Password) {
		return errors.New("用户名与密码不匹配！")
	}

	return nil
}
