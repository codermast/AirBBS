package daos

import (
	"codermast.com/airbbs/models/po"
	"codermast.com/airbbs/models/ro"
	"codermast.com/airbbs/models/vo"
	"codermast.com/airbbs/utils"
	"errors"
	"github.com/jinzhu/copier"
	"time"
)

// GetAllUsers 获取所有用户 Get /user/all
func GetAllUsers() []vo.UserVO {
	var userVos []vo.UserVO
	DB.Table("users").Find(&userVos)
	return userVos
}

// GetUserByID 根据 ID 获取指定用户 Get /user/:uid
func GetUserByID(userID string) (vo.UserVO, error) {
	var userVo vo.UserVO
	// 根据 ID 查询用户
	result := DB.Table("users").First(&userVo, userID)

	if result != nil && result.RowsAffected == 0 {
		return vo.UserVO{}, errors.New("用户不存在！")
	}

	return userVo, nil
}

// CreateUser 创建用户 POST /user/register
func CreateUser(user *po.User) error {
	// 保存用户，Gorm 会自动创建 UserID
	result := DB.Table("users").Create(user)

	return result.Error
}

// GetUserByUserName 根据 UserName 查用户
func GetUserByUserName(username string) (po.User, error) {
	var userQuery po.User
	result := DB.First(&userQuery, "username = ?", username)

	// 查询为空
	if result != nil && result.RowsAffected == 0 {
		return userQuery, result.Error
	}
	return userQuery, nil
}

// UpdateUser 更新指定 userID 的用户信息 PUT /user/:uid
func UpdateUser(userRo *ro.UserUpdateInfoRequest) error {
	_, err := GetUserByID(userRo.ID)

	// 用户不存在
	if err != nil {
		return errors.New("用户不存在")
	}

	var user po.User

	_ = copier.Copy(&user, userRo)

	// 更新操作，仅更新非零值的字段
	result := DB.Table("users").Updates(&user)

	// 更新失败
	if result.Error != nil {
		return errors.New("更新失败！")
	}
	return result.Error
}

func DeleteUserByID(userID string) error {
	result := DB.Delete(&po.User{}, userID)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return result.Error
}

func UserLogin(user *po.User) error {
	oldPassword := user.Password

	result := DB.Table("users").Where("username = ?", user.Username).First(&user)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("用户名不存在！")
	}

	if !utils.ComparePassword(user.Password, oldPassword) {
		return errors.New("用户名与密码不匹配！")
	}
	// 此时登录成功
	user.LoginTime = time.Now()
	// 更新最后登录时间
	DB.Table("users").Where("id = ?", user.ID).Update("login_time", user.LoginTime)

	return nil
}

// IsAdminByUserId 判断指定用户是否为管理员
func IsAdminByUserId(userId string) error {
	var user po.User
	result := DB.First(&user, userId)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("用户不存在！")
	}

	if user.Admin {
		return nil
	}
	return errors.New("权限不足！")
}

func GetMailByAccount(account string) (string, error) {
	var user po.User
	result := DB.First(&user, account)

	if result.Error != nil || result.RowsAffected == 0 {
		return "", errors.New("用户不存在！")
	}

	return user.Mail, nil
}

func ResetUserPassword(mail string, password string) error {
	result := DB.Table("users").Where("mail = ?", mail).Update("password", password)

	return result.Error
}

func UpdateUserPhoto(url string, id string) error {
	DB.Table("users").Where("id = ?", id).Update("photo", url)
	return nil
}
