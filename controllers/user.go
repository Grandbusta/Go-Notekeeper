package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
	"notekeeper/config"
	"notekeeper/models"
	"notekeeper/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var u models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}
	_, Eerr := mail.ParseAddress(u.Email)
	if len(u.Password) <= 0 || len(u.Email) <= 0 || Eerr != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	user := models.User{Email: u.Email, Password: string(bytes)}
	exist := db.Where("email=?", u.Email).First(&u)
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to signup")
		return
	}
	if exist.RowsAffected > 0 {
		utils.RespondWithError(w, http.StatusConflict, "user already exist")
		return
	}
	result := db.Create(&user)
	if result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to signup")
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, map[string]interface{}{
		"status": http.StatusCreated,
		"data":   map[string]uint{"userId": user.ID},
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// db := config.DB
}
