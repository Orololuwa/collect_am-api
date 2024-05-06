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
)

func (m *Repository) AddBusiness(w http.ResponseWriter, r *http.Request){
	var body dtos.AddBusiness
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "internal server error")
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
				IsCorporateAffair: body.IsCorporateAffair || false,
				Logo: body.Logo,
				UserId: user.ID,
			},
		)
		if txErr != nil {
			return txErr
		}
		id = businessId

		_, txErr = m.KYC.CreateKyc(ctx, tx, 
			models.KYC{ 
				CertificateOfRegistration: body.CertificateOfRegistration,
				ProofOfAddress: body.ProofOfAddress,
				BVN: body.BVN,
				BusinessId: businessId,
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
	business, err := m.Business.GetUserBusiness(ctx, nil, user.ID, models.Business{})
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
        return
	}

	helpers.ClientResponseWriter(w, business, http.StatusOK, "business retrieved successfully")
}