package member

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateMember(member *model.Member) error
	GetAllMembers() []ResponseMember
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateMember(member *model.Member) error {
	err := r.db.Create(member).Error
	return err
}

func (r *repository) GetAllMembers() []ResponseMember {
	var members []ResponseMember
	r.db.Find(&members)
	return members
}