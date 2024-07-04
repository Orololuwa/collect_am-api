package models

import "time"

type Address struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
	UnitNumber    string     `json:"unitNumber"`
	AddressLine   string     `json:"addressLine"`
	City          string     `json:"city"`
	State         string     `json:"state"`
	CountryCode   string     `json:"countryCode"`
	PostalCode    string     `json:"postalCode"`
	AddressLineI  string     `json:"addressLineI"`
	AddressLineII string     `json:"addressLineII"`
	Customer      Customer
}
