package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"notekeeper/config"
	"notekeeper/middlewares"
	"notekeeper/models"
	"notekeeper/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

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
		return
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
	db := config.DB
	params := mux.Vars(r)
	noteId := params["noteId"]
	uid, err := middlewares.ExtractId(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	user := models.User{}
	existErr := db.Model(&models.User{}).First(&user, uid).Error
	if errors.Is(existErr, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusNotFound, "Unauthorized")
		return
	}
	note := Note{}
	fErr := db.Model(&models.Note{}).Where("user_id=? AND id=?", uid, noteId).Find(&note).Error
	if fErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occured")
		return
	}
	if note.Content == "" {
		utils.RespondWithJson(w, http.StatusNotFound, map[string]interface{}{
			"status": http.StatusNotFound,
			"data":   nil,
		})
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"userId": uid,
			"note":   note,
		}})
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	n := models.Note{}
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}
	if len(n.Content) <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	uid, err := middlewares.ExtractId(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
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

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	note := Note{}
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}
	if len(note.Content) <= 0 && note.ID <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	uid, eErr := middlewares.ExtractId(r)
	if eErr != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	user := models.User{}
	existErr := db.Model(&models.User{}).First(&user, uid).Error
	if errors.Is(existErr, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusNotFound, "Unauthorized")
		return
	}
	var n models.Note
	fErr := db.Model(&models.Note{}).First(&n, note.ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusNotFound, "note not found")
		return
	}
	uErr := db.Model(&models.Note{}).Where("user_id=? AND id=?", uid, n.ID).Update("content", note.Content).Error
	if uErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to update note")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"noteId": n.ID,
		},
	})
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	note := Note{}
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}
	if note.ID <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	uid, eErr := middlewares.ExtractId(r)
	if eErr != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var n models.Note
	fErr := db.Model(&models.Note{}).First(&n, note.ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusNotFound, "note not found")
		return
	}
	dErr := db.Unscoped().Where("user_id=? AND id=?", uid, n.ID).Delete(&models.Note{}).Error
	if dErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to delete note")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
	})
}
