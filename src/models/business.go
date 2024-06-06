package models

import (
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	// ID int `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	Description string `db:"description"`
	Sector string `db:"sector"`
	IsCorporateAffair string `db:"is_corporate_affairs" dataType:"bool"`
	IsSetupComplete bool `db:"is_setup_complete"`
	Logo string `db:"logo"`
	UserID int `db:"user_id"`
	Kyc Kyc//one to one
	Products []Product
	// CreatedAt time.Time `db:"created_at"`
	// UpdatedAt time.Time `db:"updated_at"`
}