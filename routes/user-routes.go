package routes

import (
	"github.com/EmreSahna/go_mysql_book_management/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/find", controllers.GetUserByUsernameAndPassword).Methods("POST")
}
