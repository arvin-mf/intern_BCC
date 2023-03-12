package repository

import (
	"intern_BCC/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return OrderRepository{db}
}

func (r *OrderRepository) CreateOrder(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetAllOrder(id uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Model(&orders).Where("customer_id = ?", id).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetOrderByID(id uint) (model.Order, error) {
	var order model.Order
	err := r.db.Find(&order, id).Error
	return order, err
}

func (r *OrderRepository) CreateReview(review *model.Review) error {
	return r.db.Create(review).Error
}
