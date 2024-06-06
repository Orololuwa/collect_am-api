package models

import (
	"gorm.io/gorm"
)

type Kyc struct {

	gorm.Model
	// ID int `db:"id"`
	CertificateOfRegistration string `db:"certificate_of_registration"`
	ProofOfAddress string `db:"proof_of_address"`
	BVN string `db:"bvn"`
	BusinessAddress string `db:"business_address"`
	// CreatedAt time.Time `db:"created_at"`
	// UpdatedAt time.Time `db:"updated_at"`
	BusinessID int `db:"business_id"`
}