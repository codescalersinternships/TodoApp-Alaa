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
	r.POST("/todo", CreateTodo)
	id := xid.New().String()
	newTodo := todo{
		ID:   id,
		Task: "Read",
	}

	jsonValue, _ := json.Marshal(newTodo)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer((jsonValue)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllTodos(t *testing.T) {
	r := SetUpRouter()
	r.GET("/todo", GetAllTodos)
	req, _ := http.NewRequest("GET", "/todo", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var newTodo []todo
	json.Unmarshal(w.Body.Bytes(), &newTodo)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, newTodo)
}
