package models

import (
	"time"

	"github.com/Orololuwa/collect_am-api/src/enums"
)

type Product struct {
	ID          uint                `json:"id"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
	DeletedAt   *time.Time          `json:"deletedAt,omitempty"`
	Code        string              `gorm:"not null;unique" json:"code"`
	Name        string              `gorm:"not null" json:"name"`
	Description string              `gorm:"type:varchar(255);not null" json:"description"`
	Price       uint                `gorm:"not null"  json:"price"`
	Category    string              `gorm:"default:others"  json:"category"`
	Count       uint                `json:"count"`
	BusinessID  uint                `json:"businessId"`
	Status      enums.ProductStatus `gorm:"default:'active'" json:"status"`
}
