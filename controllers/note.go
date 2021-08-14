package controllers

import (
	"encoding/json"
	"net/http"
	"notekeeper/config"
	"notekeeper/models"
)

// functions for a user
// Get all notes
// Get a single note
// Create a new note
// Delete a note
// Update a note
var note = &models.Note{}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	data := db.Find(note)
	if data.RowsAffected == 0 {
		json.NewEncoder(w).Encode(note)
	}
	json.NewEncoder(w).Encode(data)
}

func GetSingleNote(w http.ResponseWriter, r *http.Request) {
	data := models.Note{Content: "Single note"}
	json.NewEncoder(w).Encode(data)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	data := models.Note{Content: "new note and first one"}
	result := db.Model(&models.Note{}).Create(&data)
	json.NewEncoder(w).Encode(result)
}
