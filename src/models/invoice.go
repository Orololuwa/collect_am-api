package models

import (
	"time"

	"github.com/Orololuwa/collect_am-api/src/enums"
)

type Invoice struct {
	ID             uint            `json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeletedAt      *time.Time      `json:"deletedAt,omitempty"`
	Code           string          `gorm:"not null;unique" json:"code"`
	DueDate        time.Time       `json:"dusDate"`
	Price          uint            `gorm:"not null,default:0"  json:"price"`
	Status         enums.IStatus   `gorm:"default:'active'" json:"status"`
	Tax            uint            `gorm:"not null,default:0"  json:"tax"`
	ServiceCharge  uint            `gorm:"not null,default:0"  json:"serviceCharge"`
	Discount       uint            `gorm:"not null,default:0"  json:"discount"`
	BusinessID     uint            `json:"businessId"`
	ListedProducts []ListedProduct `json:"listedProducts"`
	CustomerID     uint            `json:"customerId"`
}
