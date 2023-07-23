package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/EmreSahna/go_mysql_book_management/database"
	"github.com/EmreSahna/go_mysql_book_management/models"
	"github.com/EmreSahna/go_mysql_book_management/utils"
	"net/http"
)

type FindUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserByUsernameAndPassword(w http.ResponseWriter, r *http.Request) {
	var request FindUserRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		fmt.Println("error while parsing")
	}

	dbForUser := database.GetDB()
	var user models.User
	dbForUser.Where("username=?", request.Username).Where("password=?", request.Password).Find(&user)

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := utils.ParseBody(r, &u)
	if err != nil {
		return
	}

	dbForUser := database.GetDB()
	dbForUser.Create(&u)

	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
