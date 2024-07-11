package models

import "time"

type ListedProduct struct {
	ID             uint       `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`
	PriceListed    uint       `gorm:"not null,default:0" json:"priceListed"`
	QuantityListed uint       `gorm:"not null,default:0" json:"quantityListed"`
	InvoiceID      uint       `json:"invoiceId"`
	ProductID      uint       `json:"productId"`
}
