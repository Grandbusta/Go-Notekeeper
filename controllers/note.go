package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"notekeeper/config"
	"notekeeper/middlewares"
	"notekeeper/models"
	"notekeeper/utils"

	"gorm.io/gorm"
)

// Get a single note
// Delete a note
// Update a note
type Note struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	uid, err := middlewares.ExtractId(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}
	user := models.User{}
	existErr := db.Model(&models.User{}).First(&user, uid).Error
	if errors.Is(existErr, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusNotFound, "Unauthorized")
		return
	}
	notes := []Note{}
	fErr := db.Model(&models.Note{}).Where("user_id=?", uid).Find(&notes).Error
	if fErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to get notes")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"userId": user.ID,
			"notes":  notes,
		},
	})
}

func GetSingleNote(w http.ResponseWriter, r *http.Request) {
	data := models.Note{Content: "Single note"}
	json.NewEncoder(w).Encode(data)
}

var n models.Note

func CreateNote(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}
	if len(n.Content) < 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	uid, err := middlewares.ExtractId(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
	}
	user := models.User{}
	existErr := db.Model(&models.User{}).First(&user, uid).Error
	if errors.Is(existErr, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusNotFound, "Unauthorized")
		return
	}
	data := models.Note{Content: n.Content, UserId: user.ID}
	cErr := db.Model(&models.Note{}).Create(&data).Error
	if cErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to create note")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"noteId": data.ID,
		},
	})
}
