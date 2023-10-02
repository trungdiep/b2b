package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/db"
	"example/web-service-gin/models"
	"example/web-service-gin/tasks"
	"fmt"
	"log"
	"os"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file.")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	// r.Use(CORSMiddleware())
	// r.Use(RequestIDMiddleware())
	// r.Use(gzip.Gzip(gzip.DefaultCompression))
	db.InitDb()

	var dbm = db.GetDB()

	dbm.AutoMigrate(models.Merchant{}, models.Payment{})

	v1 := r.Group("/v1")
	{
		merchant := new(controllers.MerchantController)
		payment := new(controllers.PaymentController)

		v1.GET("/merchant/:id/payments", payment.All)
		v1.POST("/merchant", merchant.Create)
		v1.GET("/merchants", merchant.All)
		v1.GET("/merchant/:id", merchant.One)
		v1.PUT("/merchant/:id", merchant.Update)
		v1.DELETE("/merchant/:id", merchant.Delete)
		v1.POST("/payment", payment.Create)
	}
	// port := os.Getenv("PORT")
	jobrunner.Start()
	jobrunner.Schedule("@every 300s", tasks.SendMail{})

	r.Run()
}
