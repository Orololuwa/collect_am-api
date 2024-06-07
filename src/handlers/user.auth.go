package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/types"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request){
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
	
	ctx := context.Background()
	emailExists, err := m.User.GetAUser(ctx, nil, models.User{Email: body.Email})
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	if emailExists != nil {
		helpers.ClientError(w, errors.New("email exists"), http.StatusBadRequest, "")
		return
	}

	phoneExists, err := m.User.GetAUser(ctx, nil, models.User{Phone: body.Phone})
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	if phoneExists != nil {
		helpers.ClientError(w, errors.New("phone exists"), http.StatusBadRequest, "")
		return
	}
	
	// validate password
	isPasswordValid, validationMessage := helpers.IsPasswordValid(body.Password)
	if !isPasswordValid {
		helpers.ClientError(w, errors.New(validationMessage), http.StatusBadRequest, "")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	userId, err := m.User.CreateAUser(ctx, nil, models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Phone: body.Phone, Password: string(hashedPassword)})
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}

	helpers.ClientResponseWriter(w, userId, http.StatusCreated, "user account created successfully")
}

func (m *Repository) LoginUser(w http.ResponseWriter, r *http.Request){
	var body dtos.UserLoginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
		return
	}

	log.Println(body)

	err = m.App.Validate.Struct(body)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		log.Println(err)
		helpers.ClientError(w, err, http.StatusBadRequest, errors.Error())
		return
	}

	ctx := context.Background()
	user, err := m.User.GetAUser(ctx, nil, models.User{Email: body.Email})
	if err != nil {
		helpers.ClientError(w, err, http.StatusBadRequest, "")
		return
	}
	if user == nil {
		helpers.ClientError(w, errors.New("invalid email"), http.StatusBadRequest, "")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		helpers.ClientError(w, errors.New("invalid password"), http.StatusBadRequest, "")
		return
	}

	tokenString, err := helpers.CreateJWTToken(body.Email)

	if err != nil {
		helpers.ClientError(w, err, http.StatusInternalServerError, "")
	}

	data := types.LoginSuccessResponse{Email: body.Email, Token: tokenString}

	helpers.ClientResponseWriter(w, data, http.StatusOK, "logged in successfully")
}