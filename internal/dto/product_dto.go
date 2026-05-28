package dto

type ProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
	Price int64  `json:"price" validate:"required"`
}