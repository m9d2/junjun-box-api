package repository

import (
	"junjun-box-api/model"
)

type MemberRepository struct {
}

func (r MemberRepository) GetByOpenId(openId string) *model.Member {
	var member model.Member
	db := DB.First(&member, "open_id", openId)
	if db.RowsAffected == 0 {
		return nil
	}
	return &member
}

func (r MemberRepository) Save(member *model.Member) *model.Member {
	DB.Save(member)
	return member
}

func (r MemberRepository) Update(member *model.Member) *model.Member {
	updateFields := make(map[string]interface{})

	if member.NickName != "" {
		updateFields["nick_name"] = member.NickName
	}
	if member.AvatarUrl != "" {
		updateFields["avatar_url"] = member.AvatarUrl
	}
	if !member.LastLoginTime.IsZero() {
		updateFields["last_login_time"] = member.LastLoginTime
	}
	if member.LastLoginIp != "" {
		updateFields["last_login_ip"] = member.LastLoginIp
	}
	if member.LastLoginLocation != "" {
		updateFields["last_login_location"] = member.LastLoginLocation
	}

	if len(updateFields) > 0 {
		DB.Model(&member).Where("open_id", member.OpenId).Updates(updateFields)
	}

	DB.Where("open_id", member.OpenId).First(&member)
	return member
}
