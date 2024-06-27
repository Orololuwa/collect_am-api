package handlers

import (
	"errors"
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/types"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) SignUp(payload dtos.UserSignUp) (userId uint, errData *ErrorData) {
	// ctx := context.Background()
	emailExists, err := m.User.GetOneByEmail(payload.Email)
	if err != nil && err.Error() != "record not found" {
		return userId, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}
	if emailExists.ID != 0 {
		return userId, &ErrorData{Message: "email exists", Error: errors.New(""), Status: http.StatusBadRequest}
	}

	phoneExists, err := m.User.GetOneByPhone(payload.Phone)
	if err != nil && err.Error() != "record not found" {
		return userId, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}
	if phoneExists.ID != 0 {
		return userId, &ErrorData{Message: "phone exists", Error: errors.New(""), Status: http.StatusBadRequest}
	}

	// validate password
	isPasswordValid, validationMessage := helpers.IsPasswordValid(payload.Password)
	if !isPasswordValid {
		return userId, &ErrorData{Message: validationMessage, Error: errors.New(""), Status: http.StatusBadRequest}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return userId, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	userId, err = m.User.InsertUser(models.User{FirstName: payload.FirstName, LastName: payload.LastName, Email: payload.Email, Phone: payload.Phone, Password: string(hashedPassword)})
	if err != nil {
		return userId, &ErrorData{Error: err, Status: http.StatusBadRequest}
	}

	return userId, nil
}

func (m *Repository) LoginUser(payload dtos.UserLoginBody) (data types.LoginSuccessResponse, errData *ErrorData) {
	user, err := m.User.GetOneByEmail(payload.Email)
	if err != nil {
		if err.Error() == "record not found" {
			return data, &ErrorData{Message: "invalid email or password", Error: err, Status: http.StatusBadRequest}
		} else {
			return data, &ErrorData{Error: err, Status: http.StatusBadRequest}
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return data, &ErrorData{Message: "invalid email or password", Error: err, Status: http.StatusBadRequest}
	}

	tokenString, err := helpers.CreateJWTToken(payload.Email)

	if err != nil {
		return data, &ErrorData{Error: err, Status: http.StatusInternalServerError}
	}

	data = types.LoginSuccessResponse{Email: payload.Email, Token: tokenString}
	return data, nil
}
