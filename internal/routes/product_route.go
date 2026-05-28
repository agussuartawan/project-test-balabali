package routes

import (
	"github.com/agussuartawan/project-test-balabali/internal/handlers"
	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Group, productHandler *handlers.ProductHandler) {
	e.POST("/products", productHandler.Create)
	e.GET("/products", productHandler.GetAll)
	e.GET("/products/:id", productHandler.GetById)
	e.PUT("/products/:id", productHandler.Update)
	e.DELETE("/products/:id", productHandler.Delete)
}