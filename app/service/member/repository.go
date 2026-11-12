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
	JoinMembersToMemberClassifications(memberId string) (ResponseJoinMembersToMemberClassifications, error)
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

// func (r *repository) JoinMembersToMemberClassifications(memberId string) (ResponseJoinMembersToMemberClassifications, error) {
// 	var memberClassificationHub ResponseJoinMembersToMemberClassifications
// 	err := r.db.Debug().Preload("MemberClassificationHub", func(db *gorm.DB) *gorm.DB {
// 		return db.Joins("Member").Joins("MemberClassification")
// 	}).Where("member_id=?", memberId).Find(&memberClassificationHub).Error
// 	return memberClassificationHub, err
// }

func (r *repository) JoinMembersToMemberClassifications(memberId string) (ResponseJoinMembersToMemberClassifications, error) {
	var memberClassificationHub ResponseJoinMembersToMemberClassifications

	err := r.db.Table("member_classification_hubs").
		Select("member_classification_hubs.*, members.*, member_classifications.*").
		Joins("join members on members.id = member_classification_hubs.member_id").
		Joins("join member_classifications on member_classifications.id = member_classification_hubs.member_classification_id").
		Where("member_classification_hubs.member_id = ?", memberId).
		Scan(&memberClassificationHub).Error

	return memberClassificationHub, err
}
