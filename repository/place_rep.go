package repository

import (
	"intern_BCC/entity"
	"intern_BCC/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func (r *PlaceRepository) GetAllPlace(pagin *model.PaginParam) ([]entity.Place, int, error) {
	var places []entity.Place
	err := r.db.
		Model(entity.Place{}).
		Limit(pagin.Limit).
		Offset(pagin.Offset).
		Find(&places).Error
	if err != nil {
		return nil, 0, err
	}
	var totalElements int64
	err = r.db.Model(entity.Place{}).Count(&totalElements).Error
	if err != nil {
		return nil, 0, err
	}
	return places, int(totalElements), err
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

func (h *PlaceRepository) BindBody(c *gin.Context, body interface{}) interface{} {
	return c.ShouldBindWith(body, binding.JSON)
}

func (h *PlaceRepository) BindParam(c *gin.Context, param interface{}) error {
	if err := c.ShouldBindUri(param); err != nil {
		return err
	}
	return c.ShouldBindWith(param, binding.Query)
}
