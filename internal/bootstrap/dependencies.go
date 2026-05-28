package bootstrap

import (
	"github.com/agussuartawan/project-test-balabali/internal/database"
	"github.com/agussuartawan/project-test-balabali/internal/handlers"
	"github.com/agussuartawan/project-test-balabali/internal/repositories"
	"github.com/agussuartawan/project-test-balabali/internal/services"
	"gorm.io/gorm"
)

type Dependencies struct {
	TransactionManager *database.TransactionManager

	OrderRepository   *repositories.OrderRepository
	ProductRepository *repositories.ProductRepository

	OrderService   *services.OrderService
	ProductService *services.ProductService

	OrderHandler   *handlers.OrderHandler
	ProductHandler *handlers.ProductHandler
}

func InitDependencies(db *gorm.DB) *Dependencies {
	tm := database.NewTransactionManager(db)
	orderRepo := repositories.NewOrderRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderService := services.NewOrderService(tm, orderRepo, productRepo)
	productService := services.NewProductService(productRepo)
	orderHandler := handlers.NewOrderHandler(orderService)
	productHandler := handlers.NewProductHandler(productService)

	return &Dependencies{
		TransactionManager: tm,
		OrderRepository: orderRepo,
		ProductRepository: productRepo,
		OrderService: orderService,
		ProductService: productService,
		OrderHandler: orderHandler,
		ProductHandler: productHandler,
	}
}