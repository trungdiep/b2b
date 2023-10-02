package controllers

import (
	"example/web-service-gin/forms"
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var paymentModels = new(models.PaymentModel)
var paymentForm = new(forms.PaymentForm)
var filterPaymentForm = new(forms.FilterPaymentForm)

type PaymentController struct{}

func (ctrl PaymentController) Create(c *gin.Context) {
	var form forms.CreatePaymentForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := paymentForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}
	id, err := paymentModels.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Payment could not be created."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment created", "id": id})
}

func (ctrl PaymentController) All(c *gin.Context) {
	var form forms.FilterPaymentForm
	var datas []models.Payment
	var err error
	merchantId := getMerchantID(c)

	if validationErr := c.ShouldBindQuery(&form); validationErr == nil {
		datas, err = paymentModels.Filter(merchantId, form, c.Request)
	} else {
		datas, err = paymentModels.All(merchantId, c.Request)
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": datas})
}
