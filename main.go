package main

import (
	"github.com/EmreSahna/go_mysql_book_management/config"
	"github.com/EmreSahna/go_mysql_book_management/database"
	"github.com/EmreSahna/go_mysql_book_management/middleware"
	routes "github.com/EmreSahna/go_mysql_book_management/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v\n", err)
	}

	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatalf("Could not close database: %v\n", err)
		}
	}()

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	routes.RegisterBookRoutes(r)
	routes.RegisterAuthorRoutes(r)
	routes.RegisterUserRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
