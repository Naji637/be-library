package member

import (
	"be-library/app/model"
	util "be-library/app/util/member"
	"errors"
)

type Service interface {
	CreateMember(member RequestMember) error
	GetAllMembers() []ResponseMember
}

type service_ struct {
	repository Repository
	member     util.Member
}

func NewService(repository Repository, member util.Member) *service_ {
	return &service_{repository, member}
}

func (s *service_) CreateMember(member RequestMember) error {
	membership := s.member.CreateMembership()
	if membership != "" {
		err := s.repository.CreateMember(&model.Member{MembershipNumber: membership, Name: member.Name, Alumni: member.Alumni})
		if err == nil {
			return nil
		}
	}
	return errors.New("error create membership")
}

func (s *service_) GetAllMembers() []ResponseMember {
	members := s.repository.GetAllMembers()
	return members
}
