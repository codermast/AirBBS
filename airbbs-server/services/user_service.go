package services

import (
	"codermast.com/airbbs/daos"
	"codermast.com/airbbs/models"
	"codermast.com/airbbs/utils"
	"errors"
)

// GetAllUsers 查询所有用户
func GetAllUsers() []models.User {
	users := daos.GetAllUsers()
	return users
}

// GetUserByID 根据ID获取指定用户
func GetUserByID(userID string) (models.UserVO, error) {
	return daos.GetUserByID(userID)
}

// CreateUser 创建用户
func CreateUser(user *models.User) error {

	// 加密密码
	hashedPassword := utils.EncryptPassword(user.Password)

	// 随机昵称
	randomNickname := utils.RandomNickname(10)
	user.Nickname = randomNickname

	// 密码加密后存入数据库
	user.Password = string(hashedPassword)

	// 1. 判断用户名是否重复
	_, err := daos.GetUserByUserName(user.Username)
	if err == nil {
		// 此时用户名已经存在
		return errors.New("用户名已经存在！")
	}
	return daos.CreateUser(user)
}

// UpdateUser 更新用户
func UpdateUser(userVo *models.UserVO) error {
	return daos.UpdateUser(userVo)
}

// DeleteUserByID 根据 ID 删除用户
func DeleteUserByID(userID string) error {
	err := daos.DeleteUserByID(userID)
	return err
}

// UserLogin 用户登录
func UserLogin(user *models.User) error {

	return daos.UserLogin(user)
}

func ResetUserPassword(account string, password string) error {
	mail, err := daos.GetMailByAccount(account)

	if err != nil {
		return err
	}

	// 密码加密
	encryptPassword := utils.EncryptPassword(password)

	return daos.ResetUserPassword(mail, encryptPassword)
}
