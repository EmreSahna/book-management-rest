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

type BookCreateRequest struct {
	Title       string  `json:"title"`
	AuthorId    int     `json:"author_id"`
	ISBN        string  `json:"isbn"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type BookResponse struct {
	Title       string         `json:"title"`
	Author      AuthorResponse `json:"author"`
	ISBN        string         `json:"isbn"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Quantity    int            `json:"quantity"`
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	authorId := r.URL.Query().Get("authorId")
	dbForBook := database.GetDB()
	if authorId != "" {
		dbForBook = dbForBook.Where("author_id=?", authorId)
	}

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

	var getAuthor models.Author
	dbForBook.Where("ID=?", getBook.AuthorId).Find(&getAuthor)

	var bookResponse = BookResponse{
		Title: getBook.Title,
		Author: AuthorResponse{
			FirstName: getAuthor.FirstName,
			LastName:  getAuthor.LastName,
		},
		ISBN:        getBook.ISBN,
		Description: getBook.Description,
		Price:       getBook.Price,
		Quantity:    getBook.Quantity,
	}

	res, _ := json.Marshal(bookResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var request BookCreateRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		fmt.Println("error while parsing")
	}

	dbForBook := database.GetDB()
	var a models.Author
	dbForBook.First(&a, request.AuthorId)

	var b = models.Book{
		Title:       request.Title,
		AuthorId:    a.ID,
		ISBN:        request.ISBN,
		Description: request.Description,
		Price:       request.Price,
		Quantity:    request.Quantity,
	}

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

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
