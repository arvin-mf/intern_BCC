package repository

import (
	"intern_BCC/entity"
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

func (r *OwnerRepository) CreateOwner(model model.CreateOwnerRequest) (*entity.Owner, error) {
	hashPassword, err := crypto.HashValue("123456")
	if err != nil {
		return nil, err
	}
	owner := entity.Owner{
		Email:    model.Email,
		Password: hashPassword,
		Whatsapp: model.Whatsapp,
	}
	result := r.db.Create(&owner)
	if result.Error != nil {
		return nil, err
	}
	return &owner, nil
}

func (r *OwnerRepository) FindByEmail(email string) (entity.Owner, error) {
	owner := entity.Owner{}
	err := r.db.Where("email = ?", email).First(&owner).Error
	return owner, err
}
