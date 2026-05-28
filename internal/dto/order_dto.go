package dto

type OrderRequest struct {
	Items []OrderItemRequest `json:"items" validate:"required,dive"`
}

type OrderItemRequest struct {
	ProductID uint `json:"productId" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}