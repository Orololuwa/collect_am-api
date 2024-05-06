package models

import "time"

type Business struct {
	ID int `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	Description string `db:"description"`
	Sector string `db:"sector"`
	IsCorporateAffair bool `db:"is_corporate_affairs"`
	IsSetupComplete bool `db:"is_setup_complete"`
	Logo string `db:"logo"`
	UserId int `db:"user_id"`
	KYC KYC //one to one
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}