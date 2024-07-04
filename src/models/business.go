package models

import "time"

type Business struct {
	ID                uint       `json:"id"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	DeletedAt         *time.Time `json:"deletedAt,omitempty"`
	Name              string     `json:"name"`
	Email             string     `json:"email" gorm:"not null;unique"`
	Description       string     `json:"description"`
	Sector            string     `json:"sector"`
	IsCorporateAffair bool       `json:"isCorporateAffair"`
	IsSetupComplete   bool       `json:"isSetupComplete"`
	Logo              string     `json:"logo"`
	UserID            int        `json:"userId"`
	Kyc               *Kyc       `json:"kyc"`
	Products          []Product  `json:"products"`
	Customers         []Customer `json:"customers"`
}
