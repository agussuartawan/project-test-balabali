package handlers

import (
	"net/http"

	"github.com/agussuartawan/project-test-balabali/internal/dto"
	"github.com/agussuartawan/project-test-balabali/internal/errors"
	"github.com/agussuartawan/project-test-balabali/internal/services"
	"github.com/agussuartawan/project-test-balabali/internal/utils"
	"github.com/labstack/echo/v4"

	_ "github.com/agussuartawan/project-test-balabali/internal/model"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateUser godoc
//
// @Summary Create product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body dto.ProductRequest true "Create Product"
// @Success 201 {object} utils.BaseResponse{data=model.Product}
// @Router /products [post]
func (h *ProductHandler) Create(c echo.Context) error {
	var request dto.ProductRequest

	if err := c.Bind(&request); err != nil {
		return utils.Error(c, errors.NewBadRequestError("invalid request body", nil))
	}
	if err := c.Validate(&request); err != nil {
		return utils.Error(c, 
			errors.NewBadRequestError(
				"Validation error", 
				utils.FormatValidationError(err, request),
			),
		)
	}

	product, err := h.productService.Create(c.Request().Context(), &request); if err != nil {
		return utils.Error(c, err)
	}

	return utils.Success(c, http.StatusCreated, "Product created successfully", product)
}

// GetAllProducts godoc
// @Summary Get all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} utils.BaseResponse{data=[]model.Product}
// @Router /products [get]
func (h *ProductHandler) GetAll(c echo.Context) error {
	products, err := h.productService.GetAll(c.Request().Context()); if err != nil {
		return utils.Error(c, err)
	}
	return utils.Success(c, http.StatusOK, "Products fetched successfully", products)
}

// GetProductById godoc
// @Summary Get product by id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.BaseResponse{data=model.Product}
// @Router /products/{id} [get]
func (h *ProductHandler) GetById(c echo.Context) error {
	id, err := utils.ParseID(c.Param("id")); if err != nil {
		return utils.Error(c, err)
	}
	
	product, err := h.productService.GetById(c.Request().Context(), id); if err != nil {
		return utils.Error(c, err)
	}
	return utils.Success(c, http.StatusOK, "Product fetched successfully", product)
}

// UpdateProduct godoc
// @Summary Update product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body dto.ProductRequest true "Update Product"
// @Success 200 {object} utils.BaseResponse{data=model.Product}
// @Router /products/{id} [put]
func (h *ProductHandler) Update(c echo.Context) error {
	id, err := utils.ParseID(c.Param("id")); if err != nil {
		return utils.Error(c, err)
	}

	var request dto.ProductRequest
	if err := c.Bind(&request); err != nil {
		return utils.Error(c, errors.NewBadRequestError("invalid request body", nil))
	}
	if err := c.Validate(&request); err != nil {
		return utils.Error(c, 
			errors.NewBadRequestError(
				"Validation error", 
				utils.FormatValidationError(err, request),
			),
		)
	}	

	product, err := h.productService.Update(c.Request().Context(), id, &request); if err != nil {
		return utils.Error(c, err)
	}
	return utils.Success(c, http.StatusOK, "Product updated successfully", product)
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.BaseResponse
// @Router /products/{id} [delete]
func (h *ProductHandler) Delete(c echo.Context) error {
	id, err := utils.ParseID(c.Param("id")); if err != nil {
		return utils.Error(c, err)
	}
	
	err = h.productService.Delete(c.Request().Context(), id); if err != nil {
		return utils.Error(c, err)
	}
	return utils.Success(c, http.StatusOK, "Product deleted successfully", nil)
}