package handlers

import (
	"fmt"
	"testing"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/go-faker/faker/v4"
)

func TestSignUp(t *testing.T) {
	body := dtos.UserSignUp{}

	err := faker.FakeData(&body)
	if err != nil {
		t.Log(err)
	}
	body.Password = fmt.Sprintf("%s123#", body.Password)
	body.Email = "johndoe@null.com"
	body.Phone = "+2340000000002"

	// Test for success
	_, errData := testHandlers.SignUp(body)

	if errData != nil {
		t.Errorf("SignUp handler returned an error, expected a successful function execution")
	}

	// Test for emailExists and phoneExists validation
	// emailExists
	body.Password = "Testpass123#"
	body.Email = "johndoe@exists.com"

	_, errData = testHandlers.SignUp(body)

	if errData == nil {
		t.Errorf("SignUp handler returned no error, expected an error for emailExists case")
	}

	//
	body.Email = faker.Email()
	body.Phone = "+2340000000001"

	_, errData = testHandlers.SignUp(body)

	if errData == nil {
		t.Errorf("SignUp handler returned no error, expected an error for phoneExists case")
	}

	// Test for invalid password
	body.Phone = faker.E164PhoneNumber()
	body.Password = "invalid"

	_, errData = testHandlers.SignUp(body)

	if errData == nil {
		t.Errorf("SignUp handler returned no error, expected an error for invalid password case")
	}

	// Test for failed db operation on createUser
	body.Password = "Testpass123#"
	body.FirstName = "fail"

	_, errData = testHandlers.SignUp(body)

	if errData == nil {
		t.Errorf("SignUp handler returned no error, expected an error for failed DB operation case")
	}
}

func TestLoginHandler(t *testing.T) {
	body := dtos.UserLoginBody{}

	// Test for success
	body.Password = "Testpass123###"
	body.Email = "test_correct@test.com"

	_, errData := testHandlers.LoginUser(body)

	if errData != nil {
		t.Errorf("Login handler returned an error, expected a successful function execution")
	}

	// Test for error if user doesn't exist
	body.Email = "johndoe@null.com"

	_, errData = testHandlers.LoginUser(body)

	if errData == nil {
		t.Errorf("Login handler returned no error, expected an error for userExists validation case")
	}

	// Test for failed password hash auth
	body.Email = "hash_fail@test.com"
	_, errData = testHandlers.LoginUser(body)

	if errData == nil {
		t.Errorf("Login handler returned no error, expected an error for password hash case")
	}
}
