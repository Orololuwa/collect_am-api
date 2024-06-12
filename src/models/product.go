package models

import "time"

type Product struct {
	ID                     uint       `json:"id"`
	CreatedAt              time.Time  `json:"createdAt"`
	UpdatedAt              time.Time  `json:"updatedAt"`
	DeletedAt              *time.Time `json:"deletedAt,omitempty"`
	Code string `gorm:"not null;unique"`
	Name string `gorm:"not null"`
	Description string `gorm:"type:varchar(255);not null"`
	Price uint `gorm:"not null"`
	Category string `gorm:"default:others"`
	Count uint
	BusinessID uint
}