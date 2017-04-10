package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "database.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	if !db.HasTable(&Survey{}) {
		db.CreateTable(&Survey{})
	}

	db.AutoMigrate(&Survey{})
}

type Survey struct {
	ID   int `gorm:"AUTO_INCREMENT,primary_key"`
	Time int64
	IP   string

	Input
}

func (s Survey) Save() {
	db.Model(&Survey{}).Save(&s)
}

func GetTotalSurveys() []Survey {
	u := []Survey{}
	db.Model(&Survey{}).Scan(&u)
	return u
}

type Input struct {
	Level           int `json:"level"`
	Refree          int `json:"refree"`
	Proportionality int `json:"proportionality"`
	Timing          int `json:"timing"`
	Morality        int `json:"morality"`
	Idea            int `json:"idea"`
	Quality         int `json:"quality"`
	Partition       int `json:"partition"`
	Points          int `json:"points"`
	Broadcast       int `json:"broadcast"`
}
