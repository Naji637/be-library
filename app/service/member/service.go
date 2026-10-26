package member

import "be-library/app/model"

type Service interface {
	CreateMember(member RequestMember) error
	GetAllMembers() []ResponseMember

}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateMember(member RequestMember) error {
	err := s.repository.CreateMember(&model.Member{MembershipNumber: member.MembershipNumber, Name: member.Name, Alumni: member.Alumni})
	return err
}

func (s *service_) GetAllMembers() []ResponseMember {
	members := s.repository.GetAllMembers()
	return members
}