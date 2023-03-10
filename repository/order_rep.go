package repository

import (
	"intern_BCC/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return OrderRepository{db}
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetAllOrder(id uint) ([]entity.Order, error) {
	var orders []entity.Order

	err := r.db.Model(&orders).Where("customer_id = ?", id).Find(&orders).Error

	return orders, err
}
