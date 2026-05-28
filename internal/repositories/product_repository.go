package repositories

import (
	"context"
	"errors"

	apperrors "github.com/agussuartawan/project-test-balabali/internal/errors"
	"github.com/agussuartawan/project-test-balabali/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(c context.Context, product *model.Product) error {
	return r.db.WithContext(c).Create(product).Error
}

func (r *ProductRepository) GetAll(c context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(c).Find(&products).Error; err != nil {
		return nil, err
	}
	
	return products, nil
}

func (r *ProductRepository) GetById(c context.Context, id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.WithContext(c).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Product not found", nil)
		}
		
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Update(c context.Context, id uint, product *model.Product) error {
	if err := r.db.WithContext(c).Model(&model.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Delete(c context.Context, id uint) error {
	if err := r.db.WithContext(c).Delete(&model.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Exists(c context.Context, id uint) bool {
	var count int64
	if err := r.db.WithContext(c).
		Model(&model.Product{}).
		Select("id").
		Where("id = ?", id).
		Count(&count).
		Error; err != nil {
		return false
	}

	return count > 0
}

func (r *ProductRepository) GetByIds(c context.Context, ids []uint) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.WithContext(c).Where("id IN ?", ids).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) DecrementStock(c context.Context, tx *gorm.DB, id uint, quantity int) error {
	var product model.Product

	if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).
		First(&product, id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Product not found", nil)
		}
		return err
	}

	if product.Stock < quantity {
		return apperrors.NewUnprocessableEntityError("Insufficient stock for product " + product.Name, nil)
	}

	product.Stock -= quantity
	if err := tx.Save(&product).Error; err != nil {
		return err
	}

	return nil
}