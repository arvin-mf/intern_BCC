package repository

import (
	"intern_BCC/entity"
	"intern_BCC/model"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
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

func (r *SpaceRepository) GetSpaceByParam(pagin *model.PaginParam, cat *model.CategoryRequest) ([]entity.Space, int, error) {
	var spaces []entity.Space
	if cat.Kategori == "" && cat.Search == "" {
		err := r.db.
			Model(entity.Space{}).
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	} else if cat.Search == "" {
		err := r.db.
			Model(entity.Space{}).
			Where("kategori = ?", cat.Kategori).
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	} else if cat.Kategori == "" {
		err := r.db.
			Model(entity.Space{}).
			Where("nama LIKE ?", "%"+cat.Search+"%").
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		err := r.db.
			Model(entity.Space{}).
			Where("kategori = ? AND nama LIKE ?", cat.Kategori, "%"+cat.Search+"%").
			Limit(pagin.Limit).Offset(pagin.Offset).Find(&spaces).Error
		if err != nil {
			return nil, 0, err
		}
	}

	var totalElements int64
	err := r.db.Model(entity.Space{}).Count(&totalElements).Error
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

func (r *SpaceRepository) AddPicture(id uint, link string) error {
	var space entity.Space
	err := r.db.First(&space, id).Error
	if err != nil {
		return err
	}
	space.Foto = link
	err = r.db.Save(&space).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SpaceRepository) DeleteSpaceByID(id uint) error {
	var space entity.Space
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

var supClient = supabasestorageuploader.NewSupabaseClient(
	"https://hszlytnrahmcgiyjovbp.supabase.co",
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imhzemx5dG5yYWhtY2dpeWpvdmJwIiwicm9sZSI6ImFub24iLCJpYXQiOjE2Nzg0NDU0NzMsImV4cCI6MTk5NDAyMTQ3M30.xg_QyHoaT-XjUPRVi8NE1p_CpBOS86F1ip0cuBYwMgA",
	"foto",
	"",
)

func (h *SpaceRepository) Upload(c *gin.Context) (string, error) {
	file, err := c.FormFile("foto")
	if err != nil {
		return "", err
	}
	link, err := supClient.Upload(file)
	if err != nil {
		return "", err
	}
	return link, nil
}
