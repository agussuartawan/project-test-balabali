package services

import (
	"context"
	"strings"

	"github.com/agussuartawan/project-test-balabali/internal/dto"
	apperrors "github.com/agussuartawan/project-test-balabali/internal/errors"
	"github.com/agussuartawan/project-test-balabali/internal/model"
	"github.com/agussuartawan/project-test-balabali/internal/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) Create(c context.Context, product *dto.ProductRequest) (*model.Product, error) {
	if err := s.validateProduct(product); err != nil {
		return nil, err
	}

	productModel := &model.Product{
		Name:  strings.TrimSpace(product.Name),
		Stock: product.Stock,
		Price: product.Price,
	}

	if err := s.productRepository.Create(c, productModel); err != nil {
		return nil, err
	}

	return productModel, nil
}

func (s *ProductService) GetAll(c context.Context) ([]model.Product, error) {
	return s.productRepository.GetAll(c)
}

func (s *ProductService) GetById(c context.Context, id uint) (*model.Product, error) {
	return s.productRepository.GetById(c, id)
}

func (s *ProductService) Update(c context.Context, id uint, product *dto.ProductRequest) (*model.Product, error) {
	if err := s.validateProduct(product); err != nil {
		return nil, err
	}

	if !s.productRepository.Exists(c, id) {
		return nil, apperrors.NewNotFoundError("Product not found", nil)
	}

	productModel := &model.Product{
		Name:  strings.TrimSpace(product.Name),
		Stock: product.Stock,
		Price: product.Price,
	}

	if err := s.productRepository.Update(c, id, productModel); err != nil {
		return nil, err
	}

	return productModel, nil
}

func (s *ProductService) Delete(c context.Context, id uint) error {
	if !s.productRepository.Exists(c, id) {
		return apperrors.NewNotFoundError("Product not found", nil)
	}

	return s.productRepository.Delete(c, id)
}

func (s *ProductService) validateProduct(product *dto.ProductRequest) error {
	errors := make(map[string][]string)
	if product.Stock < 0 {
		errors["stock"] = []string{"Stock cannot be negative"}
	}
	if product.Price <= 0 {
		errors["price"] = []string{"Price must be greater than 0"}
	}

	if len(errors) > 0 {
		return apperrors.NewBadRequestError("Validation error", errors)
	}

	return nil
}