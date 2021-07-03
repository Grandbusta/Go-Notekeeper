package main

import (
	"encoding/json"
	"log"
	"net/http"
	"notekeeper/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := models.Note{1, "new note and first one"}
	json.NewEncoder(w).Encode(data)
}

func main() {
	m := mux.NewRouter()

	m.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", m))
}
