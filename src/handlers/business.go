package handlers

import (
	"net/http"
	"reflect"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/models"
	"gorm.io/gorm"
)

// cleanBusinessData cleans business data for a update operation
func cleanBusinessData(bodyMap map[string]interface{}) map[string]interface{} {
    businessData := make(map[string]interface{})
    keys := []string{"name", "description", "sector", "is_corporate_affair", "logo"}
    helpers.AssignIfExists(bodyMap, businessData, keys...)
    return businessData
}

// cleanKycData cleans kyc data for a update operation
func cleanKycData(bodyMap map[string]interface{}) map[string]interface{} {
    kycData := make(map[string]interface{})
    keys := []string{"certificate_of_registration", "proof_of_address", "bvn"}
    helpers.AssignIfExists(bodyMap, kycData, keys...)
    return kycData
}

func (m *Repository) CreateBusiness(payload dtos.AddBusiness, options ...*Extras)(id uint, errData *ErrorData){
	var user models.User
	if len(options) > 0 && options[0] != nil {
		user = *options[0].User
	} 

	err := m.conn.Transaction(func(tx *gorm.DB) error {
		businessId, txErr := m.Business.InsertBusiness(
			models.Business{ 
				Name: payload.Name, 
				Email: payload.Email, 
				Description: payload.Description,
				Sector: payload.Sector,
				IsCorporateAffair: payload.IsCorporateAffair,
				Logo: payload.Logo,
				UserID: int(user.ID),
			},
			tx,
		)
		if txErr != nil {
			return txErr
		}
		id = businessId

		_, txErr = m.Kyc.InsertKyc(
			models.Kyc{ 
				CertificateOfRegistration: payload.CertificateOfRegistration,
				ProofOfAddress: payload.ProofOfAddress,
				BVN: payload.BVN,
				BusinessID: uint(businessId),
			},
			tx,
		)
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		return id, &ErrorData{ Error: err, Status: http.StatusBadRequest}
	}
	return id, nil
}

func (m *Repository) GetBusiness(businessId uint, options ...*Extras)(data *models.Business, errData *ErrorData){	
	business, err := m.Business.GetOneById(businessId)
	if err != nil && err.Error() != "record not found"{
		return &business, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}
	if err != nil && err.Error() == "record not found"{
		return nil, nil
	}

	// Check if Kyc struct is empty and set it to nil if it is
	if business.Kyc != nil && reflect.DeepEqual(*business.Kyc, models.Kyc{}) {
		business.Kyc = nil
	}

	return &business, nil
}

func (m *Repository) UpdateBusiness(id uint, payload map[string]interface{}, options ...*Extras)(errData *ErrorData){
	business, err := m.Business.GetOneById(id)
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}


    businessData := cleanBusinessData(payload)
    kycData := cleanKycData(payload)

	err = m.conn.Transaction(func(tx *gorm.DB) error {
		txErr := m.Business.UpdateBusiness(
			businessData,
			models.Business{
				ID: business.ID,
			},
			tx,
		)
		if txErr != nil {
			return txErr
		}

		txErr = m.Kyc.UpdateKyc(
			kycData,
			models.Kyc{
				BusinessID: business.ID,
			},
			tx,
		)
		if txErr != nil {
			return txErr
		}

		return nil

	})
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return nil
}