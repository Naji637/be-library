package member_classification_hub

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateMemberClassificationHub(mch *model.MemberClassificationHub) error
	UpdateMemberClassificationHub(memberClassificationId string, memberId string) error
	DeleteMemberClassificationHub(memberId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateMemberClassificationHub(mch *model.MemberClassificationHub) error {
	return r.db.Create(mch).Error
}

func (r *repository) UpdateMemberClassificationHub(memberClassificationId string, memberId string) error {
	return r.db.Model(&model.MemberClassificationHub{}).Where("member_id=?", memberId).Updates(map[string]interface{}{"member_classification_id": memberClassificationId}).Error
}

func (r *repository) DeleteMemberClassificationHub(memberId string) error {
	return r.db.Where("member_id=?", memberId).Delete(&model.MemberClassificationHub{}).Error
}
