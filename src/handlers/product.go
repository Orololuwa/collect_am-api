package handlers

import (
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/enums"
	"github.com/Orololuwa/collect_am-api/src/models"
)

func (m *Repository) AddProduct(payload dtos.AddProduct, options ...*Extras) (id uint, errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	product := models.Product{
		Category:    payload.Category,
		Code:        payload.Code,
		Count:       payload.Count,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Status:      enums.ProductStatuses.Active,
		BusinessID:  business.ID,
	}

	id, err := m.Product.InsertProduct(product)
	if err != nil {
		return id, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return id, errData
}

func (m *Repository) UpdateProduct(payload dtos.UpdateProduct, options ...*Extras) (errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	product := models.Product{
		Category:    payload.Category,
		Count:       payload.Count,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Status:      enums.ProductStatuses.Active,
		BusinessID:  business.ID,
	}

	err := m.Product.UpdateProduct(models.Product{ID: payload.ID}, product)
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return errData
}
