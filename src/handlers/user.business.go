package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/serializer"
)

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

	ctx := context.Background()
	err = m.DB.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		businessId, txErr := m.Business.CreateBusiness(ctx, tx, 
			models.Business{ 
				Name: body.Name, 
				Email: body.Email, 
				Description: body.Description,
				Sector: body.Sector,
				IsCorporateAffair: body.IsCorporateAffair,
				Logo: body.Logo,
				UserID: int(user.ID),
				IsSetupComplete: true,
			},
		)
		if txErr != nil {
			return txErr
		}
		id = businessId

		_, txErr = m.KYC.CreateKyc(ctx, tx, 
			models.Kyc{ 
				CertificateOfRegistration: body.CertificateOfRegistration,
				ProofOfAddress: body.ProofOfAddress,
				BVN: body.BVN,
				BusinessID: businessId,
			},
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

	ctx := context.Background()
	business, err := m.Business.GetUserBusiness(ctx, nil, int(user.ID), models.Business{})
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
        return
	}

	if business == nil {
		helpers.ClientResponseWriter(w, business, http.StatusOK, "business retrieved successfully")
		return
	}

	var dst serializer.Business	
	err = helpers.SerializeStruct(business, &dst)
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
        return
	}

	// data := func() any {
	// 	if business == nil {
	// 		return nil
	// 	}
	// 	return dst
	// }()


	helpers.ClientResponseWriter(w, dst, http.StatusOK, "business retrieved successfully")
}

func (m *Repository) UpdateBusiness(w http.ResponseWriter, r *http.Request){
	var body dtos.UpdateBusiness
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

	ctx := context.Background()
	err = m.DB.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		txErr := m.Business.UpdateBusiness(ctx, tx, 
			models.Business{ 
				Name: body.Name, 
				Description: body.Description,
				Sector: body.Sector,
				IsCorporateAffair: body.IsCorporateAffair,
				Logo: body.Logo,
				UserID: int(user.ID),
				IsSetupComplete: true,
			},
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