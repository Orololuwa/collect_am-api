package helpers

import (
	"regexp"
	"unicode"

	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/enums"
	"github.com/go-playground/validator/v10"
)

func IsPasswordValid(password string) (bool, string) {
	// Check if the password is at least 8 characters long
	if len(password) < 8 {
		return false, "password length cannot be less than 8"
	}

	// Check if the password contains at least one uppercase letter
	hasUppercase := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
			break
		}
	}
	if !hasUppercase {
		return false, "password must contain at least one uppercase character"
	}

	// Check if the password contains at least one lowercase letter
	hasLowercase := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLowercase = true
			break
		}
	}
	if !hasLowercase {
		return false, "password must contain at least one lowercase character"
	}

	// Check if the password contains at least one digit
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return false, "password must contain atleast one digit"
	}

	// Check if the password contains at least one special character
	hasSpecialChar, _ := regexp.MatchString(`[!@#$%^&*()_+{}\[\]:;<>,.?/~]`, password)
	if !hasSpecialChar {
		return false, "password must contain at least one special character"
	}

	return true, ""
}

func DiscountValidator(fl validator.FieldLevel) bool {
	invoice := fl.Top().Interface().(dtos.CreateInvoice)

	if invoice.DiscountType == enums.EDiscountType.Percentage {
		return invoice.Discount >= 0.00 && invoice.Discount <= 99.00
	}
	return invoice.Discount >= 0.00
}
