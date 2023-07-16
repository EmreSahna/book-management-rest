package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:!Mysql123@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}
	log.Println("Connected to database")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
