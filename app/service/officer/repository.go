package officer

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOfficer(officer *model.Officer)error
	GetAllOfficer()[]ResponseOfficer
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateOfficer(officer *model.Officer)error{
	err:=r.db.Create(officer).Error
	return err
} 

func (r *repository) GetAllOfficer()[]ResponseOfficer{
	var officers []ResponseOfficer
	r.db.Find(&officers)
	return officers
}