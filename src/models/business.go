package models

import "time"

type Business struct {
	ID int `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	Description string `db:"description"`
	Sector string `db:"sector"`
	IsCorporateAffair bool `db:"is_corporate_affairs"`
	KYCId int `db:"kyc_id"`
	Logo string `db:"logo"`
	KYC KYC
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}