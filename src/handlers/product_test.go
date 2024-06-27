package handlers

import (
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"github.com/go-faker/faker/v4"
)

func TestAddProduct(t *testing.T) {
	body := dtos.AddProduct{}

	// case success
	err := faker.FakeData(&body)
	if err != nil {
		t.Log(err)
	}

	_, errData := testHandlers.AddProduct(body)
	if errData != nil {
		t.Errorf("AddProduct handler returned an error, expected a successful call")
	}

	// case: failed InsertProduct operation
	body.Code = "invalid"

	_, errData = testHandlers.AddProduct(body)
	if errData == nil {
		t.Errorf("AddProduct handler returned no error, expected an error for failed db operation on InsertProduct")
	}
}

func TestUpdateProduct(t *testing.T) {
	body := dtos.UpdateProduct{}

	// case success
	err := faker.FakeData(&body)
	if err != nil {
		t.Log(err)
	}

	errData := testHandlers.UpdateProduct(body)
	if errData != nil {
		t.Errorf("UpdateProduct handler returned an error, expected a successful call")
	}

	// case: failed InsertProduct operation
	body.Category = "invalid"

	errData = testHandlers.UpdateProduct(body)
	if errData == nil {
		t.Errorf("UpdateProduct handler returned no error, expected an error for failed db operation on InsertProduct")
	}
}

func TestGetAllProducts(t *testing.T) {
	query := repository.FilterQueryPagination{}

	// case success
	_, _, errData := testHandlers.GetAllProducts(query)
	if errData != nil {
		t.Errorf("GetAllProducts handler returned an error, expected a successful call")
	}

	// case: failed Find operation
	query.Page = 1

	_, _, errData = testHandlers.GetAllProducts(query)
	if errData == nil {
		t.Errorf("GetAllProducts handler returned no error, expected an error for failed db operation on InsertProduct")
	}
}
