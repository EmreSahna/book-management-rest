package models

import (
	"github.com/EmreSahna/go_mysql_book_management/pkg/config"
	"github.com/jinzhu/gorm"
)

var dbForUser *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	dbForUser = config.GetDB()
	dbForUser.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	dbForUser.NewRecord(u)
	dbForUser.Create(&u)
	return u
}

func GetUserByUsernameAndPassword(username string, password string) (*User, *gorm.DB) {
	var user User
	db := dbForUser.Where("username=?", username).Where("password=?", password).Find(&user)
	return &user, db
}
