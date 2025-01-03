package models

import (
	"time"

	"github.com/Orololuwa/collect_am-api/src/enums"
)

type Customer struct {
	ID         uint                `json:"id"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
	DeletedAt  *time.Time          `json:"deletedAt,omitempty"`
	BusinessID uint                `json:"businessId"`
	Status     enums.IStatus       `gorm:"default:'active'" json:"status"`
	Type       enums.ICustomerType `json:"type" gorm:"not null"`
	FirstName  string              `json:"firstName"`
	LastName   string              `json:"lastName"`
	Name       string              `json:"name"`
	Email      string              `json:"email" gorm:"not null"`
	Phone      string              `json:"phone" gorm:"not null"`
	Address    Address             `json:"address"`
	Invoices   []Invoice           `json:"invoices"`
}
