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
