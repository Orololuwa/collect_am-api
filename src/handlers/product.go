package handlers

import (
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/enums"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
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

func (m *Repository) GetAllProducts(query repository.FilterQueryPagination, options ...*Extras) (products []models.Product, pagination repository.Pagination, errData *ErrorData) {
	products, pagination, err := m.Product.FindAllWithPagination(query)
	if err != nil {
		return products, pagination, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return products, pagination, nil
}

func (m *Repository) GetProduct(id uint, options ...*Extras) (product models.Product, errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	products, err := m.Product.FindOneById(repository.FindOneBy{ID: id, BusinessId: business.ID})
	if err != nil {
		return products, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return products, nil
}
