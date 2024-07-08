package dtos

type AddProduct struct {
	Code        string `json:"code" validate:"required" faker:"sentence"`
	Name        string `json:"name" validate:"required" faker:"name"`
	Description string `json:"description" validate:"required" faker:"sentence"`
	Price       uint   `json:"price" validate:"required" faker:"boundary_start=1, boundary_end=1000"`
	Category    string `json:"category" validate:"required" faker:"sentence"`
	Count       uint   `json:"count" validate:"required" faker:"boundary_start=1, boundary_end=100"`
}

type UpdateProduct struct {
	ID          uint   `json:"id,omitempty" validate:"omitempty" faker:"-"`
	Name        string `json:"name,omitempty" validate:"omitempty" faker:"name"`
	Description string `json:"description,omitempty" validate:"omitempty" faker:"sentence"`
	Price       uint   `json:"price,omitempty" validate:"omitempty" faker:"boundary_start=1, boundary_end=1000"`
	Category    string `json:"category,omitempty" validate:"omitempty" faker:"sentence"`
	Count       uint   `json:"count,omitempty" validate:"omitempty" faker:"boundary_start=1, boundary_end=100"`
}

type GetAllProductsFilter struct {
	BusinessID uint
	Code       string `json:"code" faker:"sentence"`
}
