package repositories

import (
	"context"
	"errors"

	apperrors "github.com/agussuartawan/project-test-balabali/internal/errors"
	"github.com/agussuartawan/project-test-balabali/internal/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAll(c context.Context) ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.WithContext(c).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) GetPendingOrders(c context.Context) ([]uint, error) {
	var orderIDs []uint
	if err := r.db.WithContext(c).
		Model(&model.Order{}).
		Where("status = ?", "pending").
		Select("id").
		Find(&orderIDs).Error; err != nil {
		return nil, err
	}

	return orderIDs, nil
}

func (r *OrderRepository) UpdateStatus(c context.Context, ids []uint, status string) error {
	if err := r.db.WithContext(c).
		Model(&model.Order{}).
		Where("id IN ?", ids).
		Updates(map[string]any{"status": status}).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Order not found", nil)
		}
		return err
	}
	
	return nil
}

func (r *OrderRepository) GetById(c context.Context, id uint) (*model.Order, error) {
	var order model.Order
	if err := r.db.WithContext(c).First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Order not found", nil)
		}
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetOrderItems(c context.Context, orderID uint) ([]model.OrderItem, error) {
	var orderItems []model.OrderItem
	if err := r.db.WithContext(c).Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		return nil, err
	}
	
	return orderItems, nil
}

func (r *OrderRepository) Create(c context.Context, tx *gorm.DB, order *model.Order) error {
	return tx.WithContext(c).Create(order).Error
}