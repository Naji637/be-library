package officer

import "be-library/app/model"

type Service interface {
	CreateOfficer(officer RequestOfficer) error
	GetAllOfficer() []ResponseOfficer
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateOfficer(officer RequestOfficer) error {
	err := s.repository.CreateOfficer(&model.Officer{Name: officer.Name, Username: officer.Username, Password: officer.Password})
	return err
}
func (s *service_) GetAllOfficer() []ResponseOfficer {
	officers := s.repository.GetAllOfficer()
	return officers
}
