package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
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

func (m *Repository) AddBusiness(w http.ResponseWriter, r *http.Request){
	var body dtos.AddBusiness
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	err = m.App.Validate.Struct(body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}


	user, ok := r.Context().Value("user").(*models.User)
    if !ok || user == nil {
		helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
        return
    }

	var id int

	err = m.conn.Transaction(func(tx *gorm.DB) error {
		businessId, txErr := m.Business.InsertBusiness(
			models.Business{ 
				Name: body.Name, 
				Email: body.Email, 
				Description: body.Description,
				Sector: body.Sector,
				IsCorporateAffair: body.IsCorporateAffair,
				Logo: body.Logo,
				UserID: int(user.ID),
			},
			tx,
		)
		if txErr != nil {
			return txErr
		}
		id = int(businessId)

		_, txErr = m.Kyc.InsertKyc(
			models.Kyc{ 
				CertificateOfRegistration: body.CertificateOfRegistration,
				ProofOfAddress: body.ProofOfAddress,
				BVN: body.BVN,
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
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	helpers.ClientResponseWriter(w, id, http.StatusCreated, "business added successfully")
}

func (m *Repository) GetBusiness(w http.ResponseWriter, r *http.Request){		
	user, ok := r.Context().Value("user").(*models.User)
    if !ok || user == nil {
		helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
        return
    }

	// ctx := context.Background()
	business, err := m.Business.GetOneByUserId(user.ID)
	if err != nil && err.Error() != "record not found"{
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	if err != nil && err.Error() == "record not found"{
		helpers.ClientResponseWriter(w, nil, http.StatusOK, "business retrieved successfully")
		return
	}

	// Check if Kyc struct is empty and set it to nil if it is
	if business.Kyc != nil && reflect.DeepEqual(*business.Kyc, models.Kyc{}) {
		business.Kyc = nil
	}

	log.Println(business)

	helpers.ClientResponseWriter(w, business, http.StatusOK, "business retrieved successfully")
}

func (m *Repository) UpdateBusiness(w http.ResponseWriter, r *http.Request){
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	// Decode JSON data into a struct
	var bodyStruct dtos.UpdateBusiness
    err = json.Unmarshal([]byte(bodyBytes), &bodyStruct)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	
	// validate struct
	err = m.App.Validate.Struct(bodyStruct)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
    if !ok || user == nil {
		helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
        return
    }

	business, err := m.Business.GetOneByUserId(user.ID)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

    // Decode JSON data into a map
	var bodyMap map[string]interface{}
    err = json.Unmarshal([]byte(bodyBytes), &bodyMap)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

    businessData := cleanBusinessData(bodyMap)
    kycData := cleanKycData(bodyMap)

	err = m.conn.Transaction(func(tx *gorm.DB) error {
		txErr := m.Business.UpdateBusiness(
			businessData,
			models.Business{
				ID: business.ID,
				UserID: int(user.ID),
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
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	helpers.ClientResponseWriter(w, nil, http.StatusCreated, "business updated successfully")
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

func (m *Repository) GetBusinessV2(options ...*Extras)(data *models.Business, errData *ErrorData){	
	var user models.User
	if len(options) > 0 && options[0] != nil {
		user = *options[0].User
	}

	business, err := m.Business.GetOneByUserId(user.ID)
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

func (m *Repository) UpdateBusinessV2(payload map[string]interface{}, options ...*Extras)(errData *ErrorData){
	var user models.User
	if len(options) > 0 && options[0] != nil {
		user = *options[0].User
	}

	business, err := m.Business.GetOneByUserId(user.ID)
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
				UserID: int(user.ID),
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