package repository

import (
	"intern_BCC/model"

	"gorm.io/gorm"
)

type OptionRepository struct {
	db *gorm.DB
}

func NewOptionRepository(db *gorm.DB) OptionRepository {
	return OptionRepository{db}
}

func (r *OptionRepository) CreateOption(option *model.Option) error {
	return r.db.Create(option).Error
}

func (r *OptionRepository) CreateDate(date *model.Date) error {
	return r.db.Create(date).Error
}
