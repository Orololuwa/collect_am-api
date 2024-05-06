package models

import "time"

type KYC struct {
	ID int `db:"id"`
	CertificateOfRegistration string `db:"certificate_of_registration"`
	ProofOfAddress string `db:"proof_of_address"`
	BVN string `db:"bvn"`
	BusinessAddress string `db:"business_address"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	BusinessId int `db:"business_id"`
}