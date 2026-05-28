package handlers

import (
	"net/http"

	"github.com/agussuartawan/project-test-balabali/internal/dto"
	"github.com/agussuartawan/project-test-balabali/internal/errors"
	"github.com/agussuartawan/project-test-balabali/internal/services"
	"github.com/agussuartawan/project-test-balabali/internal/utils"
	"github.com/labstack/echo/v4"

	_ "github.com/agussuartawan/project-test-balabali/internal/response"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create order
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body dto.OrderRequest true "Create Order"
// @Success 201 {object} utils.BaseResponse{data=response.OrderResponse}
// @Router /orders [post]
func (h *OrderHandler) Create(c echo.Context) error {
	var request dto.OrderRequest
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

	order, err := h.orderService.Create(c.Request().Context(), &request); if err != nil {
		return utils.Error(c, err)
	}
	return utils.Success(c, http.StatusCreated, "Order created successfully", order)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} utils.BaseResponse{data=[]response.OrderResponse}
// @Router /orders [get]
func (h *OrderHandler) GetAll(c echo.Context) error {
	orders, err := h.orderService.GetAll(c.Request().Context()); if err != nil {
		return utils.Error(c, err)
	}
	return utils.Success(c, http.StatusOK, "Orders fetched successfully", orders)
}

// GetOrderById godoc
// @Summary Get order by id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} utils.BaseResponse{data=response.OrderResponse}
// @Router /orders/{id} [get]
func (h *OrderHandler) GetById(c echo.Context) error {
	id, err := utils.ParseID(c.Param("id")); if err != nil {
		return utils.Error(c, err)
	}

	order, err := h.orderService.GetById(c.Request().Context(), id); if err != nil {
		return utils.Error(c, err)
	}
	return utils.Success(c, http.StatusOK, "Order fetched successfully", order)
}