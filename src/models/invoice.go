package models

import (
	"time"

	"github.com/Orololuwa/collect_am-api/src/enums"
)

type Invoice struct {
	ID             uint                 `json:"id"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
	DeletedAt      *time.Time           `json:"deletedAt,omitempty"`
	Code           string               `gorm:"not null" json:"code"`
	Description    string               `gorm:"not null" json:"description"`
	DueDate        time.Time            `json:"dueDate"`
	Price          float64              `gorm:"not null,default:0"  json:"price"`
	Status         enums.IInvoiceStatus `gorm:"default:'draft'" json:"status"`
	Tax            float64              `gorm:"not null,default:0"  json:"tax"`
	ServiceCharge  float64              `gorm:"not null,default:0"  json:"serviceCharge"`
	Discount       float64              `gorm:"not null,default:0"  json:"discount"`
	DiscountType   enums.IDiscountType  `gorm:"not null,default:'fixed'"  json:"discountType"`
	Total          float64              `gorm:"not null,default:0"  json:"total"`
	BusinessID     uint                 `json:"businessId"`
	ListedProducts []ListedProduct      `json:"listedProducts"`
	CustomerID     uint                 `json:"customerId"`
}
