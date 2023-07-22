package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	AuthorId    uint    `json:"author_id"`
	ISBN        string  `json:"isbn"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
