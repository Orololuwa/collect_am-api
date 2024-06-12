package models

import "time"

type Kyc struct {
	ID                     uint       `json:"id"`
	CreatedAt              time.Time  `json:"createdAt"`
	UpdatedAt              time.Time  `json:"updatedAt"`
	DeletedAt              *time.Time `json:"deletedAt,omitempty"`
	CertificateOfRegistration string  `json:"certificateOfRegistration"`
	ProofOfAddress         string     `json:"proofOfAddress"`
	BVN                    string     `json:"bvn"`
	BusinessAddress        string     `json:"businessAddress"`
	BusinessID             uint       `json:"businessId"`
}