package repository

import (
	"intern_BCC/model"
	"intern_BCC/sdk/crypto"

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

func (r *OwnerRepository) GetAllOwner() ([]model.Owner, error) {
	var owners []model.Owner
	err := r.db.Find(&owners).Error
	return owners, err
}

func (r *OwnerRepository) GetOwnerByID(id uint) (model.Owner, error) {
	owner := model.Owner{}
	err := r.db.Preload("Spaces").First(&owner, id).Error
	return owner, err
}

func (r *OwnerRepository) DeleteOwnerByID(ID uint) error {
	var owner model.Owner
	err := r.db.Delete(&owner, ID).Error
	return err
}
