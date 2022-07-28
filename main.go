package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	_ "swag-gin-demo/docs"
	model "swag-gin-demo/models"
	"time"

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

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server exiting")
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

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	router.Run()
}
