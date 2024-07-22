package daos

import (
	"codermast.com/airbbs/models/po"
	"codermast.com/airbbs/models/vo"
	"errors"
)

// FollowUser 关注指定用户	POST /follow/:id
func FollowUser(userId string, targetId string) error {
	var follow po.Follow

	follow.Follower = userId
	follow.Followed = targetId

	result := DB.Table("follows").Create(&follow)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("关注失败！")
	}
	return nil
}

// UnfollowUser 取消关注指定用户	DELETE /follow/:id
func UnfollowUser(userId string, targetId string) error {
	var follow po.Follow

	follow.Follower = userId
	follow.Followed = targetId

	result := DB.Table("follows").Where("follower = ? AND followed = ?", userId, targetId).Delete(&follow)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("取关失败！")
	}
	return nil
}

// GetUserFans 查看指定用户的粉丝列表	GET /follow/:id
func GetUserFans(userId string) ([]vo.UserVO, error) {

	var fansIDs []string
	result := DB.Table("follows").Select("follower").Where("followed = ?", userId).Find(&fansIDs)

	if result.Error != nil {
		return nil, result.Error
	}

	fansList := make([]vo.UserVO, len(fansIDs))

	for i := 0; i < len(fansIDs); i++ {
		userVo, _ := GetUserByID(fansIDs[i])
		fansList[i] = userVo
	}

	return fansList, nil
}
