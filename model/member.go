package model

import "time"

type Member struct {
	Id            uint      `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement;"`
	OpenId        string    `json:"openid" gorm:"column:open_id;type:varchar(255);uniqueIndex:open_id;"`
	NickName      string    `json:"nickName" gorm:"column:nick_name;type:varchar(255);"`
	AvatarUrl     string    `json:"avatarUrl" gorm:"column:avatar_url;type:varchar(255);"`
	LastLoginIp   string    `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(255);"`
	LastLoginTime time.Time `json:"last_login_time" gorm:"column:last_login_time;type:datetime;"`
}

func (m Member) TableName() string {
	return "member"
}