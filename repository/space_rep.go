package repository

import (
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

func (r *SpaceRepository) CreateSpace(space *model.Space) error {
	return r.db.Create(space).Error
}

func (r *SpaceRepository) GetAllSpace(pagin *model.PaginParam) ([]model.Space, int, error) {
	var spaces []model.Space
	err := r.db.
		Model(model.Space{}).
		Limit(pagin.Limit).
		Offset(pagin.Offset).
		Find(&spaces).Error
	if err != nil {
		return nil, 0, err
	}
	var totalElements int64
	err = r.db.Model(model.Space{}).Count(&totalElements).Error
	if err != nil {
		return nil, 0, err
	}
	return spaces, int(totalElements), err
}

func (r *SpaceRepository) GetSpaceByParam(pagin *model.PaginParam, cat *model.CategoryRequest) ([]model.Space, int, error) {
	var spaces []model.Space
	if cat.Kategori == "" && cat.Search == "" {
		err := r.db.
			Model(model.Space{}).
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	} else if cat.Search == "" {
		err := r.db.
			Model(model.Space{}).
			Where("kategori = ?", cat.Kategori).
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	} else if cat.Kategori == "" {
		err := r.db.
			Model(model.Space{}).
			Where("nama LIKE ?", "%"+cat.Search+"%").
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		err := r.db.
			Model(model.Space{}).
			Where("kategori = ? AND nama LIKE ?", cat.Kategori, "%"+cat.Search+"%").
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	}

	var totalElements int64
	err := r.db.Model(model.Space{}).Count(&totalElements).Error
	if err != nil {
		return nil, 0, err
	}
	return spaces, int(totalElements), err
}

func (r *SpaceRepository) GetSpaceByID(id uint) (model.Space, error) {
	space := model.Space{}
	err := r.db.Preload("Facilities").Preload("Options").Preload("Options.Dates").First(&space, id).Error
	return space, err
}

func (r *SpaceRepository) GetCustomerByID(id uint) (model.Customer, error) {
	customer := model.Customer{}
	err := r.db.First(&customer, id).Error
	return customer, err
}

func (r *SpaceRepository) DeleteSpaceByID(id uint) error {
	var space model.Space
	err := r.db.Delete(&space, id).Error
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
