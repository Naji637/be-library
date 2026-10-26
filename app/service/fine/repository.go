package fine

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateFine(fine *model.Fine) error
	GetAllFine() []ResponseFine
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateFine(fine *model.Fine) error {
	err := r.db.Create(fine).Error
	return err
}

func (r *repository) GetAllFine() []ResponseFine {
	var fines []ResponseFine
	r.db.Find(&fines)
	return fines
}
