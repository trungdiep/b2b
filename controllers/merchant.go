package controllers

import (
	"example/web-service-gin/forms"
	"example/web-service-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var merchantModels = new(models.MerchantModel)
var merchantForm = new(forms.MerchantForm)

type MerchantController struct{}

func (ctrl MerchantController) Create(c *gin.Context) {
	var form forms.CreateMerchantForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := merchantForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}
	id, err := merchantModels.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Merchant could not be created."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Merchant created", "id": id})
}

func getMerchantID(c *gin.Context) string {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.Param("id")
}

func (ctrl MerchantController) All(c *gin.Context) {

	datas, err := merchantModels.All(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": datas})
}

func (ctrl MerchantController) One(c *gin.Context) {
	id := c.Param("id")
	getId, err := strconv.ParseInt(id, 10, 64)

	if getId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	data, err := merchantModels.One(getId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Merchant not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})

}

func (ctrl MerchantController) Update(c *gin.Context) {
	id := c.Param("id")
	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}
	var form forms.CreateMerchantForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := merchantForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = merchantModels.Update(getID, form)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotAcceptable,
			gin.H{"Message": "Merchant could not be updated"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant updated"})
}

func (ctrl MerchantController) Delete(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"Message": "Invalid parameter"})
		return
	}

	err = merchantModels.Delete(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable,
			gin.H{"Message": "Merchant could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant deleted"})

}
