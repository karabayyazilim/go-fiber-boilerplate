package models

import (
	"gorm.io/gorm"
	"karabayyazilim/src/internal/config"
)

var db *gorm.DB

func init() {
	db = config.Database()
	err := db.AutoMigrate(&User{})

	if err != nil {
		return
	}
}

type User struct {
	gorm.Model
	Name  string
	Email string
}
