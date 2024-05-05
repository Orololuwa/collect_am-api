package dtos

type AddBusiness struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Description string `json:"description" validate:"required,containsany"`
	Sector string `json:"sector" validate:"required"`
	IsCorporateAffair bool `json:"isCorporateAffairs" validate:"required,boolean"`
	Logo string `json:"logo" validate:"url"`
	CertificateOfRegistration string `json:"certificateOfRegistration" validate:"required,url"`
	ProofOfAddress string `json:"proof_of_address" validate:"required,url"`
	BVN string `json:"bvn" validate:"required"`
}