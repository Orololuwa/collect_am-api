package handlers

import (
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/enums"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"github.com/Orololuwa/collect_am-api/src/types"
	"gorm.io/gorm"
)

func (repo *Repository) AddCustomer(payload dtos.CreateCustomer, options ...*Extras) (id uint, errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	customer := models.Customer{
		BusinessID: business.ID,
		Status:     enums.EStatus.Active,
		Type:       payload.Type,
		Email:      payload.Email,
		Phone:      payload.Phone,
	}

	if payload.Type == enums.ECustomerType.Individual {
		customer.FirstName = payload.FirstName
		customer.LastName = payload.LastName
	} else {
		customer.Name = payload.Name
	}

	address := models.Address{
		UnitNumber:    payload.UnitNumber,
		AddressLine:   payload.AddressLine,
		City:          payload.City,
		State:         payload.State,
		CountryCode:   payload.CountryCode,
		PostalCode:    payload.PostalCode,
		AddressLineI:  payload.AddressLineI,
		AddressLineII: payload.AddressLineII,
	}

	err := repo.conn.Transaction(func(tx *gorm.DB) error {
		customerId, txErr := repo.Customer.InsertCustomer(customer)
		if txErr != nil {
			return txErr
		}
		id = customerId

		address.CustomerID = customerId
		_, txErr = repo.Address.InsertAddress(address)
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		return id, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return id, errData
}

func (repo *Repository) EditCustomer(payload types.EditCustomerPayload, options ...*Extras) (errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	customer, err := repo.Customer.FindOneById(repository.FindOneBy{ID: payload.Id, BusinessID: business.ID})
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	updatedDetails := models.Customer{
		Phone: payload.Phone,
	}

	if customer.Type == enums.ECustomerType.Individual {
		updatedDetails.FirstName = payload.FirstName
		updatedDetails.LastName = payload.LastName
	} else {
		updatedDetails.Name = payload.Name
	}

	err = repo.Customer.UpdateCustomer(repository.FindOneBy{ID: payload.Id, BusinessID: business.ID}, updatedDetails)
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return errData
}
