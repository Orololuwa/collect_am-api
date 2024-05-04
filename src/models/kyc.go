package models

import "time"

type KYC struct {
	ID int
	CertificateOfRegistration string
	ProofOfAddress string
	BVN string
	BusinessAddress string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	BusinessId int `db:"business_id"`
}