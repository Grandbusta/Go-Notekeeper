package models

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"content"`
	Password string `json:"password"`
}
