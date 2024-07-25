package services

import (
	"codermast.com/airbbs/daos"
	"codermast.com/airbbs/models/vo"
	"errors"
)

// FollowUser 关注指定用户	POST /follow/:id
func FollowUser(userId string, targetId string) error {
	// 先判断 要关注的用户 ID 是否合法

	_, err := daos.GetUserByID(targetId)

	if err != nil {
		return errors.New("要关注的用户不存在！")
	}

	return daos.FollowUser(userId, targetId)
}

// UnfollowUser 取消关注指定用户	DELETE /follow/:id
func UnfollowUser(userId string, targetId string) error {
	return daos.UnfollowUser(userId, targetId)
}

// GetUserFans 查看指定用户的粉丝列表	GET /follow/:id
func GetUserFans(userId string) ([]vo.UserVO, error) {
	// 先判断要查询的用户是否合法
	_, err := daos.GetUserByID(userId)

	if err != nil {
		return nil, errors.New("要关注的用户不存在！")
	}

	return daos.GetUserFans(userId)
}

func GetAlreadyFollowed(curUserId string, targetUserId string) bool {
	return daos.GetAlreadyFollowed(curUserId, targetUserId)
}
