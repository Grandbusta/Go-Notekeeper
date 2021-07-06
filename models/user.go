package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"content"`
	Password string `json:"password"`
	Notes    []Note `json:"notes"`
}
