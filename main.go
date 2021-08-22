package main

import (
	"fmt"
	"log"
	"net/http"
	"notekeeper/config"
	"notekeeper/controllers"
	"notekeeper/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	//Load env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println(".env loaded")
	}
	//Connect to database
	config.DbConnect()
}

func handleResponse() {
	m := mux.NewRouter()
	getSingleNote := controllers.GetSingleNote
	getAllNotes := controllers.GetAllNotes
	m.HandleFunc("/notes", getAllNotes).Methods("GET")
	m.HandleFunc("/notes/{id}", getSingleNote).Methods("GET")
	m.HandleFunc("/new-note", controllers.CreateNote).Methods("POST")
	m.Use(middlewares.AuthUser)
	log.Fatal(http.ListenAndServe(":8081", m))
}

func main() {
	handleResponse()
}
