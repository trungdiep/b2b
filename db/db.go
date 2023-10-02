package db

import (
	"fmt"
	"log"
	"os"

	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("NET"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_NAME"),
	)
	var err error
	db, err = ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Connected.")
	}
}

func ConnectDB(dsn string) (*gorm.DB, error) {

	var err error

	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
