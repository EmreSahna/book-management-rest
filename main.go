package main

import (
	"fmt"
	"github.com/EmreSahna/go_mysql_book_management/config"
	"github.com/EmreSahna/go_mysql_book_management/database"
	"log"
	"os"
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

	server := NewServer(fmt.Sprintf(":%s", os.Getenv("PORT")))
	server.Run()
}
