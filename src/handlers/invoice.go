package handlers

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/Orololuwa/collect_am-api/src/enums"
	"github.com/Orololuwa/collect_am-api/src/helpers/utils"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"github.com/Orololuwa/collect_am-api/src/types"
	"gorm.io/gorm"
)

func cleanEditInvoicePayload(body map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	fields := map[string]bool{
		"description":    true,
		"dueDate":        true,
		"tax":            true,
		"discount":       true,
		"discountType":   true,
		"serviceCharge":  true,
		"customerId":     true,
		"listedProducts": true,
	}

	for key, value := range body {
		if _, ok := fields[key]; ok {
			resKey := utils.CamelToSnakeCase(key)

			log.Printf("dataType: %+v\n", reflect.TypeOf(value))

			if resKey != "" {
				result[resKey] = value
			}
		}
	}

	return result
}

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
		Status:        enums.EInvoiceStatus.Draft,
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
			if payload.Discount >= totalPrice {
				return errors.New("discount must be less than totalPrice")
			}
			discount = float64(payload.Discount)
		}

		discountedTotal := totalPrice - discount

		// calculate tax and service charge on discounted total
		tax := (payload.Tax / 100) * float64(discountedTotal)
		serviceCharge := (payload.ServiceCharge / 100) * float64(discountedTotal)
		totalAmountToBePaid := discountedTotal + tax + serviceCharge

		repo.Invoice.Update(repository.FindOneBy{ID: invoiceId, BusinessID: business.ID}, models.Invoice{Price: totalPrice, Total: totalAmountToBePaid}, tx)

		return nil
	})
	if err != nil {
		return id, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return id, errData
}

func (repo *Repository) GetInvoice(payload types.GetAnInvoicePayload, options ...*Extras) (customer models.Invoice, errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	customer, err := repo.Invoice.FindOneById(repository.FindOneBy{ID: payload.Id, BusinessID: business.ID})
	if err != nil {
		return customer, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return customer, errData
}

func (repo *Repository) GetAllInvoices(query map[string]interface{}, options ...*Extras) (customers []models.Invoice, pagination repository.Pagination, errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}

	query["business_id"] = business.ID

	customers, pagination, err := repo.Invoice.FindAllWithPagination(query)
	if err != nil {
		return customers, pagination, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return customers, pagination, nil
}

var invoiceValidationMap = map[string]utils.FieldInfo{
	"description":    {reflect.String},
	"dueDate":        {reflect.String},
	"tax":            {reflect.Float64},
	"discount":       {reflect.Float64},
	"discountType":   {reflect.String},
	"serviceCharge":  {reflect.Float64},
	"customerId":     {reflect.Int},
	"listedProducts": {reflect.Slice},
}
var listedProductsValidationMap = map[string]utils.FieldInfo{
	"id":          {reflect.Int},
	"quantity":    {reflect.Int},
	"priceListed": {reflect.Float64},
}

func (repo *Repository) EditInvoice(payload types.EditInvoicePayload, options ...*Extras) (errData *ErrorData) {
	var business models.Business
	if len(options) > 0 && options[0] != nil {
		business = *options[0].Business
	}
	id := payload.ID

	body, err := utils.ValidateMap(payload.Body, invoiceValidationMap, true)
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	invoice, err := repo.Invoice.FindOneById(repository.FindOneBy{ID: id, BusinessID: business.ID})
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	savedListedProducts := invoice.ListedProducts
	savedListedProductsMap := make(map[uint]models.ListedProduct, 0)
	for _, savedProduct := range savedListedProducts {
		savedListedProductsMap[savedProduct.ID] = savedProduct
	} //a map of the listedproducts with the id as the key

	priceTotal := invoice.Price
	discount := invoice.Discount
	discountType := invoice.DiscountType
	tax := invoice.Tax
	serviceCharge := invoice.ServiceCharge
	listedProducts := make([]models.ListedProduct, 0) //this is what will be updated in batch update

	if body["customer_id"] != nil {
		customerId := uint(body["customer_id"].(int))
		_, err := repo.Customer.FindOneById(repository.FindOneBy{ID: customerId, BusinessID: business.ID})
		if err != nil {
			return &ErrorData{Error: err, Status: http.StatusBadRequest}
		}
	}
	if body["due_date"] != nil {
		dd := body["due_date"].(string)
		dueDate, err := time.Parse("2006-01-02", dd)
		if err != nil {
			return &ErrorData{Error: err, Status: http.StatusBadRequest}
		}
		body["due_date"] = dueDate
	}
	if body["discount"] != nil {
		discount = float64(body["discount"].(float64))
	}
	if body["discount_type"] != nil {
		discountType = body["discount_type"].(enums.IDiscountType)
	}
	if body["tax"] != nil {
		tax = float64(body["tax"].(float64))
	}
	if body["service_charge"] != nil {
		serviceCharge = float64(body["service_charge"].(float64))
	}
	if body["listed_products"] != nil {
		rawListedProducts := body["listed_products"].([]map[string]interface{})

		for _, productMap := range rawListedProducts {

			cleanedProductMap, err := utils.ValidateMap(productMap, listedProductsValidationMap, false)
			if err != nil {
				return &ErrorData{Error: err, Status: http.StatusBadRequest}
			}
			product := models.ListedProduct{
				ID:             uint(cleanedProductMap["id"].(int)),
				QuantityListed: uint(cleanedProductMap["quantity"].(int)),
				PriceListed:    float64(cleanedProductMap["priceListed"].(float64)),
			}
			if savedProduct, ok := savedListedProductsMap[uint(product.ID)]; ok {
				savedProduct.QuantityListed = product.QuantityListed
				savedProduct.PriceListed = product.PriceListed

				savedListedProductsMap[uint(product.ID)] = savedProduct
			}
			listedProducts = append(listedProducts, product)
		}

		newPriceTotal := 0.00
		for _, value := range savedListedProductsMap {
			newPriceTotal += (value.PriceListed * float64(value.QuantityListed))
		}
		priceTotal = newPriceTotal
	}

	var calculatedDiscount float64
	if discountType == enums.EDiscountType.Percentage {
		if discount >= 100 {
			return &ErrorData{Error: err, Status: http.StatusBadRequest, Message: "discount must be less than totalPrice"}
		}
		calculatedDiscount = (discount / 100) * priceTotal
	} else {
		if discount >= priceTotal {
			return &ErrorData{Error: err, Status: http.StatusBadRequest, Message: "discount must be less than totalPrice"}
		}
		calculatedDiscount = float64(discount)
	}

	discountedTotal := priceTotal - calculatedDiscount

	// calculate tax and service charge on discounted total
	calculatedTax := (tax / 100) * discountedTotal
	calculatedServiceCharge := (serviceCharge / 100) * discountedTotal
	totalAmountToBePaid := discountedTotal + calculatedTax + calculatedServiceCharge

	body["total"] = totalAmountToBePaid
	body["price"] = priceTotal

	delete(body, "listed_products")

	log.Printf("savedListedProducts: %+v\n\n", savedListedProducts)
	log.Printf("@listed_products_to_be_updated: %+v\n\n", listedProducts)
	log.Printf("calculatedTax: %+v\n\n", calculatedTax)
	log.Printf("calculatedServiceCharge: %+v\n\n", calculatedServiceCharge)
	log.Printf("priceTotal: %+v\n\n", priceTotal)
	log.Printf("discountedTotal: %+v\n\n", discountedTotal)
	log.Printf("totalAmountToBePaid: %+v\n\n", totalAmountToBePaid)
	log.Printf("body%+v\n\n", body)

	err = repo.conn.Transaction(func(tx *gorm.DB) error {
		err := repo.Invoice.UpdateWithMap(repository.FindOneBy{ID: invoice.ID, BusinessID: business.ID}, body, tx)
		if err != nil {
			return err
		}

		err = repo.ListedProduct.BatchUpdate(listedProducts, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return errData
}
