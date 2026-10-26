package fine

import "be-library/app/model"

type Service interface {
	CreateFine(fine RequestFine) error
	GetAllFines() []ResponseFine
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateFine(fine RequestFine) error {
	err := s.repository.CreateFine(&model.Fine{BorrowingID: fine.BorrowingId, TotalDays: fine.TotalDays, TotalFine: fine.TotalFine, Status: fine.Status})
	return err
}
func (s *service_) GetAllfines() []ResponseFine {
	fines := s.repository.GetAllFine()
	return fines
}
