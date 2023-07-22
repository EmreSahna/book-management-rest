package models

type Book struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
