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
	JoinMembersToMemberClassifications(membershpNumber string) (*ResponseJoinMembersToMemberClassifications, error)
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

func (r *repository) JoinMembersToMemberClassifications(membershpNumber string) (*ResponseJoinMembersToMemberClassifications, error) {
	var mch ResponseJoinMembersToMemberClassifications

	query := `
			SELECT m.id, m.membership_number, m.name, mc.max_books, mc.max_days
			FROM members AS m
			INNER JOIN member_classification_hubs AS mch ON m.id = mch.member_id
			INNER JOIN member_classifications AS mc ON mc.id = mch.member_classification_id
			WHERE m.membership_number = ?`
	if err := r.db.Raw(query, membershpNumber).Scan(mch).Error; err != nil {
		return nil, err
	}
	return &mch, nil
}
