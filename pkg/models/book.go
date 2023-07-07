package models

import (
	"github.com/hauselkinga/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}