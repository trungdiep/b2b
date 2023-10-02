package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type MerchantForm struct{}

type CreateMerchantForm struct {
	MerchantName string `form:"merchant_name" json:"merchant_name" binding:"required,max=130"`
	Email        string `binding:"required,email"`
}

func (f MerchantForm) MerchantName(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			message = "Please enter the merchant name"
		}
	case "max":
		message = "Merchant name should be less than 130"
	default:
		message = "Something went wrong, please try again later"
	}
	return message
}

func (f MerchantForm) Email(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			message = "Please enter the merchant email"
		}
	case "email":
		message = "Merchant email should be a email address"
	default:
		message = "Something went wrong, please try again later"
	}
	return message
}

// Create ...
func (f MerchantForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "MerchantName" {
				return f.MerchantName(err.Tag())
			}
			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Update ...
func (f MerchantForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "MerchantName" {
				return f.MerchantName(err.Tag())
			}
			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
