package entity

import (
	"time"
)

type Product struct {
	BaseEntity
	Name            string     `json:"name" gorm:"type:varchar(255);not null"`
	Description     string     `json:"description" gorm:"type:text"`
	SKU             string     `json:"sku" gorm:"type:varchar(50);unique;not null"`
	CategoryID      uint       `json:"category_id" gorm:"not null"`
	Category        Category   `json:"category" gorm:"foreignKey:CategoryID"`
	Brand           string     `json:"brand" gorm:"type:varchar(100)"`
	Unit            string     `json:"unit" gorm:"type:varchar(20);not null"` // e.g., "tablet", "ml", "mg"
	BasePrice       float64    `json:"base_price" gorm:"type:decimal(10,2);not null"`
	Currency        string     `json:"currency" gorm:"type:varchar(3);not null;default:'USD'"`
	StockQuantity   int        `json:"stock_quantity" gorm:"type:int;not null;default:0"`
	MinOrderUnit    int        `json:"min_order_unit" gorm:"type:int;not null;default:1"`
	MaxOrderUnit    int        `json:"max_order_unit" gorm:"type:int"`
	IsActive        bool       `json:"is_active" gorm:"not null;default:true"`
	ExpiryDate      *time.Time `json:"expiry_date" gorm:"type:date"`
	ManufactureDate *time.Time `json:"manufacture_date" gorm:"type:date"`
	ImageURL        string     `json:"image_url" gorm:"type:varchar(255)"`

	// Pricing tiers for bulk purchases
	PricingTiers []PricingTier `gorm:"foreignKey:ProductID"`
}

// PricingTier represents different price points based on quantity
type PricingTier struct {
	BaseEntity
	ProductID    uint    `json:"product_id" gorm:"not null"`
	MinQuantity  int     `json:"min_quantity" gorm:"type:int;not null"`
	MaxQuantity  int     `json:"max_quantity" gorm:"type:int"`
	PricePerUnit float64 `json:"price_per_unit" gorm:"type:decimal(10,2);not null"`
}
