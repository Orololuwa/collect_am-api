package handlers

import (
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/go-faker/faker/v4"
)

func TestCreateBusiness(t *testing.T) {
	// Case success
	body := dtos.AddBusiness{}

	err := faker.FakeData(&body)
	if err != nil {
		t.Log(err)
	}
	body.IsCorporateAffair = true

	_, errData := testHandlers.CreateBusiness(body)
	if errData != nil {
		t.Errorf("CreateBusiness handler returned an error, expected a successful call")
	}

	// test for failure on InsertBusiness repo function
	body.Email = "invalid"

	_, errData = testHandlers.CreateBusiness(body)
	if errData == nil {
		t.Errorf("CreateBusiness handler returned no error, expected an error on failed InsertBusiness call")
	}

	// test for failure on InsertKyc repo function
	body.BVN = "invalid"

	_, errData = testHandlers.CreateBusiness(body)
	if errData == nil {
		t.Errorf("CreateBusiness handler returned no error, expected an error on failed InsertKyc call")
	}
}

func TestGetBusiness(t *testing.T) {
	// Case success
	_, errData := testHandlers.GetBusiness(2)

	if errData != nil {
		t.Errorf("GetBusiness handler returned an error, expected a successful call")
	}

	//Case error (normal errors)
	_, errData = testHandlers.GetBusiness(0)

	if errData == nil {
		t.Errorf("GetBusiness handler returned no error, expected an error on failed GetOneById call")
	}

	//Case error (record not dound)
	business, errData := testHandlers.GetBusiness(1)

	if errData != nil && business != nil {
		t.Errorf("GetBusiness handler returned an error and a business data, expected nil, nil on (record not found) call on GetOneById")
	}
}

func TestUpdateBusiness(t *testing.T) {
	// Case Success
	body := map[string]interface{}{
		"name":                        "Manchester United Football Club",
		"description":                 "Biggest Football Club in England",
		"sector":                      "sports",
		"is_corporate_affair":         false,
		"logo":                        "http://logo.test",
		"certificate_of_registration": "http://random_url.test",
		"proof_of_address":            "http://random_url.test",
		"bvn":                         "000000456",
	}

	errData := testHandlers.UpdateBusiness(2, body)
	if errData != nil {
		t.Errorf("UpdateBusiness handler returned an error, expected a successful call")
	}

	// Case failed business validation
	errData = testHandlers.UpdateBusiness(0, body)
	if errData == nil {
		t.Errorf("UpdateBusiness handler returned no error, expected an error for failed validation on Business.GetOneById")
	}

	// Case failed insert or transaction manager
	body["name"] = "invalid"
	errData = testHandlers.UpdateBusiness(2, body)
	if errData == nil {
		t.Errorf("UpdateBusiness handler returned no error, expected an error for failed validation on Business.GetOneById")
	}
}
