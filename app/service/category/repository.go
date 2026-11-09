package category

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCategory(category *model.Category) error
	GetAllCategories() []ResponseCategory
	GetCategorById(id string) ResponseCategory
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
	r.db.Find(&categories)
	return categories
}
func (r *repository) GetCategorById(id string) ResponseCategory {
	var category ResponseCategory
	r.db.Model(model.Category{}).Where("id=?", id).Find(&category)
	return category
}
