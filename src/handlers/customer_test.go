package handlers

import (
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/types"
	"github.com/go-faker/faker/v4"
)

func TestAddCustomer(t *testing.T) {
	var body dtos.CreateCustomer

	err := faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}

	_, errData := testHandlers.AddCustomer(body)
	if errData != nil {
		t.Errorf("AddCustomer handler returned an error, expected a successful call")
	}

	// test for failed InsertCustomer
	body.Email = "invalid"

	_, errData = testHandlers.AddCustomer(body)
	if errData == nil {
		t.Errorf("AddCustomer handler returned no error, expected an error for failed db operation on InsertCustomer")
	}

	// test for failed InsertAddress
	err = faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}
	body.UnitNumber = "invalid"

	_, errData = testHandlers.AddCustomer(body)
	if errData == nil {
		t.Errorf("AddCustomer handler returned no error, expected an error for failed db operation on InsertAddress")
	}
}

func TestUpdateCustomer(t *testing.T) {
	var body types.EditCustomerPayload

	err := faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}

	errData := testHandlers.EditCustomer(body)
	if errData != nil {
		t.Errorf("EditCustomer handler returned an error, expected a successful call")
	}

	// test for failed UpdateCustomer
	body.Id = 1

	errData = testHandlers.EditCustomer(body)
	if errData == nil {
		t.Errorf("EditCustomer handler returned no error, expected an error for failed db operation on UpdateCustomer")
	}
}
