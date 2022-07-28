package main

import (
	"net/http"
	_ "swag-gin-demo/docs"
	model "swag-gin-demo/models"

	middleware "swag-gin-demo/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) GetAllTodos(c *gin.Context) {
	var lists []model.TodoList

	if result := h.db.Find(&lists); result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, &lists)

}

func (h *Handler) CreateTodo(c *gin.Context) {
	var list model.TodoList

	if err := c.BindJSON(&list); err != nil {
		return
	}

	if result := h.db.Create(&list); result.Error != nil {
		return
	}

	c.JSON(http.StatusCreated, &list)

}

func (h *Handler) DeleteTodo(c *gin.Context) {

	id := c.Param("id")

	if result := h.db.Delete(&model.TodoList{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error : ": result.Error.Error(),
		})

		return
	}

	c.Status(http.StatusOK)

}

func (h *Handler) GetTodoByID(c *gin.Context) {
	var list model.TodoList
	if err := h.db.Where("id= ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error ": "ID Not Found !!"})
		return
	}
	c.JSON(http.StatusOK, list)

}

func main() {
	router := gin.New()
	db, err := model.ConnectDB()
	if err != nil {
		panic("Failed to connect to database !!")
	}

	handler := newHandler(db)
	router.Use(middleware.GinBodyMiddleware())
	router.GET("/todo", handler.GetAllTodos)
	router.Use(middleware.GinBodyMiddleware())
	router.POST("/todo", handler.CreateTodo)
	router.Use(middleware.GinBodyMiddleware())
	router.GET("/todo/:id", handler.GetTodoByID)
	router.Use(middleware.GinBodyMiddleware())
	router.DELETE("/todo/:id", handler.DeleteTodo)
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
