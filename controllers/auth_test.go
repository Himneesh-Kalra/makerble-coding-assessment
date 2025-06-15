package controllers_test

import (
	"log"
	"os"
	"testing"

	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/Himneesh-Kalra/makerble-coding-assessment/controllers"
	"github.com/Himneesh-Kalra/makerble-coding-assessment/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error

	dsn := "host=localhost user=postgres password=postgres dbname=hospital_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test DB: %v", err)
	}

	_ = db.AutoMigrate(&models.User{})

	db.Exec("DELETE FROM users")
	code := m.Run()

	os.Exit(code)
}

func performRequest(router *gin.Engine, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestRegisterHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/register", controllers.Register(db))

	payload := `{
		"name": "Dr. Jane",
		"email": "jane@example.com",
		"password": "securepass",
		"role": "doctor"
	}`

	w := performRequest(router, "POST", "/register", []byte(payload))

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestLoginInvalidCredentials(t *testing.T) {
	router := gin.Default()
	router.POST("/login", controllers.Login(db))

	payload := `{
		"email": "wrong@example.com",
		"password": "wrongpass"
	}`

	w := performRequest(router, "POST", "/login", []byte(payload))

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}
