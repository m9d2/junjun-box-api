package service

import (
	"junjun-box-api/model"
	"junjun-box-api/repository"
)

type MemberService struct {
	memberRepository repository.MemberRepository
}

func (s MemberService) GetMember(openId string) *model.Member {
	return s.memberRepository.GetByOpenId(openId)
}

func (s MemberService) UpdateMember(m *model.Member) *model.Member {
	s.memberRepository.Update(m)
	return m
}

func (s MemberService) Save(m *model.Member) *model.Member {
	s.memberRepository.Save(m)
	return m
}
