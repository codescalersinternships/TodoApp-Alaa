package main

import (
	"net/http"
	_ "swag-gin-demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

type message struct {
	Message string `json:"message"`
}

var todoList = []todo{
	{"1", "Restful Api Server"},
	{"2", "Docker Image"},
	{"3", "Tests"},
	{"4", "Postman"},
	{"5", "Github actions"},
}

// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} todo
// @Router /todo [get]
func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [get]
func GetTodoByID(c *gin.Context) {
	ID := c.Param("id")

	for _, todo := range todoList {
		if todo.ID == ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	r := message{"Todo not Found!!"}
	c.JSON(http.StatusNotFound, r)
}

// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body todo true "todo data"
// @Success 200 {object} todo
// @Failure 400 {object} message
// @Router /todo [post]
func CreateTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		r := message{"Error!! Can't Create TodoList "}
		c.JSON(http.StatusBadRequest, r)
		return
	}

	todoList = append(todoList, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [delete]
func DeleteTodo(c *gin.Context) {
	ID := c.Param("id")
	for index, todo := range todoList {
		if todo.ID == ID {
			todoList = append(todoList[:index], todoList[index+1:]...)
			r := message{"successfully deleted todo"}
			c.JSON(http.StatusOK, r)
			return
		}
	}

	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
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

	router.LoadHTMLFiles("./tempaltes", ".html")

	router.GET("/todo", GetAllTodos)
	router.POST("/todo", CreateTodo)
	router.GET("/todo/:id", GetTodoByID)
	router.DELETE("/todo/:id", DeleteTodo)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
