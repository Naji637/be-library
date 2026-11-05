package condition

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCondition(condition *model.Condition) error
	GetAllConditions() []ResponseCondition
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCondition(condition *model.Condition) error {
	err := r.db.Create(condition).Error
	return err
}	

func (r *repository) GetAllConditions() []ResponseCondition {
	var conditions []ResponseCondition
	r.db.Find(&conditions)
	return conditions
}