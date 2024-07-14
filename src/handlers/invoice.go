package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Orololuwa/collect_am-api/src/enums"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"github.com/Orololuwa/collect_am-api/src/types"
	"gorm.io/gorm"
)

func (repo *Repository) CreateInvoice(payload types.CreateInvoicePayload, options ...*Extras) (id uint, errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	dd := payload.DueDate
	dueDate, err := time.Parse("2006-01-02", dd)
	if err != nil {
		return id, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	codeExists, err := repo.Invoice.FindOneBy(models.Invoice{Code: payload.Code, BusinessID: business.ID})
	if err != nil && err.Error() != "record not found" {
		return id, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}
	if codeExists.ID > 0 {
		return id, &ErrorData{Error: errors.New("invoice with this code exists for your business"), Status: http.StatusBadRequest}
	}

	_, err = repo.Customer.FindOneById(repository.FindOneBy{ID: payload.CustomerID, BusinessID: business.ID})
	if err != nil {
		return id, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	invoice := models.Invoice{
		Code:          payload.Code,
		Description:   payload.Description,
		DueDate:       dueDate,
		Status:        enums.EInvoiceStatus.Pending,
		Tax:           payload.Tax,
		ServiceCharge: payload.ServiceCharge,
		Discount:      payload.Discount,
		DiscountType:  payload.DiscountType,
		CustomerID:    payload.CustomerID,
		BusinessID:    business.ID,
	}

	var listedProducts []models.ListedProduct
	ll := payload.ListedProducts

	totalPrice := 0.00

	err = repo.conn.Transaction(func(tx *gorm.DB) error {
		invoiceId, txErr := repo.Invoice.Insert(invoice, tx)
		if txErr != nil {
			return txErr
		}
		id = invoiceId

		for _, lProd := range ll {
			product, err := repo.Product.FindOneById(repository.FindOneBy{ID: lProd.ProductID, BusinessID: business.ID})
			if err != nil {
				return err
			}

			listedProduct := models.ListedProduct{
				PriceListed:    product.Price,
				QuantityListed: lProd.QuantityListed,
				ProductID:      product.ID,
				InvoiceID:      invoiceId,
			}
			listedProducts = append(listedProducts, listedProduct)
			totalPrice += product.Price * float64(lProd.QuantityListed)
		}
		ids, txErr := repo.ListedProduct.BatchInsert(listedProducts, tx)
		if txErr != nil {
			return txErr
		}
		log.Println(ids)

		var discount float64
		if payload.DiscountType == enums.EDiscountType.Percentage {
			discount = (payload.Discount / 100) * float64(totalPrice)
		} else {
			discount = float64(payload.Discount)
		}
		totalAmountToBePaid := totalPrice - discount - payload.Tax - payload.ServiceCharge

		repo.Invoice.Update(repository.FindOneBy{ID: invoiceId, BusinessID: business.ID}, models.Invoice{Price: totalPrice, Total: totalAmountToBePaid}, tx)

		return nil
	})
	if err != nil {
		return id, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return id, errData
}
