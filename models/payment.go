package models

import (
	"example/web-service-gin/db"
	"example/web-service-gin/forms"
	"net/http"
)

type Payment struct {
	ID           int64   `json:"id" gorm:"primary_key"`
	MerchantBuy  int64   `json:"merchant_buy"`
	MerchantSale int64   `json:"merchant_sale"`
	Date         int64   `json:"date" gorm:"serializer:unixtime;type:date"`
	DueDate      int64   `json:"due_date" gorm:"serializer:unixtime;type:date"`
	TotalAmount  float64 `json:"total_amount"`
	CreatedAt    int64   `json:"created_at" gorm:"autoUpdateTime"`
}

type PaymentModel struct{}

func (m PaymentModel) Create(form forms.CreatePaymentForm) (payment_id int64, err error) {

	payment := Payment{
		MerchantBuy:  form.MerchantBuy,
		MerchantSale: form.MerchantSale,
		Date:         form.Date.Unix(),
		DueDate:      form.DueDate.Unix(),
		TotalAmount:  form.TotalAmount,
	}
	result := db.GetDB().Create(&payment)
	return payment.ID, result.Error
}

func (m PaymentModel) All(merchant string, r *http.Request) (payments []Payment, err error) {
	result := db.GetDB().Scopes(Paginate(r)).Find(&payments)
	return payments, result.Error
}

func (m PaymentModel) Filter(merchant string, conditions forms.FilterPaymentForm, r *http.Request) (payments []Payment, err error) {
	c := make(map[string]interface{})
	if conditions.Date != "" {
		c["date"] = conditions.Date
	}
	if conditions.DueDate != "" {
		c["due_date"] = conditions.DueDate
	}
	if conditions.TotalAmount != "" {
		c["total_amount"] = conditions.TotalAmount
	}
	result := db.GetDB().Where(c).Scopes(Paginate(r)).Find(&payments)
	return payments, result.Error
}

func (m PaymentModel) One(merchant, payment_id int64) (payment Payment, err error) {
	result := db.GetDB().Find(&payment, "merchant_buy = ? or merchant_sale = ?", merchant, merchant, payment_id)
	return payment, result.Error
}

func (m PaymentModel) GetEmails() (merchants []Merchant, err error) {

	result := db.GetDB().Model(&Merchant{}).Select("distinct merchants.email").Joins(
		"inner join payments on merchants.id = payments.merchant_buy and payments.due_date > current_date()",
	).Scan(&merchants)

	return merchants, result.Error
}
