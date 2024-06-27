package handlers

import (
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/go-faker/faker/v4"
)

func TestCreateProduct(t *testing.T) {
	body := dtos.AddProduct{}

	// case success
	err := faker.FakeData(&body)
    if err != nil {
        t.Log(err)
    }

	_, errData :=  testHandlers.CreateProduct(body)
	if errData != nil {
		t.Errorf("AddProduct handler returned an error, expected a successful call")
	}

	// case: failed InsertProduct operation
	body.Code = "invalid"

	_, errData =  testHandlers.CreateProduct(body)
	if errData == nil {
		t.Errorf("AddProduct handler returned no error, expected an error for failed db operation on InsertProduct")
	}
}