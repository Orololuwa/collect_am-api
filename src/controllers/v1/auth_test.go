package v1

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
	body.Email = "johndoe@null.com"
	body.Phone = "+2340000000002"

    jsonBody, err := json.Marshal(body)
    if err != nil {
        t.Log("Error:", err)
        return
    }

	// Test for success
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(v1.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("SignUp handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusCreated)
	}

	// Test for missing body
	req, _ = http.NewRequest("POST", "/auth/signup", bytes.NewBuffer([]byte(``)))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(v1.SignUp)
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

	handler = http.HandlerFunc(v1.SignUp)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("SignUp handler returned wrong response code for failed validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

}

func TestLoginHandler(t *testing.T){
	body := dtos.UserLoginBody{}
	err := faker.FakeData(&body)
    if err != nil {
        t.Log(err)
    }
	body.Password = fmt.Sprintf("%s123#", body.Password)
	// jsonBody, err := json.Marshal(body)
    // if err != nil {
    //     t.Log("Error:", err)
    //     return
    // }

	// Test for missing request body
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer([]byte(``)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(v1.LoginUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Login handler returned wrong response code for missing request body: got %d, wanted %d", rr.Code, http.StatusInternalServerError)
	}

	// test validator with an invalid email
	body.Email = "invalid"
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Log("Error:", err)
		return
	}
	
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(v1.LoginUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Login handler returned wrong response code for failed validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// Test for failed db operation
	body.Email = "johndoe@fail.com"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(v1.LoginUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Login handler returned wrong response code for failed db operation on userExists validation: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// Test for error if user doesn't exist
	body.Email = "johndoe@null.com"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }
	
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(v1.LoginUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Login handler returned wrong response code for error if user doesn't exist: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// Test for failed password hash auth
	body.Email = "test_fail@test.com"
	jsonBody, err = json.Marshal(body)
	if err != nil {
		t.Log("Error:", err)
		return
	}
	
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(v1.LoginUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Login handler returned wrong response code for failed password hash auth: got %d, wanted %d", rr.Code, http.StatusBadRequest)
	}

	// Test for success	
	body.Password = "Testpass123###"
	body.Email = "test_correct@test.com"
	jsonBody, err = json.Marshal(body)
	if err != nil {
        t.Log("Error:", err)
        return
    }

	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(v1.LoginUser)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Login handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}
}