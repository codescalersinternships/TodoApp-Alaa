package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	model "swag-gin-demo/models"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllTodos(t *testing.T) {

	r := SetUpRouter()
	app := App{}
	r.GET("/todo", app.GetAllTodos)
	request, _ := http.NewRequest("Get", "/todo", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestCreateTodo(t *testing.T) {
	r := SetUpRouter()
	handler := App{}
	r.POST("/todo", handler.CreateTodo)
	new := model.TodoList{
		ID:   "3",
		Task: "Github Actions",
	}
	jsonList, _ := json.Marshal(new)
	req, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonList))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	status := response.Body.String()
	want := "{\"id\":3,\"task\":\"Github Actions\"}\n"

	if status != want {
		t.Errorf("Error!!. Expected %q, want %q", want, status)
	}

}

func TestGetTodoByID(t *testing.T) {
	r := SetUpRouter()
	handler := App{}
	r.GET("/todo/1", handler.GetTodoByID)
	request, _ := http.NewRequest(http.MethodPost, "/todo/1", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)

	}
}

func TestDeleteTodo(t *testing.T) {
	r := SetUpRouter()
	handler := App{}
	r.GET("/todo/1", handler.DeleteTodo)
	request := httptest.NewRequest(http.MethodPost, "/todo/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)
	}

}
