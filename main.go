package main

import (
	"fmt"
	"log"
	"net/http"
	"notekeeper/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func handleResponse() {
	m := mux.NewRouter()

	m.HandleFunc("/notes", controllers.GetAllNotes).Methods("GET")
	m.HandleFunc("/notes/{id}", controllers.GetSingleNote).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", m))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println(".env loaded")
	}
	handleResponse()
}
