package po

import "time"

type Follow struct {
	Id         string    `gorm:"primaryKey;autoIncrement" json:"id"`
	Follower   string    `gorm:"default:0" json:"follower"`
	Followed   string    `gorm:"default:0" json:"followed"`
	FollowTime time.Time `gorm:"autoCreateTime" json:"follow_time"`
}

func (f *Follow) TableName() string {
	return "follows"
}
