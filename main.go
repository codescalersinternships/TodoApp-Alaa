package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	model "swag-gin-demo/models"
	"time"

	middleware "swag-gin-demo/middleware"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

func newApp(db *gorm.DB) *App {
	return &App{db}
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func (app *App) GetAllTodos(c *gin.Context) {
	lists , err:= model.GetAllTodosHandler()
	if err!= nil{
		r := ErrorMsg{"Error!!, Can't get all todos"}
		c.JSON(http.StatusBadRequest, r)
	}
	c.JSON(http.StatusAccepted, &lists)
}

func(app *App) CreateTodo(c *gin.Context) {
	list :=  &model.TodoList{}
	if err := c.BindJSON(&list); err != nil {
		r := ErrorMsg{"Error!!"}
		c.JSON(http.StatusBadRequest, r)
		return
	}
	todolist := list.CreateTodoHandler()
	c.JSON(http.StatusCreated, &todolist)
}

func (app *App) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	deletedList, err := model.DeleteTodoHandler(id)
	if err != nil{
		r := ErrorMsg{"Failed to delete"}
		c.JSON(http.StatusInternalServerError, r)
		return
	}
	c.JSON(http.StatusOK, deletedList)
}

func (app *App) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	task , err := model.GetTodoByIDHandler(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{ "Error ": "ID Not Found !!"})
		return	
	}
	c.JSON(http.StatusAccepted, task)
}

func(app *App)MarkCompleted(c *gin.Context){
	var list model.TodoList
	if err := app.db.Where("id= ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error ": "ID Not Found !!"})
		return
	}
	if list.Done == true {
		list.Done = false
	}else {
		list.Done = true
	}
	
	app.db.Save(&list)
	c.JSON(http.StatusAccepted, list)

}


func main() {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:  []string{"content-type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

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
	app := newApp(db)
	router.Use(middleware.GinBodyMiddleware())
	router.GET("/todo", app.GetAllTodos)
	router.Use(middleware.GinBodyMiddleware())
	router.POST("/todo", app.CreateTodo)
	router.Use(middleware.GinBodyMiddleware())
	router.GET("/todo/:id", app.GetTodoByID)
	router.Use(middleware.GinBodyMiddleware())
	router.DELETE("/todo/:id", app.DeleteTodo)
	router.Use(middleware.GinBodyMiddleware())
	router.PATCH("/todo/:id", app.MarkCompleted)
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
