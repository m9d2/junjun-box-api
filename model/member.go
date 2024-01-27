package model

import "time"

type Member struct {
	Id            uint      `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement;"`
	OpenId        string    `json:"openid" gorm:"column:open_id;type:varchar(50);uniqueIndex:open_id;"`
	NickName      string    `json:"nickname" gorm:"column:nick_name;type:varchar(255);"`
	AvatarUrl     string    `json:"avatarUrl" gorm:"column:avatar_url;type:varchar(255);"`
	LastLoginIp   string    `json:"lastLoginIp" gorm:"column:last_login_ip;type:varchar(255);"`
	LastLoginTime time.Time `json:"lastLoginTime" gorm:"column:last_login_time;type:datetime;"`
	CreateTime    time.Time `json:"createTime" gorm:"column:create_time;type:datetime;"`
}

func (m Member) TableName() string {
	return "member"
}
