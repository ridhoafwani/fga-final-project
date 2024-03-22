package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/database"
	"github.com/ridhoafwani/fga-final-project/handlers"
	"github.com/ridhoafwani/fga-final-project/models"
	"github.com/stretchr/testify/assert"
)

func setUpAuthHandler() *handlers.AuthHandler {
	db := database.DatabaseConnection()
	return &handlers.AuthHandler{
		Db: db,
	}
}

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}



func TestWelcomeHandler(t *testing.T) {
    r := SetUpRouter()
    r.GET("/", WelcomeHandler)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterHandler(t *testing.T){
	r := SetUpRouter()
	authHandler := setUpAuthHandler()
    r.POST("/", authHandler.Register)
	requetBody := models.User{
		Username: "test123",
		Password: "12345678",
		Email: "tester@gmail.com",
		Age: 10,

	}
	jsonValue, _ := json.Marshal(requetBody)
    req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}