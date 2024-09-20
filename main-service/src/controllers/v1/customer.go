package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/types"
	"github.com/go-chi/chi/v5"
)

func CleanCustomersQuery(queryParams url.Values) map[string]interface{} {
	// Initialize the filter map
	filter := make(map[string]interface{})

	if pageStr := queryParams.Get("page"); pageStr != "" {
		filter["page"], _ = strconv.Atoi(pageStr)
	}

	if pageSizeStr := queryParams.Get("pageSize"); pageSizeStr != "" {
		filter["pageSize"], _ = strconv.Atoi(pageSizeStr)
	}

	if emailStr := queryParams.Get("email"); emailStr != "" {
		filter["email"] = emailStr
	}

	if phoneStr := queryParams.Get("phone"); phoneStr != "" {
		filter["phone"] = phoneStr
	}

	return filter
}

func (m *V1) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var body dtos.CreateCustomer
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

	business, ok := r.Context().Value("business").(*models.Business)
	if !ok || business == nil {
		helpers.ClientError(w, errors.New("no business ties"), http.StatusForbidden, "")
		return
	}

	extras := &handlers.Extras{User: user, Business: business}

	id, errData := m.Handlers.AddCustomer(body, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, id, http.StatusCreated, "customer added successfully")
}

func (m *V1) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var body dtos.UpdateCustomer

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

	business, ok := r.Context().Value("business").(*models.Business)
	if !ok || business == nil {
		helpers.ClientError(w, errors.New("no business ties"), http.StatusForbidden, "")
		return
	}

	extras := &handlers.Extras{User: user, Business: business}

	customerId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	payload := types.EditCustomerPayload{
		FindByID:       dtos.FindByID{Id: uint(customerId)},
		UpdateCustomer: body,
	}

	errData := m.Handlers.EditCustomer(payload, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, nil, http.StatusCreated, "customer added successfully")
}

func (m *V1) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
		return
	}

	business, ok := r.Context().Value("business").(*models.Business)
	if !ok || business == nil {
		helpers.ClientError(w, errors.New("no business ties"), http.StatusForbidden, "")
		return
	}

	extras := &handlers.Extras{User: user, Business: business}

	query := CleanCustomersQuery(r.URL.Query())

	products, pagination, errData := m.Handlers.GetAllCustomers(query, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriterWithPagination(w, products, pagination, http.StatusCreated, "product updated successfully")
}

func (m *V1) GetCustomer(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		helpers.ClientError(w, errors.New("unauthorized"), http.StatusUnauthorized, "")
		return
	}

	business, ok := r.Context().Value("business").(*models.Business)
	if !ok || business == nil {
		helpers.ClientError(w, errors.New("no business ties"), http.StatusForbidden, "")
		return
	}

	extras := &handlers.Extras{User: user, Business: business}

	productId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	payload := types.GetACustomerPayload{}
	payload.Id = uint(productId)

	product, errData := m.Handlers.GetCustomer(payload, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, product, http.StatusCreated, "product updated successfully")
}
