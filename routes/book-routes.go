package routes

import (
	"github.com/EmreSahna/go_mysql_book_management/handlers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/book", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", handlers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", handlers.DeleteBook).Methods("DELETE")
}
