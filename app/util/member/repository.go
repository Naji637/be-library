package util

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetByYearsMonths(year uint8, month uint8) model.Membership
	CreateMembership(member model.Membership) error
	UpdateMembership(year, month, number uint8) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByYearsMonths(year uint8, month uint8) model.Membership {
	var member model.Membership
	r.db.Where("year=? and month=?", year, month).Find(&member)
	return member
}

func (r *repository) CreateMembership(member model.Membership) error {
	return r.db.Create(&member).Error
}

func (r *repository) UpdateMembership(year, month, number uint8) error {
	return r.db.Model(&model.Membership{}).Where("year=? and month=?", year, month).Updates(map[string]interface{}{"number": number}).Error
}
