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

func GetAllTodosHandler()([]TodoList, error){
	db, err := ConnectDB()
	if err != nil{
		return nil, err
	}
	var allTodos []TodoList
	errr := db.Find(&allTodos).Error
	return allTodos, errr
}

func (newTodo *TodoList)CreateTodoHandler()(*TodoList){
	db, err := ConnectDB()
	if err != nil{
		return nil
	}
	db.Create(&newTodo)
	return newTodo
}

func DeleteTodoHandler(id string)(TodoList, error){
	deletedTodo := TodoList{}
	db, err := ConnectDB()
	if err != nil{
		return deletedTodo, err
	}
	err2 := db.Where("ID = ?",id).Delete(&deletedTodo).Error
	return deletedTodo, err2
}

func GetTodoByIDHandler(id string)(TodoList, error){
	getTask := TodoList{}
	db, err := ConnectDB()
	if err != nil{
		return getTask, err
	}
	errr := db.Where("ID=?", id).First(&getTask).Error
	return getTask, errr
}
