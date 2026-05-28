package model

type OrderItem struct {
	ID        uint  `gorm:"primaryKey"`
	OrderID   uint  `gorm:"not null"`
	ProductID uint  `gorm:"not null"`
	Quantity  int   `gorm:"not null"`
	Price     int64 `gorm:"not null"`
	Subtotal  int64 `gorm:"not null"`

	BaseModel
	Product Product `gorm:"foreignKey:ProductID"`
}