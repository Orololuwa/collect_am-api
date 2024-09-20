package v1

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/helpers/utils"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/types"
	"github.com/go-chi/chi/v5"
)

func CleanInvoiceQuery(queryParams url.Values) map[string]interface{} {
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

func (m *V1) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var body dtos.CreateInvoice
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

	var payload types.CreateInvoicePayload
	payload.CreateInvoice = body

	id, errData := m.Handlers.CreateInvoice(payload, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, id, http.StatusCreated, "invoice created successfully")
}

func (m *V1) GetAllInvoices(w http.ResponseWriter, r *http.Request) {
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

	query := CleanInvoiceQuery(r.URL.Query())

	products, pagination, errData := m.Handlers.GetAllInvoices(query, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriterWithPagination(w, products, pagination, http.StatusCreated, "product updated successfully")
}

func (m *V1) GetInvoice(w http.ResponseWriter, r *http.Request) {
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

	payload := types.GetAnInvoicePayload{}
	payload.Id = uint(productId)

	product, errData := m.Handlers.GetInvoice(payload, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, product, http.StatusCreated, "product updated successfully")
}

func (m *V1) EditInvoice(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	var bodyMap map[string]interface{}
	err = json.Unmarshal([]byte(bodyBytes), &bodyMap)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	body, err := utils.ValidateMap(bodyMap, dtos.InvoiceValidationMap, true)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	log.Println(body)
	// if body["listed_products"] != nil {
	// 	rawListedProducts := body["listed_products"].([]map[string]interface{})
	// 	var newRawListedProducts []map[string]interface{}

	// 	for _, productMap := range rawListedProducts {

	// 		cleanedProductMap, err := utils.ValidateMap(productMap, dtos.ListedProductsValidationMap, true)
	// 		if err != nil {
	// 			helpers.ClientError(w, err, http.StatusBadRequest, "")
	// 			return
	// 		}
	// 		newRawListedProducts = append(newRawListedProducts, cleanedProductMap)
	// 	}

	// 	body["listed_products"] = newRawListedProducts
	// }
	if listedProducts, ok := body["listed_products"].([]interface{}); ok {
		var products []map[string]interface{}
		for _, lp := range listedProducts {
			if product, ok := lp.(map[string]interface{}); ok {
				cleanedProduct, err := utils.ValidateMap(product, dtos.ListedProductsValidationMap, true)
				if err != nil {
					helpers.ClientError(w, err, http.StatusBadRequest, "")
					return
				}
				products = append(products, cleanedProduct)
			} else {
				log.Fatal("Error asserting listedProduct element to map[string]interface{}")
			}
		}
		body["listed_products"] = products
	}
	log.Println(body)

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

	invoiceId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	var payload types.EditInvoicePayload
	payload.Body = body
	payload.ID = uint(invoiceId)

	errData := m.Handlers.EditInvoice(payload, extras)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, nil, http.StatusCreated, "invoice updated successfully")
}
