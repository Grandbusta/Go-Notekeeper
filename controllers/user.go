package controllers

import (
	"encoding/json"
	"net/http"
	"notekeeper/config"
	"notekeeper/models"
	"notekeeper/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}
	if len(u.Password) <= 0 || len(u.Email) <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	user := models.User{Email: u.Email, Password: u.Password}
	result := db.Create(&user)
	if result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to signup")
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, map[string]interface{}{
		"status": http.StatusCreated,
		"data":   map[string]uint{"userID": user.ID},
	})
}
