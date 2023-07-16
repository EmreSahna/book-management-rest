package models

import (
	"github.com/EmreSahna/go_mysql_book_management/config"
	"github.com/jinzhu/gorm"
)

var dbForBook *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	dbForBook = config.GetDB()
	dbForBook.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	dbForBook.NewRecord(b)
	dbForBook.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	dbForBook.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := dbForBook.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	dbForBook.Where("ID=?", ID).Delete(book)
	return book
}
