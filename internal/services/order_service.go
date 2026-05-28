package services

import (
	"context"
	"log"
	"time"

	"github.com/agussuartawan/project-test-balabali/internal/database"
	"github.com/agussuartawan/project-test-balabali/internal/dto"
	apperrors "github.com/agussuartawan/project-test-balabali/internal/errors"
	"github.com/agussuartawan/project-test-balabali/internal/model"
	"github.com/agussuartawan/project-test-balabali/internal/repositories"
	"github.com/agussuartawan/project-test-balabali/internal/response"
	"gorm.io/gorm"
)

type OrderService struct {
	tm *database.TransactionManager
	orderRepository *repositories.OrderRepository
	productRepository *repositories.ProductRepository
}

func NewOrderService(
	tm *database.TransactionManager,
	orderRepository *repositories.OrderRepository, 
	productRepository *repositories.ProductRepository,
) *OrderService {
	return &OrderService{
		tm: tm,
		orderRepository: orderRepository, 
		productRepository: productRepository,
	}
}

func (s *OrderService) Create(c context.Context, req *dto.OrderRequest) (*response.OrderResponse, error) {
	if err := s.validateOrder(req); err != nil {
		return nil, err
	}

	productIds := s.getProductIds(req.Items)
	products, err := s.productRepository.GetByIds(c, productIds); if err != nil {
		return nil, err
	}

	if len(products) != len(productIds) {
		return nil, apperrors.NewNotFoundError("Some products not found", nil)
	}

	orderItems := make([]model.OrderItem, len(req.Items))
	totalPrice := int64(0)
	for i, item := range req.Items {
		var product model.Product
		for _, p := range products {
			if p.ID == item.ProductID {
				product = p
				break
			}
		}

		orderItems[i] = model.OrderItem{
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			Price: product.Price,
			Subtotal: product.Price * int64(item.Quantity),
		}
		totalPrice += orderItems[i].Subtotal
	}

	order := &model.Order{
		Status: "pending",
		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	if err := s.tm.WithTransaction(c, func(tx *gorm.DB) error {
		// decrement stock atomically with lock
		for _, item := range orderItems {
			if err := s.productRepository.DecrementStock(c, tx, item.ProductID, item.Quantity); err != nil {
				return err
			}
		}

		// create new order
		if err := s.orderRepository.Create(c, tx, order); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// simple background worker to process order using go routine
	go func() {
		time.Sleep(5 * time.Second)
		if err := s.ProcessOrderByID(context.Background(), order.ID); err != nil {
			log.Printf("error processing order %d: %v\n", order.ID, err)
		}
	}()
	
	return s.toOrderResponse(order, products), nil
}

func (s *OrderService) GetAll(c context.Context) ([]*response.OrderResponse, error) {
	orders, err := s.orderRepository.GetAll(c); if err != nil {
		return nil, err
	}

	orderResponses := make([]*response.OrderResponse, len(orders))
	products := make([]model.Product, 0) // ignore order items for order list
	for i, order := range orders {
		orderResponses[i] = s.toOrderResponse(&order, products)
	}

	return orderResponses, nil
}

func (s *OrderService) GetById(c context.Context, id uint) (*response.OrderResponse, error) {
	order, err := s.orderRepository.GetById(c, id); if err != nil {
		return nil, err
	}

	orderItems, err := s.orderRepository.GetOrderItems(c, order.ID); if err != nil {
		return nil, err
	}

	productIDs := make([]uint, len(orderItems))
	for i, item := range orderItems {
		productIDs[i] = item.ProductID
	}
	products, err := s.productRepository.GetByIds(c, productIDs); if err != nil {
		return nil, err
	}

	order.OrderItems = orderItems
	return s.toOrderResponse(order, products), nil
}

func (s *OrderService) ProcessOrderByID(c context.Context, id uint) error {
	order, err := s.orderRepository.GetById(c, id); if err != nil {
		log.Printf("error getting order %d: %v\n", id, err)
		return nil
	}

	if order.Status != "pending" {
		log.Printf("order %d is not pending, skipping", id)
		return nil
	}

	orderIDs := []uint{id}
	if err := s.orderRepository.UpdateStatus(c, orderIDs, "processed"); err != nil {
		log.Printf("error updating status for order %d: %v\n", id, err)
		return nil
	}

	return nil
}

func (s *OrderService) ProcessOrder(c context.Context) error {
	orderIDs, err := s.orderRepository.GetPendingOrders(c); if err != nil {
		return err
	}
	log.Printf("found %d pending orders\n", len(orderIDs))
	if len(orderIDs) == 0 {
		return nil
	}

	// simulate processing time for 5 seconds
	time.Sleep(5 * time.Second)

	// process each order
	if err := s.orderRepository.UpdateStatus(c, orderIDs, "processed"); err != nil {
		return err
	}

	return nil
}

func (s *OrderService) getProductIds(items []dto.OrderItemRequest) []uint {
	productIds := make([]uint, len(items))
	for i, item := range items {
		productIds[i] = item.ProductID
	}
	return productIds
}

func (s *OrderService) validateOrder(order *dto.OrderRequest) error {
	errors := make(map[string][]string)
	if len(order.Items) == 0 {
		errors["items"] = []string{"Items must be contains at least one item"}
	}

	for _, item := range order.Items {
		if item.Quantity <= 0 {
			errors["items"] = []string{"Quantity must be greater than 0"}
		}
	}
	if len(errors) > 0 {
		return apperrors.NewBadRequestError("Validation error", errors)
	}

	return nil
}

func (s *OrderService) toOrderResponse(order *model.Order, products []model.Product) *response.OrderResponse {
	productMap := make(map[uint]model.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	orderItems := make([]response.OrderItemResponse, len(order.OrderItems))
	for i, item := range order.OrderItems {
		orderItems[i] = response.OrderItemResponse{
			Product: response.ProductResponse{
				ID: productMap[item.ProductID].ID,
				Name: productMap[item.ProductID].Name,
				Price: productMap[item.ProductID].Price,
			},
			Quantity: item.Quantity,
			Price: item.Price,
			Subtotal: item.Subtotal,
		}
	}

	return &response.OrderResponse{
		ID:         order.ID,
		Status:     order.Status,
		TotalPrice: order.TotalPrice,
		OrderItems: orderItems,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}