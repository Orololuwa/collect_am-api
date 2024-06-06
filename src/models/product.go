package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code string `gorm:"not null;unique"`
	Name string `gorm:"not null"`
	Description string `gorm:"type:varchar(255);not null"`
	Price uint `gorm:"not null"`
	Category string `gorm:"default:others"`
	BusinessID uint
}