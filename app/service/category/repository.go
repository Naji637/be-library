package category

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCategory(category *model.Category) error
	GetAllCategories() []ResponseCategory
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCategory(category *model.Category) error {
	err := r.db.Create(category).Error
	return err
}
func (r *repository) GetAllCategories() []ResponseCategory {
	var categories []ResponseCategory
	r.db.Debug().Find(&categories)
	return categories
}