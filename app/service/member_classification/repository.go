package member_classification

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateMemberClassification(memberClassification *model.MemberClassification) error
	GetAllMemberClassifications() []ResponseMemberClassification
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateMemberClassification(memberClassification *model.MemberClassification) error {
	err := r.db.Create(memberClassification).Error
	return err
}

func (r *repository) GetAllMemberClassifications() []ResponseMemberClassification {
	var memberClassifications []ResponseMemberClassification
	r.db.Find(&memberClassifications)
	return memberClassifications
}
