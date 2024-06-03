// main.go
package main

import (
	"log"
	"net/http"
	"github.com/RomanBagger/myproject.git/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	log.Println("Starting server on :10000")
	if err := http.ListenAndServe(":10000", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
