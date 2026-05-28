package routes

import (
	"github.com/agussuartawan/project-test-balabali/internal/handlers"
	"github.com/labstack/echo/v4"
)

func OrderRoute(e *echo.Group, orderHandler *handlers.OrderHandler) {
	e.POST("/orders", orderHandler.Create)
	e.GET("/orders", orderHandler.GetAll)
	e.GET("/orders/:id", orderHandler.GetById)
}