package dtos

import (
	"reflect"

	"github.com/Orololuwa/collect_am-api/src/enums"
	"github.com/Orololuwa/collect_am-api/src/helpers/utils"
)

type CreateListedProduct struct {
	QuantityListed uint `json:"quantityListed" validate:"required,min=1" faker:"boundary_start=1, boundary_end=20"`
	ProductID      uint `json:"productId" validate:"required" faker:"boundary_start=1, boundary_end=1000"`
}

type CreateInvoice struct {
	Code           string                `json:"code" validate:"required" faker:"word"`
	Description    string                `json:"description" validate:"required" faker:"sentence"`
	DueDate        string                `json:"dueDate" faker:"date"`
	Tax            float64               `json:"tax" validate:"omitempty,min=0,max=99" faker:"boundary_start=1, boundary_end=5"`
	ServiceCharge  float64               `json:"serviceCharge" validate:"omitempty,min=0,max=99" faker:"boundary_start=1, boundary_end=5"`
	Discount       float64               `json:"discount" validate:"omitempty,discount" faker:"boundary_start=1, boundary_end=20"`
	DiscountType   enums.IDiscountType   `json:"discountType" validate:"omitempty" faker:"oneof: fixed, percentage"`
	ListedProducts []CreateListedProduct `json:"listedProducts" validate:"required,dive"`
	CustomerID     uint                  `json:"customerId" validate:"required" faker:"boundary_start=1, boundary_end=1000"`
}

type EditInvoice struct {
	Description   string              `json:"description" validate:"omitempty" faker:"sentence"`
	DueDate       string              `json:"dueDate" faker:"date"`
	Tax           float64             `json:"tax" validate:"omitempty,min=0,max=99" faker:"boundary_start=1, boundary_end=5"`
	ServiceCharge float64             `json:"serviceCharge" validate:"omitempty,min=0,max=99" faker:"boundary_start=1, boundary_end=5"`
	Discount      float64             `json:"discount" validate:"omitempty,discount" faker:"boundary_start=1, boundary_end=20"`
	DiscountType  enums.IDiscountType `json:"discountType" validate:"omitempty" faker:"oneof: fixed, percentage"`
	CustomerID    uint                `json:"customerId" validate:"omitempty" faker:"boundary_start=1, boundary_end=1000"`
}

var InvoiceValidationMap = map[string]utils.FieldInfo{
	"description":    {reflect.String},
	"dueDate":        {reflect.String},
	"tax":            {reflect.Float64},
	"discount":       {reflect.Float64},
	"discountType":   {reflect.String},
	"serviceCharge":  {reflect.Float64},
	"customerId":     {reflect.Float64},
	"listedProducts": {reflect.Slice},
}
var ListedProductsValidationMap = map[string]utils.FieldInfo{
	"id":          {reflect.Float64},
	"quantity":    {reflect.Float64},
	"priceListed": {reflect.Float64},
}
