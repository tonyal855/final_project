package main

import (
	"bytes"
	"final_project/db"
	"final_project/server/controllers"
	"final_project/server/repositories/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUnit(t *testing.T) {
	db := db.ConnectGorm()

	photoRepo := gorm.NewPhotoRepo(db)

	socialMediaRepo := gorm.NewSocialMediaRepo(db)

	commentRepo := gorm.NewCommentRepo(db)

	userRepo := gorm.NewUserRepo(db)
	userController := controllers.NewUserController(userRepo, photoRepo, commentRepo, socialMediaRepo)

	jsonBodyReg := []byte(`{"age": 9, "email": "unittest@mail.com", "password":"password", "username":"unittest"}`)
	bodyReaderReg := bytes.NewReader(jsonBodyReg)

	jsonBodyLog := []byte(`{"email": "unittest@mail.com", "password":"password"}`)
	bodyReaderLog := bytes.NewReader(jsonBodyLog)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/users/register", userController.CreateUser)
	r.POST("/users/login", userController.Login)

	req, err := http.NewRequest(http.MethodPost, "/users/register", bodyReaderReg)

	if err != nil {
		t.Errorf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusCreated {
		t.Errorf("Expected to get status %d but instead got %d\n", http.StatusCreated, w.Code)
	}

	t.Log("SUCCESS TEST REGISTER")

	reqLogin, errL := http.NewRequest(http.MethodPost, "/users/login", bodyReaderLog)

	if errL != nil {
		t.Errorf("Couldn't create request: %v\n", err)
	}

	// Perform the request
	r.ServeHTTP(w, reqLogin)

	// Check to see if the response was what you expected
	if w.Code != http.StatusCreated {
		t.Errorf("Expected to get status %d but instead got %d\n", http.StatusCreated, w.Code)
	}

	t.Log("SUCCESS TEST LOGIN")

}
