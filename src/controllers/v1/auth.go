package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/go-playground/validator"
)

func (m *v1) SignUp(w http.ResponseWriter, r *http.Request){
	var body dtos.UserSignUp
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}
	
	err = m.App.Validate.Struct(body)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Println(err)
		helpers.ClientError(w, err, http.StatusBadRequest, errors.Error())
		return
	}
	
	userId, errData := handlers.Repo.SignUpV2(body)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, userId, http.StatusCreated, "user account created successfully")
}

func (m *v1) LoginUser(w http.ResponseWriter, r *http.Request){
	var body dtos.UserLoginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	err = m.App.Validate.Struct(body)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Println(err)
		helpers.ClientError(w, err, http.StatusBadRequest, errors.Error())
		return
	}

	log.Println(body)
	if handlers.Repo == nil {
		log.Fatal("handlers.Repo is not initialized")
	}

	data, errData := handlers.Repo.LoginUserV2(body)
	if errData != nil {
		helpers.ClientError(w, errData.Error, errData.Status, errData.Message)
		return
	}

	helpers.ClientResponseWriter(w, data, http.StatusOK, "logged in successfully")
}