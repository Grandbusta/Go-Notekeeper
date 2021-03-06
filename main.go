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
	m.HandleFunc("/notes", controllers.GetAllNotes).Methods("GET")
	m.HandleFunc("/notes/{noteId}", controllers.GetSingleNote).Methods("GET")
	m.HandleFunc("/new-note", controllers.CreateNote).Methods("POST")
	m.HandleFunc("/update-note", controllers.UpdateNote).Methods("PUT")
	m.HandleFunc("/delete-note", controllers.DeleteNote).Methods("DELETE")
	m.HandleFunc("/users/signup", controllers.CreateUser).Methods("POST")
	m.HandleFunc("/users/login", controllers.LoginUser).Methods("POST")
	m.Use(middlewares.Cors)
	exclude := []string{"/users/login", "/users/signup"}
	middlewares.Auth(m, exclude)
	log.Fatal(http.ListenAndServe(":8081", m))
}

func main() {
	handleResponse()
}
