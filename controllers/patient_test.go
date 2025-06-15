package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Himneesh-Kalra/makerble-coding-assessment/controllers"
	"github.com/Himneesh-Kalra/makerble-coding-assessment/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupRouterWithPatientRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/patients", controllers.CreatePatient(db))
	r.GET("/patients", controllers.GetAllPatients(db))
	r.GET("/patients/:id", controllers.GetPatientByID(db))
	r.PUT("/patients/:id", controllers.UpdatePatient(db))
	r.DELETE("/patients/:id", controllers.DeletePatient(db))
	return r
}

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=hospital_db port=5432 sslmode=disable"
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}
	testDB.AutoMigrate(&models.Patient{})
	testDB.Exec("DELETE FROM patients")
}

func TestCreatePatient(t *testing.T) {
	r := setupRouterWithPatientRoutes(testDB)
	patient := models.Patient{
		FirstName: "John",
		LastName:  "Doe",
		Age:       35,
		Gender:    "male",
		Diagnosis: "Flu",
	}
	body, _ := json.Marshal(patient)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
}

func TestGetAllPatients(t *testing.T) {
	r := setupRouterWithPatientRoutes(testDB)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patients", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestGetPatientByID(t *testing.T) {
	
	patient := models.Patient{
		FirstName: "Mark",
		LastName:  "Smith",
		Age:       40,
		Gender:    "male",
		Diagnosis: "Cough",
	}
	testDB.Create(&patient)

	r := setupRouterWithPatientRoutes(testDB)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patients/"+strconv.Itoa(int(patient.ID)), nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestUpdatePatient(t *testing.T) {
	
	patient := models.Patient{
		FirstName: "Anna",
		LastName:  "Lee",
		Age:       25,
		Gender:    "female",
		Diagnosis: "Headache",
	}
	testDB.Create(&patient)

	updated := models.Patient{
		FirstName: "Anna",
		LastName:  "Smith",
		Age:       26,
		Gender:    "female",
		Diagnosis: "Migraine",
	}
	body, _ := json.Marshal(updated)

	r := setupRouterWithPatientRoutes(testDB)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/patients/"+strconv.Itoa(int(patient.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestDeletePatient(t *testing.T) {
	
	patient := models.Patient{
		FirstName: "Ella",
		LastName:  "Brown",
		Age:       30,
		Gender:    "female",
		Diagnosis: "Allergy",
	}
	testDB.Create(&patient)

	r := setupRouterWithPatientRoutes(testDB)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/patients/"+strconv.Itoa(int(patient.ID)), nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
