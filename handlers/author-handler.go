package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/EmreSahna/go_mysql_book_management/database"
	"github.com/EmreSahna/go_mysql_book_management/models"
	"github.com/EmreSahna/go_mysql_book_management/utils"
	"net/http"
)

type AuthorCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthorResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var request AuthorCreateRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		fmt.Println("error while parsing")
	}

	var a = models.Author{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	dbForBook := database.GetDB()
	dbForBook.Create(&a)

	res, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
