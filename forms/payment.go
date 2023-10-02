package forms

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type PaymentForm struct{}

type CustomerDate time.Time

func (j *CustomerDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("02-01-2006", s)
	if err != nil {
		return err
	}
	*j = CustomerDate(t)
	return nil
}

func (j CustomerDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j *CustomerDate) UnmarshalText(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("02-01-2006", s)
	if err != nil {
		return err
	}
	*j = CustomerDate(t)
	return nil
}

// Maybe a Format function for printing your date
func (j CustomerDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j CustomerDate) Unix() int64 {
	t := time.Time(j)
	return t.Unix()
}

type CreatePaymentForm struct {
	MerchantBuy  int64        `form:"merchant_buy" json:"merchant_buy" binding:"required"`
	MerchantSale int64        `form:"merchant_sale" json:"merchant_sale" binding:"required"`
	Date         CustomerDate `form:"date" json:"date" binding:"required"`
	DueDate      CustomerDate `form:"due_date" json:"due_date" binding:"required"`
	TotalAmount  float64      `form:"total_amount" json:"total_amount" binding:"required"`
}

type FilterPaymentForm struct {
	Date        string `form:"date"`
	DueDate     string `form:"due_date"`
	TotalAmount string `form:"total_amount"`
}

func (f PaymentForm) MerchantBuy(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			message = "Please enter the merchant buy id"
		}
	default:
		message = "Something went wrong, please try again later"
	}

	return message
}

func (f PaymentForm) MerchantSale(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			message = "Please enter the merchant sale id"
		}
	default:
		message = "Something went wrong, please try again later"
	}

	return message
}

func (f PaymentForm) Date(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			message = "Please enter the invoice date"
		}
	case "datetime":
		message = "Date's format is wrong (%d-%m-%Y)"
	default:
		message = "Something went wrong, please try again later"
	}

	return message
}

func (f PaymentForm) DueDate(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			message = "Please enter the invoice due date"
		}
	case "datetime":
		message = "Due date's format is wrong (%d-%m-%Y)"
	default:
		message = "Something went wrong, please try again later"
	}

	return message
}

func (f PaymentForm) TotalAmount(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			message = "Please enter the invoice due date"
		}
	default:
		message = "Something went wrong, please try again later"
	}

	return message
}

// Create ...
func (f PaymentForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "MerchantBuy" {
				return f.MerchantBuy(err.Tag())
			}
			if err.Field() == "MerchantSale" {
				return f.MerchantSale(err.Tag())
			}
			if err.Field() == "Date" {
				return f.Date(err.Tag())
			}
			if err.Field() == "DueDate" {
				return f.DueDate(err.Tag())
			}
			if err.Field() == "TotalAmount" {
				return f.TotalAmount(err.Tag())
			}
		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
