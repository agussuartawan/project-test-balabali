package response

import "time"

type OrderResponse struct {
	ID         uint                `json:"id"`
	Status     string              `json:"status"`
	TotalPrice int64               `json:"totalPrice"`
	OrderItems []OrderItemResponse `json:"orderItems,omitempty"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
}

type OrderItemResponse struct {
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
	Price    int64           `json:"price"`
	Subtotal int64           `json:"subtotal"`
}

type ProductResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}