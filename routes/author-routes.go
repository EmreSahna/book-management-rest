package routes

import (
	"github.com/EmreSahna/go_mysql_book_management/handlers"
	"github.com/gorilla/mux"
)

var RegisterAuthorRoutes = func(router *mux.Router) {
	router.HandleFunc("/author", handlers.CreateAuthor).Methods("POST")
}
