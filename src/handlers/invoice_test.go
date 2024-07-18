package handlers

import (
	"log"
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/types"
	"github.com/go-faker/faker/v4"
)

func TestCreateInvoice(t *testing.T) {
	var body types.CreateInvoicePayload

	err := faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}

	_, errData := testHandlers.CreateInvoice(body)
	if errData != nil {
		t.Errorf("CreateInvoice handler returned an error, expected a successful call")
	}

	// test for invalid due date
	body.DueDate = "invalid"
	_, errData = testHandlers.CreateInvoice(body)
	if errData == nil {
		t.Errorf("CreateInvoice handler returned no error, expected an error for an invalid due date")
	}

	// test for existing invoice no
	err = faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}

	body.Code = "exists"
	_, errData = testHandlers.CreateInvoice(body)
	if errData == nil {
		t.Errorf("CreateInvoice handler returned no error, expected an error for existing invoice no")
	}

	// test for a failed Invoice.Insert db operation
	err = faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}

	body.Code = "invalid"
	_, errData = testHandlers.CreateInvoice(body)
	if errData == nil {
		t.Errorf("CreateInvoice handler returned no error, expected an error for a failed Invoice.Insert db operation")
	}

	// test for a failed ListedProduct.BatchInsert db operation
	err = faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}

	log.Println(make([]dtos.CreateListedProduct, 0))
	body.ListedProducts = make([]dtos.CreateListedProduct, 0)
	_, errData = testHandlers.CreateInvoice(body)
	if errData == nil {
		t.Errorf("CreateInvoice handler returned no error, expected an error for a failed ListedProduct.BatchInsert db operation")
	}
}

func TestGetAInvoice(t *testing.T) {
	var findBy types.GetAnInvoicePayload

	_, errData := testHandlers.GetInvoice(findBy)
	if errData != nil {
		t.Errorf("GetInvoice handler returned an error, expected a successful call")
	}

	// test for failed UpdateInvoice
	findBy.Id = 1

	_, errData = testHandlers.GetInvoice(findBy)
	if errData == nil {
		t.Errorf("GetInvoice handler returned no error, expected an error for failed db operation on FindOneById")
	}
}

func TestGetAllInvoices(t *testing.T) {
	query := make(map[string]interface{}, 0)

	// case success
	_, _, errData := testHandlers.GetAllInvoices(query)
	if errData != nil {
		t.Errorf("GetAllInvoices handler returned an error, expected a successful call")
	}

	// case: failed Find operation
	query["page"] = 1

	_, _, errData = testHandlers.GetAllInvoices(query)
	if errData == nil {
		t.Errorf("GetAllInvoices handler returned no error, expected an error for failed db operation on FindAllWithPagination")
	}
}
