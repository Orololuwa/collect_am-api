package dtos

import (
	"github.com/Orololuwa/collect_am-api/src/enums"
)

type CreateListedProduct struct {
	QuantityListed uint `json:"quantityListed" validate:"required,min=1" faker:"boundary_start=1, boundary_end=20"`
	ProductID      uint `json:"productId" validate:"required" faker:"boundary_start=1, boundary_end=1000"`
}

type CreateInvoice struct {
	Code           string                `json:"code" validate:"required" faker:"word"`
	Description    string                `json:"description" validate:"required" faker:"sentence"`
	DueDate        string                `json:"dueDate" faker:"date"`
	Tax            float64               `json:"tax" validate:"omitempty,min=0" faker:"boundary_start=1, boundary_end=5"`
	ServiceCharge  float64               `json:"serviceCharge" validate:"omitempty,min=0" faker:"boundary_start=1, boundary_end=5"`
	Discount       float64               `json:"discount" validate:"omitempty,min=0" faker:"boundary_start=1, boundary_end=20"`
	DiscountType   enums.IDiscountType   `json:"discountType" validate:"omitempty" faker:"oneof: fixed, percentage"`
	ListedProducts []CreateListedProduct `json:"listedProducts" validate:"required,dive"`
	CustomerID     uint                  `json:"customerId" validate:"required" faker:"boundary_start=1, boundary_end=1000"`
}
