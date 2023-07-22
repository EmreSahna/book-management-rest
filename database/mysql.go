package database

import (
	"github.com/EmreSahna/go_mysql_book_management/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var db *gorm.DB

func Connect() error {
	d, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		return err
	}

	db = d

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func Close() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
