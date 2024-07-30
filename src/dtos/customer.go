package dtos

import "github.com/Orololuwa/collect_am-api/src/enums"

type CreateCustomer struct {
	Type      enums.ICustomerType `json:"type" validate:"required" faker:"oneof: individual, corporate"`
	FirstName string              `json:"firstName" faker:"first_name"`
	LastName  string              `json:"lastName" faker:"last_name"`
	Name      string              `json:"name" faker:"name"`
	Email     string              `json:"email" validate:"required,email" faker:"email"`
	Phone     string              `json:"phone" validate:"required,e164" faker:"phone_number"`
	CreateAddress
}

type UpdateCustomer struct {
	FirstName string `json:"firstName" faker:"first_name"`
	LastName  string `json:"lastName" faker:"last_name"`
	Name      string `json:"name" faker:"name"`
	Phone     string `json:"phone" validate:"e164, omitempty" faker:"phone_number"`
}
