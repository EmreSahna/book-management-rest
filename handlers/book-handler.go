package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/EmreSahna/go_mysql_book_management/database"
	"github.com/EmreSahna/go_mysql_book_management/models"
	"github.com/EmreSahna/go_mysql_book_management/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	dbForBook := database.GetDB()
	var Books []models.Book
	dbForBook.Find(&Books)

	res, _ := json.Marshal(Books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	dbForBook := database.GetDB()
	var getBook models.Book
	dbForBook.Where("ID=?", ID).Find(&getBook)

	res, _ := json.Marshal(getBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	err := utils.ParseBody(r, &b)
	if err != nil {
		fmt.Println("error while parsing")
	}

	dbForBook := database.GetDB()
	dbForBook.Create(&b)

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	dbForBook := database.GetDB()
	var book models.Book
	dbForBook.Where("ID=?", ID).Delete(book)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	dbForBook := database.GetDB()
	var bookDetails models.Book
	db := dbForBook.Where("ID=?", ID).Find(&bookDetails)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
