package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/go-faker/faker/v4"
)

func TestCreateBusiness(t *testing.T) {
	body := dtos.AddBusiness{}

    err := faker.FakeData(&body)
    if err != nil {
        t.Log(err)
    }
	body.IsCorporateAffair = true

    jsonBody, err := json.Marshal(body)
    if err != nil {
        t.Log("Error:", err)
        return
    }


	tokenString, err := helpers.CreateJWTToken("johndoe@exists.com")
	if (err != nil){
		t.Fatal("error creating test token")
	}
	req, _ := http.NewRequest("POST", "/business", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	rr := httptest.NewRecorder()

	handler := mdTest.Authorization(http.HandlerFunc(Repo.AddBusiness))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("CreateBusiness handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusCreated)
	}
}

func TestGetBusiness(t *testing.T) {
	tokenString, err := helpers.CreateJWTToken("johndoe@exists.com")
	if (err != nil){
		t.Fatal("error creating test token")
	}
	req, _ := http.NewRequest("GET", "/business", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	rr := httptest.NewRecorder()

	handler := mdTest.Authorization(http.HandlerFunc(Repo.GetBusiness))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("GetBusiness handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}
}