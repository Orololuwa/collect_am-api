package handlers

import (
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/go-faker/faker/v4"
)

func TestAddCustomer(t *testing.T) {
	var body dtos.CreateCustomer

	err := faker.FakeData(&body)
	if err != nil {
		t.Error(err)
	}

	t.Log(body)

	_, errData := testHandlers.AddCustomer(body)
	if errData != nil {
		t.Errorf("AddCustomer handler returned an error, expected a successful call")
	}
}
