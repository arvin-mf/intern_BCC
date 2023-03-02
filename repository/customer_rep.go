package repository

import (
	"intern_BCC/entity"
	"intern_BCC/model"
	"intern_BCC/sdk/crypto"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return CustomerRepository{db}
}

func (r *CustomerRepository) CreateCustomer(model model.CreateCustomerRequest) (*entity.Customer, error) {
	hashPassword, err := crypto.HashValue(model.Password)
	if err != nil {
		return nil, err
	}
	customer := entity.Customer{
		Email:    model.Email,
		Password: hashPassword,
		Nama:     model.Nama,
	}
	result := r.db.Create(&customer)
	if result.Error != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) GetAllCustomer() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *CustomerRepository) GetCustomerByID(id uint) (entity.Customer, error) {
	customer := entity.Customer{}
	err := r.db.First(&customer, id).Error
	return customer, err
}

func (r *CustomerRepository) DeleteCustomerByID(ID uint) error {
	var customer entity.Customer
	err := r.db.Delete(&customer, ID).Error
	return err
}
