package models

type KYC struct {
	ID int
	CertificateOfRegistration string
	ProofOfAddress string
	BVN string
	BusinessAddress string
}