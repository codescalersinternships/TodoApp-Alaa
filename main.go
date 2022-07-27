package main

import (
	"log"
	"net/http"
	_ "swag-gin-demo/docs"
	"swag-gin-demo/model"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type NewList struct {
	ID   string `json:"id" binding"required"`
	Task string `json:"task" binding:"required"`
}

type ListUpdate struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} todo
// @Router /todo [get]
func GetAllTodos(c *gin.Context) {
	var lists []model.List
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&lists).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"TodoList": lists})

}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [get]
func GetTodoByID(c *gin.Context) {

	var list model.List
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID not Found!! "})
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body todo true "todo data"
// @Success 200 {object} todo
// @Failure 400 {object} message
// @Router /todo [post]
func CreateTodo(c *gin.Context) {
	var list NewList

	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	newList := model.List{ID: list.ID, Task: list.Task}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, newList)
}

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [delete]
func DeleteTodo(c *gin.Context) {
	var deletedList model.List

	if err := model.DB.Where("id = ?", c.Param("id")).First(&deletedList).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found"})
		return
	}

	model.DB.Delete(&deletedList)
	c.JSON(http.StatusOK, gin.H{"task": true})
	// db, err := model.Database()
	// if err != nil {
	// 	log.Println(err)
	// }

	// if err := db.Where("id = ?", c.Param("id")).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "ID not Found!!"})
	// 	return
	// }

	// if err := db.Delete(&list).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "List deleted"})

}

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	router := gin.Default()
	model.Database()

	// db, err := model.Database()
	// if err != nil {
	// 	log.Println(err)
	// }

	// db.DB()

	//router.LoadHTMLFiles("my-svelte-project/")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "page.html", gin.H{
	// 		"title": "Main website",
	// 	})
	// })

	router.GET("/todo", GetAllTodos)
	router.POST("/todo", CreateTodo)
	router.GET("/todo/:id", GetTodoByID)
	router.DELETE("/todo/:id", DeleteTodo)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
