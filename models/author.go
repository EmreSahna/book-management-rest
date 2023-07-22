package models

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
