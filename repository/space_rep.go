package repository

import (
	"intern_BCC/entity"
	"intern_BCC/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type SpaceRepository struct {
	db *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) SpaceRepository {
	return SpaceRepository{db}
}

func (r *SpaceRepository) CreateSpace(space *entity.Space) error {
	return r.db.Create(space).Error
}

func (r *SpaceRepository) GetAllSpace(pagin *model.PaginParam) ([]entity.Space, int, error) {
	var spaces []entity.Space
	err := r.db.
		Model(entity.Space{}).
		Limit(pagin.Limit).
		Offset(pagin.Offset).
		Find(&spaces).Error
	if err != nil {
		return nil, 0, err
	}
	var totalElements int64
	err = r.db.Model(entity.Space{}).Count(&totalElements).Error
	if err != nil {
		return nil, 0, err
	}
	return spaces, int(totalElements), err
}

func (r *SpaceRepository) GetSpaceByKategori(pagin *model.PaginParam, cat *model.CategoryRequest) ([]entity.Space, int, error) {
	var spaces []entity.Space
	err := r.db.
		Model(entity.Space{}).
		Where("kategori = ?", cat.Kategori).
		Limit(pagin.Limit).
		Offset(pagin.Offset).
		Find(&spaces).Error
	if err != nil {
		return nil, 0, err
	}
	var totalElements int64
	err = r.db.Model(entity.Space{}).Count(&totalElements).Error
	if err != nil {
		return nil, 0, err
	}
	return spaces, int(totalElements), err
}

func (r *SpaceRepository) GetSpaceByID(id uint) (entity.Space, error) {
	space := entity.Space{}
	err := r.db.Preload("Options").First(&space, id).Error
	return space, err
}

func (r *SpaceRepository) DeleteSpaceByID(ID uint) error {
	var space entity.Space
	err := r.db.Delete(&space, ID).Error
	return err
}

func (h *SpaceRepository) BindBody(c *gin.Context, body interface{}) interface{} {
	return c.ShouldBindWith(body, binding.JSON)
}

func (h *SpaceRepository) BindParam(c *gin.Context, param interface{}) error {
	if err := c.ShouldBindUri(param); err != nil {
		return err
	}
	return c.ShouldBindWith(param, binding.Query)
}
