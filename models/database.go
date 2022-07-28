package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoList struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

func ConnectDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database !!")
	}

	db.AutoMigrate(&TodoList{})

	return db, err
}
