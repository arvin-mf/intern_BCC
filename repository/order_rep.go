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

func (r *OrderRepository) GetReviewsBySpaceID(spaceID uint) ([]model.Review, int, error) {
	var reviews []model.Review
	err := r.db.Where("space_id = ?", spaceID).Find(&reviews).Error

	var count int64
	err = r.db.Model(model.Review{}).Where("space_id = ?", spaceID).Count(&count).Error
	return reviews, int(count), err
}

func (r *OrderRepository) GetSpaceByID(id uint) (model.Space, error) {
	space := model.Space{}
	err := r.db.First(&space, id).Error
	return space, err
}

func (r *OrderRepository) UpdateRating(spaceID uint, newRating float64) error {
	space := model.Space{}
	_ = r.db.First(&space, spaceID).Error
	space.Rating = newRating
	err := r.db.Save(&space).Error
	return err
}
