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

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	data := models.Note{Content: "new note and first one"}
	result := db.Model(&models.Note{}).Create(&data)
	json.NewEncoder(w).Encode(result)
}

func GetSingleNote(w http.ResponseWriter, r *http.Request) {
	data := models.Note{Content: "Single note"}
	json.NewEncoder(w).Encode(data)
}
