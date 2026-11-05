package member

import (
	"be-library/app/model"
	"be-library/app/service/member_classification_hub"
	util "be-library/app/util/member"
	"errors"
	"fmt"
)

type Service interface {
	CreateMember(member RequestMember) error
	GetAllMembers() []ResponseMember
	UpdateMember(member RequestMember, id string) error
	DeleteMember(id string) error
}

type service_ struct {
	repository    Repository
	member        util.Member
	repositoryMch member_classification_hub.Repository
}

func NewService(repository Repository, member util.Member, repositoryMch member_classification_hub.Repository) *service_ {
	return &service_{repository, member, repositoryMch}
}

func (s *service_) CreateMember(member RequestMember) error {
	membership := s.member.CreateMembership()
	if membership != "" {
		member_, err := s.repository.CreateMember(&model.Member{MembershipNumber: membership, Name: member.Name, Alumni: member.Alumni})
		fmt.Println(member_.ID)
		if err == nil {
			err := s.repositoryMch.CreateMemberClassificationHub(&model.MemberClassificationHub{MemberID: member_.ID, MemberClassificationID: member.MemberClassificationId})
			if err == nil {
				return nil
			}
		}
	}
	return errors.New("error create membership")
}

func (s *service_) GetAllMembers() []ResponseMember {
	members := s.repository.GetAllMembers()
	return members
}

func (s *service_) UpdateMember(member RequestMember, id string) error {
	if err := s.repository.UpdateMember(&model.Member{Name: member.Name, Alumni: member.Alumni}, id); err == nil {
		if err := s.repositoryMch.UpdateMemberClassificationHub(member.MemberClassificationId, id); err == nil {
			return nil
		}
	}
	return errors.New("error update member")
}

func (s *service_) DeleteMember(id string) error {
	if err := s.repositoryMch.DeleteMemberClassificationHub(id); err == nil {
		if err := s.repository.DeleteMember(id); err == nil {
			return nil
		}
	}
	return errors.New("error delete member")
}
