package services

import (
	"codermast.com/airbbs/daos"
	"codermast.com/airbbs/models/po"
	"codermast.com/airbbs/models/ro"
	"codermast.com/airbbs/models/vo"
	"codermast.com/airbbs/utils"
	"errors"
)

// GetAllUsers 获取所有用户 Get /user/all
func GetAllUsers() []vo.UserVO {
	return daos.GetAllUsers()
}

// GetUserByID 根据 ID 获取指定用户 Get /user/:uid
func GetUserByID(userID string) (vo.UserVO, error) {
	return daos.GetUserByID(userID)
}

// CreateUser 创建用户 POST /user/register
func CreateUser(user *po.User) error {

	// 加密密码
	hashedPassword := utils.EncryptPassword(user.Password)

	// 随机昵称
	randomNickname := utils.RandomNickname(10)
	user.Nickname = randomNickname

	// 密码加密后存入数据库
	user.Password = hashedPassword

	// 判断用户名是否重复
	_, err := daos.GetUserByUserName(user.Username)
	if err == nil {
		// 此时用户名已经存在
		return errors.New("用户名已经存在！")
	}
	return daos.CreateUser(user)
}

// UpdateUser 更新指定 userID 的用户信息 PUT /user/:uid
func UpdateUser(userRo *ro.UserUpdateInfoRequest) error {
	return daos.UpdateUser(userRo)
}

// DeleteUserByID 根据 ID 删除用户
func DeleteUserByID(userID string) error {
	err := daos.DeleteUserByID(userID)
	return err
}

// UserLogin 用户登录
func UserLogin(user po.User) error {

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

// UpdateUserPhoto 上传用户头像
func UpdateUserPhoto(url string, userID string) error {
	return daos.UpdateUserPhoto(url, userID)
}
