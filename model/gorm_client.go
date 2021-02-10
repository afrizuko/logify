package model

import (
	"log"

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

	db, err = gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
