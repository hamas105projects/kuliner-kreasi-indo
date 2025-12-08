package saleorder

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleOrder struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CashierID   uuid.UUID `gorm:"not null"`
	TotalAmount float64   `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
	Items       []SaleOrderItem `gorm:"foreignKey:SaleOrderID"`
}

type SaleOrderItem struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	SaleOrderID   uuid.UUID `gorm:"not null"`
	ProductID     uuid.UUID `gorm:"not null"`
	Qty           int       `gorm:"not null"`
	PriceSnapshot float64   `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
