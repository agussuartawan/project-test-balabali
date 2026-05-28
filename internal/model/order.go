package model

type Order struct {
	ID         uint   `gorm:"primaryKey"`
	Status     string `gorm:"not null"`
	TotalPrice int64  `gorm:"not null"`

	BaseModel
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}