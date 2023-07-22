package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadENV() error {
	// Load .env file
	err := godotenv.Load()
	fmt.Println("Loading .env file", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	return err
}
