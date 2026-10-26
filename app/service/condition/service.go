package condition

import "be-library/app/model"

type Service interface {
	CreateCondition(condition RequestCondition) error
	GetAllConditions() []ResponseCondition
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateCondition(condition RequestCondition) error {
	err := s.repository.CreateCondition(&model.Condition{Condition: condition.Condition})
	return err
}

func (s *service_) GetAllConditions() []ResponseCondition {
	conditions := s.repository.GetAllConditions()
	return conditions
}
