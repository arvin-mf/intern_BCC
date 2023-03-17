package repository

import (
	"intern_BCC/model"
	"math"

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
	var totalElements int64
	totalElements = r.db.Model(model.Space{}).Find(&spaces).RowsAffected
	result := r.db.
		Model(model.Space{}).
		Limit(pagin.Limit).
		Offset(pagin.Offset).
		Find(&spaces)
	err := result.Error
	if err != nil {
		return nil, 0, err
	}
	return spaces, int(totalElements), err
}

func (r *SpaceRepository) GetSpaceByParam(pagin *model.PaginParam, cat *model.CategoryRequest) ([]model.Space, int, error) {
	var spaces []model.Space
	var err error
	var totalElem int64
	if cat.Kategori == "" && cat.Search == "" {
		totalElem = r.db.Model(model.Space{}).Find(&spaces).RowsAffected
		result := r.db.
			Model(model.Space{}).
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces)
		err = result.Error
		if err != nil {
			return nil, 0, err
		}
	} else if cat.Search == "" {
		totalElem = r.db.Model(model.Space{}).Where("kategori = ?", cat.Kategori).Find(&spaces).RowsAffected
		result := r.db.
			Model(model.Space{}).
			Where("kategori = ?", cat.Kategori).
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces)
		err = result.Error
		if err != nil {
			return nil, 0, err
		}
	} else if cat.Kategori == "" {
		totalElem = r.db.Model(model.Space{}).
			Where("nama LIKE ?", "%"+cat.Search+"%").
			Find(&spaces).RowsAffected
		result := r.db.
			Model(model.Space{}).
			Where("nama LIKE ?", "%"+cat.Search+"%").
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces)
		err = result.Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		totalElem = r.db.Model(model.Space{}).
			Where("kategori = ? AND nama LIKE ?", cat.Kategori, "%"+cat.Search+"%").
			Find(&spaces).RowsAffected
		result := r.db.
			Model(model.Space{}).
			Where("kategori = ? AND nama LIKE ?", cat.Kategori, "%"+cat.Search+"%").
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces)
		err = result.Error
		if err != nil {
			return nil, 0, err
		}
	}

	return spaces, int(totalElem), err
}

func (r *SpaceRepository) GetSpaceByID(id uint) (model.Space, []model.Option, error) {
	space := model.Space{}
	err := r.db.Model(model.Space{}).Preload("Facilities").First(&space, id).Error

	var options []model.Option
	err = r.db.Model(model.Option{}).Where("space_id = ?", id).
		Preload("Dates", func(db *gorm.DB) *gorm.DB {
			return db.Limit(7)
		}).Find(&options).Error
	return space, options, err
}

func (r *SpaceRepository) CreateReview(review *model.Review) error {
	return r.db.Create(review).Error
}

func (r *SpaceRepository) GetReviewsBySpaceID(spaceID uint) ([]model.Review, int, error) {
	var reviews []model.Review
	err := r.db.Where("space_id = ?", spaceID).Find(&reviews).Error

	var count int64
	err = r.db.Model(model.Review{}).Where("space_id = ?", spaceID).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	var space model.Space
	err = r.db.First(&space, spaceID).Error
	if err != nil {
		return nil, 0, err
	}
	space.ReviewsCount = int(count)
	err = r.db.Save(&space).Error

	return reviews, int(count), err
}

func (r *SpaceRepository) UpdateRating(spaceID uint, newRating float64) error {
	space := model.Space{}
	_ = r.db.First(&space, spaceID).Error
	space.Rating = newRating
	err := r.db.Save(&space).Error
	return err
}

func (r *SpaceRepository) GetCustomerByID(id uint) (model.Customer, error) {
	customer := model.Customer{}
	err := r.db.First(&customer, id).Error
	return customer, err
}

func (r *SpaceRepository) GetPicturesByOwnerID(id uint) ([]string, error) {
	var pictures []model.Picture
	err := r.db.Model(model.Picture{}).Where("owner_id = ?", id).Find(&pictures).Error
	if err != nil {
		return nil, err
	}
	var links []string
	for _, picture := range pictures {
		links = append(links, picture.Link)
	}
	return links, err
}

func (r *SpaceRepository) DeleteSpaceByID(id uint) error {
	var space model.Space
	err := r.db.Delete(&space, id).Error
	return err
}

func (r *SpaceRepository) BindBody(c *gin.Context, body interface{}) interface{} {
	return c.ShouldBindWith(body, binding.JSON)
}

func (r *SpaceRepository) BindParam(c *gin.Context, param interface{}) error {
	if err := c.ShouldBindUri(param); err != nil {
		return err
	}
	return c.ShouldBindWith(param, binding.Query)
}

func (r *SpaceRepository) UpdateDistance(spaces []model.Space, loc model.UserLocation) error {
	var err error
	for _, space := range spaces {
		space.Jarak = 111.111 * math.Sqrt((space.Lat-loc.Lat)*(space.Lat-loc.Lat)+(space.Lon-loc.Lon)*(space.Lon-loc.Lon))
		err = r.db.Save(&space).Error
	}
	return err
}

func (r *SpaceRepository) InputLocation(input model.InputLocation) error {
	var space model.Space
	err := r.db.First(&space, input.ID).Error
	if err != nil {
		return err
	}
	space.Lat = input.Lat
	space.Lon = input.Lon
	err = r.db.Save(&space).Error
	return err
}
