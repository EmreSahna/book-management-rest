package routes

import (
	"github.com/EmreSahna/go_mysql_book_management/handlers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/find", handlers.GetUserByUsernameAndPassword).Methods("POST")
}
