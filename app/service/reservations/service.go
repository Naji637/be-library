package reservations

import (
	"be-library/app/model"
)

type Service interface {
	CreateReservation(reservation RequestReservation) error
	GetAllReservations() []ResponseReservation
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateReservation(reservation RequestReservation) error {
	err := s.repository.CreateReservation(&model.Reservation{MemberID: reservation.MemberId, BookHubID: reservation.BookHubId, PickupDeadline: reservation.PickupDeadline})
	return err
}

func (s *service_) GetAllReservations() []ResponseReservation {
	reservations := s.repository.GetAllReservations()
	return reservations
}
