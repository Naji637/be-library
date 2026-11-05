package member_classification

import (
	"be-library/app/model"
)

type Service interface {
	CreateMemberClassification(memberClassification RequestMemberClassification) error
	GetAllMemberClassifications() []ResponseMemberClassification
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateMemberClassification(memberClassification RequestMemberClassification) error {
	err := s.repository.CreateMemberClassification(&model.MemberClassification{Classification: memberClassification.Classification, MaxBooks: memberClassification.MaxBooks, MaxDays: memberClassification.MaxDays})
	return err
}

func (s *service_) GetAllMemberClassifications() []ResponseMemberClassification {
	memberClassification := s.repository.GetAllMemberClassifications()
	return memberClassification
}
