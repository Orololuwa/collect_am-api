package dtos

type CreateAddress struct {
	UnitNumber    string `json:"unitNumber" validate:"required" faker:"oneof: 6b"`
	AddressLine   string `json:"addressLine" validate:"required" faker:"sentence"`
	City          string `json:"city" validate:"required" faker:"oneof: Ikeja"`
	State         string `json:"state" validate:"required" faker:"oneof: Lagos"`
	CountryCode   string `json:"countryCode" validate:"required,len=2" faker:"oneof: NG"`
	PostalCode    string `json:"postalCode" validate:"required" faker:"oneof: 100213"`
	AddressLineI  string `json:"addressLineI" faker:"-"`
	AddressLineII string `json:"addressLineII" faker:"-"`
}
