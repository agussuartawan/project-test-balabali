package routes

import (
	"github.com/agussuartawan/project-test-balabali/internal/bootstrap"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, deps *bootstrap.Dependencies) {
	api := e.Group("/api")
	
	DocsRoute(e)
	ProductRoute(api, deps.ProductHandler)
	OrderRoute(api, deps.OrderHandler)
}