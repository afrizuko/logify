package model

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func GetConnection() *gorm.DB {

	if db != nil {
		return db
	}

	db, err = gorm.Open(mysql.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDSN() string {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "admin:S3cr3t123*@tcp(192.168.1.105:3306)/shiva?charset=utf8mb4&parseTime=True&loc=Local"
	}
	return dsn
}
