package dtos

type UserLoginBody struct {
	Email string `json:"email" validate:"required,email" faker:"email"`
}

type UserSignUp struct {
	FirstName string `json:"firstName" validate:"required" faker:"first_name"`
	LastName string `json:"lastName" validate:"required" faker:"last_name"`
	Email string `json:"email" validate:"required,email" faker:"email"`
	Phone string `json:"phone" validate:"required,e164" faker:"e_164_phone_number"`
	Password string `json:"password" validate:"required" faker:"password"`
}