package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoList struct {
	ID   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func ConnectDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database !!")
	}

	db.AutoMigrate(&TodoList{})

	return db, err
}

// func getAllTodos()([]TodoList,error){
// 	var lists[] TodoList
// 	db *gorm.DB
// 	if result := db.Find(&lists); result.Error != nil {
// 		err := result.Error
// 		return nil, err

// 	}

// 	//fmt.Println(lists)
// 	return lists, nil
// }
