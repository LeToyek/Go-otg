package main

import (
	"go-otg/internal/app/http"
	"log"

	"github.com/joho/godotenv"
)

// main function
func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if 1+1 == 2 {
		http.NewApplication()
	}
}
