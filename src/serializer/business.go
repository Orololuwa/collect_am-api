package serializer

import (
	"time"
)

type Business struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Description string `json:"description"`
	Sector string `json:"sector"`
	IsCorporateAffair bool `json:"is_corporate_affairs" dataType:"bool"`
	IsSetupComplete bool `json:"is_setup_complete"`
	Logo string `json:"logo"`
	UserId int `json:"-"`
	KYC KYC `json:"kyc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type KYC struct {
	ID int `json:"id"`
	CertificateOfRegistration string `json:"certificate_of_registration"`
	ProofOfAddress string `json:"proof_of_address"`
	BVN string `json:"bvn"`
	BusinessAddress string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BusinessId int `json:"-"`
}