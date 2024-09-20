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
	"github.com/go-chi/chi/v5"
)

func CleanProductQuery(queryParams url.Values) map[string]interface{} {
	// Initialize the filter map
	filter := make(map[string]interface{})

	if pageStr := queryParams.Get("page"); pageStr != "" {
		filter["page"], _ = strconv.Atoi(pageStr)
	}

	if pageSizeStr := queryParams.Get("pageSize"); pageSizeStr != "" {
		filter["pageSize"], _ = strconv.Atoi(pageSizeStr)
	}

	if codeStr := queryParams.Get("code"); codeStr != "" {
		filter["code"] = codeStr
	}

	return filter
}

func (m *V1) AddProduct(w http.ResponseWriter, r *http.Request) {
	var body dtos.AddProduct
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

	id, errData := m.Handlers.AddProduct(body, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, id, http.StatusCreated, "product added successfully")
}

func (m *V1) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var body dtos.UpdateProduct
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

	productId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}
	body.ID = uint(productId)

	errData := m.Handlers.UpdateProduct(body, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, nil, http.StatusCreated, "product updated successfully")
}

func (m *V1) GetAllProducts(w http.ResponseWriter, r *http.Request) {
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

	// var query types.GetAllProductsPayload

	query := CleanProductQuery(r.URL.Query())

	products, pagination, errData := m.Handlers.GetAllProducts(query, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriterWithPagination(w, products, pagination, http.StatusCreated, "product updated successfully")
}

func (m *V1) GetProduct(w http.ResponseWriter, r *http.Request) {
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

	product, errData := m.Handlers.GetProduct(uint(productId), extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, product, http.StatusCreated, "product updated successfully")
}
