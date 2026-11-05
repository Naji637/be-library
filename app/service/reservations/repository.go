package reservations

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateReservation(reservation *model.Reservation) error
	GetAllReservations() []ResponseReservation
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateReservation(reservation *model.Reservation) error {
	err := r.db.Create(reservation).Error
	return err
}

func (r *repository) GetAllReservations() []ResponseReservation {
	var reservations []ResponseReservation
	r.db.Find(&reservations)
	return reservations
}
