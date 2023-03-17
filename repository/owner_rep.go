package repository

import (
	"intern_BCC/model"
	"intern_BCC/sdk/crypto"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OwnerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnerRepository {
	return OwnerRepository{db}
}

func (r *OwnerRepository) CreateOwner(req model.CreateOwnerRequest) (*model.Owner, error) {
	hashPassword, err := crypto.HashValue("12345678")
	if err != nil {
		return nil, err
	}
	owner := model.Owner{
		Email:    req.Email,
		Password: hashPassword,
		Whatsapp: req.Whatsapp,
	}
	result := r.db.Create(&owner)
	if result.Error != nil {
		return nil, err
	}
	return &owner, nil
}

func (r *OwnerRepository) FindByEmail(email string) (model.Owner, error) {
	owner := model.Owner{}
	err := r.db.Where("email = ?", email).First(&owner).Error
	return owner, err
}

func (r *OwnerRepository) GetOwnerSpaces(id uint) ([]model.Space, error) {
	var spaces []model.Space
	err := r.db.Where("owner_id = ?", id).Find(&spaces).Error
	return spaces, err
}

func (r *OwnerRepository) GetOwnerReviews(id []uint) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Where("space_id IN ?", id).Find(&reviews).Error
	return reviews, err
}

func (r *OwnerRepository) GetOwnerSpaceByCat(ownerID uint, category int) (model.Space, error) {
	var space model.Space
	err := r.db.
		Preload("Facilities").Preload("Options").Preload("Options.Dates").
		Where("owner_id = ? AND kategori = ?", ownerID, model.Category[category-1]).
		First(&space).Error
	return space, err
}

func (r *OwnerRepository) GetReviewsBySpaceID(id uint) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Model(model.Review{}).Where("space_id = ?", id).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, err
}

func (r *OwnerRepository) UpdateDescription(id uint, desc string) error {
	var owner model.Owner
	err := r.db.First(&owner, id).Error
	if err != nil {
		return err
	}
	owner.Deskripsi = desc
	err = r.db.Save(&owner).Error
	return err
}

func (r *OwnerRepository) UpdateCapacity(id uint, capac int) error {
	var owner model.Owner
	err := r.db.First(&owner, id).Error
	if err != nil {
		return err
	}
	owner.Kapasitas = capac
	err = r.db.Save(&owner).Error
	return err
}

func (r *OwnerRepository) AddGeneralFacility(id uint, facils []string) error {
	var owner model.Owner
	err := r.db.First(&owner, id).Error
	if err != nil {
		return err
	}

	for _, description := range facils {
		facility := model.GeneralFacility{
			Ket:     description,
			OwnerID: owner.ID,
		}
		result := r.db.Create(&facility)
		if result.Error != nil {
			return err
		}
	}
	return err
}

func (r *OwnerRepository) AddFacilities(id uint, facils []string) error {
	var space model.Space
	err := r.db.First(&space, id).Error
	if err != nil {
		return err
	}

	for _, description := range facils {
		facility := model.Facility{
			Ket:     description,
			SpaceID: space.ID,
		}
		result := r.db.Create(&facility)
		if result.Error != nil {
			return err
		}
	}
	return err
}

func (r *OwnerRepository) UpdatePrice(id uint, harga int) error {
	var space model.Space
	err := r.db.First(&space, id).Error
	if err != nil {
		return err
	}
	space.Harga = harga
	err = r.db.Save(&space).Error
	return err
}

func (r *OwnerRepository) SwitchAvailability(id uint) (bool, error) {
	var date model.Date
	err := r.db.First(&date, id).Error
	if err != nil {
		return false, err
	}
	date.Tersedia = !date.Tersedia
	err = r.db.Save(&date).Error
	return date.Tersedia, err
}

func (r *OwnerRepository) GetSpaceByID(id uint) (model.Space, error) {
	space := model.Space{}
	err := r.db.First(&space, id).Error
	return space, err
}

func (r *OwnerRepository) GetOptionByID(id uint) (model.Option, error) {
	option := model.Option{}
	err := r.db.First(&option, id).Error
	return option, err
}

func (r *OwnerRepository) GetDateByID(id uint) (model.Date, error) {
	date := model.Date{}
	err := r.db.First(&date, id).Error
	return date, err
}

func (r *OwnerRepository) AddPicture(id uint, link string) error {
	var space model.Space
	err := r.db.First(&space, id).Error
	if err != nil {
		return err
	}

	space.Foto = link
	err = r.db.Save(&space).Error
	return err
}

func (r *OwnerRepository) AddGalleryPicture(picture *model.Picture) error {
	return r.db.Create(&picture).Error
}

func (r *OwnerRepository) GetAllPictures(id uint) ([]string, error) {
	var spaces []model.Space
	err := r.db.Model(model.Space{}).Where("owner_id = ?", id).Find(&spaces).Error
	if err != nil {
		return nil, err
	}
	var links []string
	for _, space := range spaces {
		links = append(links, space.Foto)
	}

	var pictures []model.Picture
	err = r.db.Model(model.Picture{}).Where("owner_id = ?", id).Find(&pictures).Error
	if err != nil {
		return nil, err
	}
	for _, pic := range pictures {
		links = append(links, pic.Link)
	}

	return links, err
}

func (r *OwnerRepository) GetAllOwner() ([]model.Owner, error) {
	var owners []model.Owner
	err := r.db.Find(&owners).Error
	return owners, err
}

func (r *OwnerRepository) GetOwnerByID(id uint) (model.Owner, error) {
	owner := model.Owner{}
	err := r.db.First(&owner, id).Error
	return owner, err
}

func (r *OwnerRepository) DeleteOwnerByID(ID uint) error {
	var owner model.Owner
	err := r.db.Delete(&owner, ID).Error
	return err
}

var supClient = supabasestorageuploader.NewSupabaseClient(
	"https://hszlytnrahmcgiyjovbp.supabase.co",
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imhzemx5dG5yYWhtY2dpeWpvdmJwIiwicm9sZSI6ImFub24iLCJpYXQiOjE2Nzg0NDU0NzMsImV4cCI6MTk5NDAyMTQ3M30.xg_QyHoaT-XjUPRVi8NE1p_CpBOS86F1ip0cuBYwMgA",
	"foto",
	"",
)

func (h *OwnerRepository) Upload(c *gin.Context) (string, error) {
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
