package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// func TestCreateTodo(t *testing.T) {
// 	r := SetUpRouter()
// 	r.POST("/todo", CreateTodo)
// 	id := xid.New().String()
// 	newTodo := todo{
// 		ID:   id,
// 		Task: "Read",
// 	}

// 	jsonValue, _ := json.Marshal(newTodo)
// 	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer((jsonValue)))
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusCreated, w.Code)
// }

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

// func TestGetTodoByID(t *testing.T) {
// 	r := SetUpRouter()
// 	r.GET("/gettodoid", GetTodoByID)
// 	req, _ := http.NewRequest("GET", "/gettodoid", bytes.NewBuffer((jsonValue)))
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusCreated, w.Code)
// }
