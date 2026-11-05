package member

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateMember(member *model.Member) (*model.Member, error)
	GetAllMembers() []ResponseMember
	UpdateMember(member *model.Member, id string) error
	DeleteMember(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateMember(member *model.Member) (*model.Member, error) {
	err := r.db.Create(member).Error
	return member, err
}

func (r *repository) GetAllMembers() []ResponseMember {
	var members []ResponseMember
	r.db.Find(&members)
	return members
}

func (r *repository) UpdateMember(member *model.Member, id string) error {
	return r.db.Model(&model.Member{}).Where("id=?", id).Updates(map[string]interface{}{"name": member.Name, "alumni": member.Alumni}).Error
}

func (r *repository) DeleteMember(id string) error {
	return r.db.Where("id=?", id).Delete(&model.Member{}).Error
}
