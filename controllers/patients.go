package controllers

import (
	"net/http"
	"strconv"

	"github.com/Himneesh-Kalra/makerble-coding-assessment/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient models.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, exists := c.Get("user")
		if exists {
			patient.CreatedBy = user.(*models.User).Email
		}

		if err := db.Create(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
			return
		}

		c.JSON(http.StatusCreated, patient)
	}
}

func GetAllPatients(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patients []models.Patient
		if err := db.Find(&patients).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients"})
			return
		}

		c.JSON(http.StatusOK, patients)
	}
}

func GetPatientByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var patient models.Patient
		if err := db.First(&patient, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patient"})
			}
			return
		}

		c.JSON(http.StatusOK, patient)
	}
}

func UpdatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var patient models.Patient
		if err := db.First(&patient, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		var updatedData models.Patient
		if err := c.ShouldBindJSON(&updatedData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		patient.FirstName = updatedData.FirstName
		patient.LastName = updatedData.LastName
		patient.Age = updatedData.Age
		patient.Gender = updatedData.Gender
		patient.Diagnosis = updatedData.Diagnosis

		if err := db.Save(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
			return
		}

		c.JSON(http.StatusOK, patient)
	}
}

func DeletePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		if err := db.Delete(&models.Patient{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
	}
}
