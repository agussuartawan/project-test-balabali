package main

import (
	"net/http"

	"github.com/agussuartawan/project-test-balabali/internal/bootstrap"
	"github.com/agussuartawan/project-test-balabali/internal/config"
	"github.com/agussuartawan/project-test-balabali/internal/routes"
	"github.com/agussuartawan/project-test-balabali/internal/scheduler"
	"github.com/agussuartawan/project-test-balabali/internal/utils"
	"github.com/agussuartawan/project-test-balabali/internal/validators"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/agussuartawan/project-test-balabali/docs"
)

// @title Project Test Balabali Backend API
// @version 1.0
// @description Project Test Balabali Backend API
// @host localhost:8080
// @BasePath /api
func main() {
	db := config.ConnectDB()
	deps := bootstrap.InitDependencies(db)

	cronManager := scheduler.NewCron()
	scheduler.RegisterOrderCron(cronManager, deps)
	cronManager.Start()

	e := echo.New()
	e.Validator = validators.NewCustomValidator()

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/health")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, utils.BaseResponse{
			Success: true,
			Message: "API is running",
			Data:    nil,
		})
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.Setup(e, deps)
	
	e.Logger.Fatal(e.Start(":8080"))
}