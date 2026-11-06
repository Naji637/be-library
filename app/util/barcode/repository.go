package util

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetByYearsMonths(year, month uint8) model.Barcode
	CreateBarcode(barcode model.Barcode) error
	UpdateBarcode(year, month, number uint8) error
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetByYearsMonths(year, month uint8) model.Barcode {
	var barcode model.Barcode
	r.db.Where("year=? and month=?", year, month).Find(&barcode)
	return barcode
}
func (r *repository) CreateBarcode(barcode model.Barcode) error {
	return r.db.Create(&barcode).Error
}
func (r *repository) UpdateBarcode(year, month, number uint8) error {
	return r.db.Model(&model.Barcode{}).Where("year=? and month=?", year, month).Updates(map[string]interface{}{"number": number}).Error
}
