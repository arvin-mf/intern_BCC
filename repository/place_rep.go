package repository

import (
	"intern_BCC/entity"

	"gorm.io/gorm"
)

type PlaceRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return PlaceRepository{db}
}

func (r *PlaceRepository) CreatePlace(place *entity.Place) error {
	return r.db.Create(place).Error
}

func (r *PlaceRepository) GetAllPlace() ([]entity.Place, error) {
	var places []entity.Place
	err := r.db.Find(&places).Error
	return places, err
}

func (r *PlaceRepository) GetPlaceByID(id uint) (entity.Place, error) {
	place := entity.Place{}
	err := r.db.Preload("Spaces").First(&place, id).Error
	return place, err
}

func (r *PlaceRepository) DeletePlaceByID(id uint) error {
	var place entity.Place
	err := r.db.Delete(&place, id).Error
	return err
}
