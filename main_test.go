package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"swag-gin-demo/model"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllTodos(t *testing.T) {

	r := SetUpRouter()
	r.GET("/todo", GetAllTodos)
	req, _ := http.NewRequest("GET", "/todo", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestCreateTodo(t *testing.T) {
	r := SetUpRouter()
	r.POST("/todo", CreateTodo)
	new := model.List{
		ID:   "1",
		Task: "Github Actions",
	}
	jsonList, _ := json.Marshal(new)
	req, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonList))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	status := response.Body.String()
	want := "{\"id\":1,\"task\":\"Github Actions\"}\n"

	if status != want {
		t.Errorf("Error!!. Expected %q, want %q", want, status)
	}

	// if status != http.StatusOK {
	// 	t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)

	// }

}

func TestGetTodoByID(t *testing.T) {
	r := SetUpRouter()
	//r.GET("/todo/6", GetTodoByID)
	request, _ := http.NewRequest(http.MethodPost, "/todo/6", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)

	}
}

func TestDeleteTodo(t *testing.T) {
	r := SetUpRouter()
	r.GET("/todo/1", DeleteTodo)
	request := httptest.NewRequest(http.MethodPost, "/todo/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)
	}

}
