package service

import (
	"junjun-box-api/model"
	"junjun-box-api/repository"
	"time"
)

type MemberService struct {
	memberRepository repository.MemberRepository
}

func (s MemberService) GetMember(openId string) *model.Member {
	m := s.memberRepository.GetByOpenId(openId)
	if m.OpenId == "" {
		m.OpenId = openId
		m.LastLoginTime = time.Now()
		s.memberRepository.Save(m)
	}
	return m
}

func (s MemberService) UpdateMember(m *model.Member) *model.Member {
	s.memberRepository.Update(m)
	return m
}
