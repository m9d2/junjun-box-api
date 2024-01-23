package repository

import (
	"junjun-box-api/model"
)

type MemberRepository struct {
}

func (r MemberRepository) GetByOpenId(openId string) *model.Member {
	var member model.Member
	DB.First(&member, "open_id", openId)
	return &member
}

func (r MemberRepository) Save(member *model.Member) *model.Member {
	DB.Save(member)
	return member
}

func (r MemberRepository) Update(member *model.Member) *model.Member {
	if member.NickName != "" {
		DB.Model(&member).Where("open_id", member.OpenId).Update("nick_name", member.NickName)
	}
	if member.AvatarUrl != "" {
		DB.Model(&member).Where("open_id", member.OpenId).Update("avatar_url", member.AvatarUrl)
	}
	DB.Where("open_id", member.OpenId).First(&member)
	return member
}
