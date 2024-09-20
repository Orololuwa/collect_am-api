package v1

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/go-chi/chi/v5"
)

func (m *V1) AddBusiness(w http.ResponseWriter, r *http.Request) {
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

	extra := &handlers.Extras{User: user}

	id, errData := m.Handlers.CreateBusiness(body, extra)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, id, http.StatusCreated, "business added successfully")
}

func (m *V1) GetBusiness(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
		return
	}

	businessId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	extra := &handlers.Extras{User: user}
	business, errData := m.Handlers.GetBusiness(uint(businessId), extra)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, &business, http.StatusOK, "business retrieved successfully")
}

func (m *V1) UpdateBusiness(w http.ResponseWriter, r *http.Request) {
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

	// Decode JSON data into a map
	var bodyMap map[string]interface{}
	err = json.Unmarshal([]byte(bodyBytes), &bodyMap)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
		return
	}

	businessId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	extra := &handlers.Extras{User: user}
	errData := m.Handlers.UpdateBusiness(uint(businessId), bodyMap, extra)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, nil, http.StatusCreated, "business updated successfully")
}
