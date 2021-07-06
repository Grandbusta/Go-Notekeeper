package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Content string `json:"content"`
	UserId  uint   `json:"user_id"`
}
