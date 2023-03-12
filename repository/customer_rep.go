package repository

import (
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

func (r *CustomerRepository) CreateCustomer(req model.CreateCustomerRequest) (*model.Customer, error) {
	hashPassword, err := crypto.HashValue(req.Password)
	if err != nil {
		return nil, err
	}
	customer := model.Customer{
		Email:    req.Email,
		Password: hashPassword,
		Nama:     req.Nama,
	}
	result := r.db.Create(&customer)
	if result.Error != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) FindByEmail(email string) (model.Customer, error) {
	customer := model.Customer{}
	err := r.db.Where("email = ?", email).First(&customer).Error
	return customer, err
}

func (r *CustomerRepository) GetAllCustomer() ([]model.Customer, error) {
	var customers []model.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *CustomerRepository) GetCustomerByID(id uint) (model.Customer, error) {
	customer := model.Customer{}
	err := r.db.First(&customer, id).Error
	return customer, err
}

func (r *CustomerRepository) DeleteCustomerByID(id uint) error {
	var customer model.Customer
	err := r.db.Delete(&customer, id).Error
	return err
}
