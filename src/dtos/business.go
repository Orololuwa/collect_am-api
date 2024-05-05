package dtos

type AddBusiness struct {
	Name string `json:"name" validate:"required" faker:"name"`
	Email string `json:"email" validate:"required,email" faker:"email"`
	Description string `json:"description" validate:"required" faker:"sentence"`
	Sector string `json:"sector" validate:"required" faker:"name"`
	IsCorporateAffair bool `json:"isCorporateAffair" validate:"boolean"`
	Logo string `json:"logo" validate:"url" faker:"url"`
	CertificateOfRegistration string `json:"certificateOfRegistration" validate:"required,url" faker:"url"`
	ProofOfAddress string `json:"proof_of_address" validate:"required,url" faker:"url"`
	BVN string `json:"bvn" validate:"required" faker:"toll_free_number"`
}