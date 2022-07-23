package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateTodo(t *testing.T) {
	r := SetUpRouter()
<<<<<<< HEAD
	r.POST("/todo", CreateTodo)
=======
	r.POST("/todo/create", CreateTodo)
>>>>>>> 1fe44b704d131c18344bd5959a50952f6d26a018
	id := xid.New().String()
	newTodo := todo{
		ID:   id,
		Task: "Read",
	}

	jsonValue, _ := json.Marshal(newTodo)
<<<<<<< HEAD
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer((jsonValue)))
=======
	req, _ := http.NewRequest("POST", "/todo/create", bytes.NewBuffer((jsonValue)))
>>>>>>> 1fe44b704d131c18344bd5959a50952f6d26a018
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllTodos(t *testing.T) {
	r := SetUpRouter()
<<<<<<< HEAD
	r.GET("/todo", GetAllTodos)
	req, _ := http.NewRequest("GET", "/todo", nil)
=======
	r.GET("/todo/get", GetAllTodos)
	req, _ := http.NewRequest("GET", "/todo/get", nil)
>>>>>>> 1fe44b704d131c18344bd5959a50952f6d26a018
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var newTodo []todo
	json.Unmarshal(w.Body.Bytes(), &newTodo)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, newTodo)
}

// func TestGetTodoByID(t *testing.T) {
// 	r := SetUpRouter()
// 	r.GET("/gettodoid", GetTodoByID)
// 	req, _ := http.NewRequest("GET", "/gettodoid", bytes.NewBuffer((jsonValue)))
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusCreated, w.Code)
// }
