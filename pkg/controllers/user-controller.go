package controllers

import (
	"encoding/json"
	"github.com/EmreSahna/go_mysql_book_management/pkg/models"
	"github.com/EmreSahna/go_mysql_book_management/pkg/utils"
	"net/http"
)

var NewUser models.User

func GetUserByUsernameAndPassword(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	utils.ParseBody(r, user)
	foundUser, _ := models.GetUserByUsernameAndPassword(user.Username, user.Password)
	res, _ := json.Marshal(foundUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
