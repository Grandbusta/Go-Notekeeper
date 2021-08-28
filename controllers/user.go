package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
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

	exist := db.Where("email=?", u.Email).First(&u)
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "unable to signup")
		return
	}
	if exist.RowsAffected > 0 {
		utils.RespondWithError(w, http.StatusConflict, "user already exist")
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	user := models.User{Email: u.Email, Password: string(bytes)}
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
	db := config.DB
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "unable to process request")
		return
	}
	if len(u.Password) <= 0 || len(u.Email) <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	user := models.User{}
	existErr := db.Model(&models.User{}).Where("email=?", u.Email).First(&user).Error
	if errors.Is(existErr, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusNotFound, "user does not exist")
		return
	}
	fmt.Println(u)
	fmt.Println(user)
	bErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if bErr != nil && bErr == bcrypt.ErrMismatchedHashAndPassword {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "user details incorrect")
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, map[string]interface{}{
		"status": http.StatusOK,
	})
}
