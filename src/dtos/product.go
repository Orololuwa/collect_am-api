package dtos

type AddProduct struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price uint `json:"price" validate:"required"`
	Category string `json:"category" validate:"required"`
	Count uint `json:"count" validate:"required"`
}