package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/go-faker/faker/v4"
)

func TestSignUp(t *testing.T){
	body := dtos.UserSignUp{}

    err := faker.FakeData(&body)
    if err != nil {
        t.Log(err)
    }
	body.Password = fmt.Sprintf("%s123#", body.Password)

    jsonBody, err := json.Marshal(body)
    if err != nil {
        t.Log("Error:", err)
        return
    }

	fmt.Println(body)

	// Test for success
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("SignUp handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusCreated)
	}

	// Test for missing body
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer([]byte(``)))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("SignUp handler returned wrong response code for missing request body: got %d, wanted %d", rr.Code, http.StatusInternalServerError)
	}

	// test validator with an invalid email
	body.Email = "invalid"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for failed validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// Test for emailExists and phoneExists validation
	// 
	body.Password = "Testpass123#"
	body.Email = "johndoe@fail.com"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for failed db operation on isEmailExist validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// 
	body.Email = "johndoe@null.com"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for success on isEmailExist validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// 
	body.Email = faker.Email()
	body.Phone = "+2340000000000"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for failed db operation on isPhoneExist validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// 
	body.Phone = "+2340000000001"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for success on isPhoneExist validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// Test for invalid password
	body.Phone = faker.E164PhoneNumber()
	body.Password = "invalid"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for invalid password: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// Test for failed db operation on createUser
	body.Password = "Testpass123#"
	body.FirstName = "fail"
	jsonBody, err = json.Marshal(body)
	if err != nil {
		t.Log("Error:", err)
		return
	}
	
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for failed db operation on createUser: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}
}