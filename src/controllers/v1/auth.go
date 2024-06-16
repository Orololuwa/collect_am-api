package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
)

func (m *V1) SignUp(w http.ResponseWriter, r *http.Request){
	var body dtos.UserSignUp
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}
	
	err = m.App.Validate.Struct(body)
	log.Println(body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	
	userId, errData := m.Handlers.SignUp(body)
	if errData != nil {
		log.Println(errData)
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, userId, http.StatusCreated, "user account created successfully")
}

func (m *V1) LoginUser(w http.ResponseWriter, r *http.Request){
	var body dtos.UserLoginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	err = m.App.Validate.Struct(body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	data, errData := m.Handlers.LoginUser(body)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, data, http.StatusOK, "logged in successfully")
}