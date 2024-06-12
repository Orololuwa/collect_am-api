package dtos

type AddBusiness struct {
	Name string `json:"name" validate:"required" faker:"name"`
	Email string `json:"email" validate:"required,email" faker:"email"`
	Description string `json:"description" validate:"required" faker:"sentence"`
	Sector string `json:"sector" validate:"required" faker:"name"`
	IsCorporateAffair bool `json:"isCorporateAffair" validate:"required"`
	Logo string `json:"logo" validate:"url" faker:"url"`
	CertificateOfRegistration string `json:"certificateOfRegistration" validate:"required,url" faker:"url"`
	ProofOfAddress string `json:"proofOfAddress" validate:"required,url" faker:"url"`
	BVN string `json:"bvn" validate:"required" faker:"toll_free_number"`
}

type UpdateBusiness struct {
	Name string `json:"name,omitempty" validate:"omitempty" faker:"name"`
	Description string `json:"description,omitempty" validate:"omitempty" faker:"sentence"`
	Sector string `json:"sector,omitempty" validate:"omitempty" faker:"name"`
	IsCorporateAffair bool `json:"is_corporate_affair,omitempty" validate:"omitempty"`
	Logo string `json:"logo,omitempty" validate:"omitempty,url" faker:"url"`
	CertificateOfRegistration string `json:"certificate_of_registration,omitempty" validate:"omitempty,url" faker:"url"`
	ProofOfAddress string `json:"proof_of_address,omitempty" validate:"omitempty,url" faker:"url"`
	BVN string `json:"bvn" validate:"omitempty" faker:"toll_free_number"`
}