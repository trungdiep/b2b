package models

import (
	"errors"
	"example/web-service-gin/db"
	"example/web-service-gin/forms"
	"net/http"
)

type Merchant struct {
	ID           int64  `json:"id" gorm:"primary_key"`
	MerchantName string `json:"merchant_name"`
	Email        string `json:"email"`
	UpdatedAt    int64  `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedAt    int64  `json:"created_at" gorm:"autoUpdateTime"`
}

type MerchantModel struct{}

func (m MerchantModel) Create(form forms.CreateMerchantForm) (int64, error) {
	merchant := &Merchant{
		MerchantName: form.MerchantName,
		Email:        form.Email,
	}
	result := db.GetDB().Create(merchant)
	return merchant.ID, result.Error
}

func (m MerchantModel) All(r *http.Request) (merchants []Merchant, err error) {
	result := db.GetDB().Scopes(Paginate(r)).Find(&merchants)
	return merchants, result.Error
}

func (m MerchantModel) One(id int64) (merchant Merchant, err error) {
	result := db.GetDB().Find(&merchant, id)

	return merchant, result.Error
}

func (m MerchantModel) Update(id int64, form forms.CreateMerchantForm) error {

	result := db.GetDB().Model(&Merchant{}).Where("id = ?", id).Updates(Merchant{MerchantName: form.MerchantName, Email: form.Email})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Update 0 records.")
	}
	return result.Error
}

func (m MerchantModel) Delete(id int64) error {
	result := db.GetDB().Delete(&Merchant{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("No records were deleted.")
	}

	return result.Error
}
