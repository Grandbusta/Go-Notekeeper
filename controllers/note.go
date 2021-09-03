package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notekeeper/config"
	"notekeeper/models"
	"notekeeper/utils"
)

// functions for a user
// Get all notes
// Get a single note
// Create a new note
// Delete a note
// Update a note
func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	notes := []models.Note{}
	data := db.Find(&notes)
	res, _ := json.Marshal(data)
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{"status": http.StatusOK, "data": res})
	fmt.Println(res)
}

func GetSingleNote(w http.ResponseWriter, r *http.Request) {
	data := models.Note{Content: "Single note"}
	json.NewEncoder(w).Encode(data)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	data := models.Note{Content: "new note and first one", UserId: 1}
	result := db.Model(&models.Note{}).Create(&data)
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}
