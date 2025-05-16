package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxonbejenari/testWebApp/models"
	"github.com/maxonbejenari/testWebApp/utils"
	"gorm.io/gorm"
	"net/http"
)

func CreatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := utils.GetRoleFromContext(c)
		if role != string(models.Receptionist) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied",
			})
			return
		}
		var p models.Patient
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Create(&p).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, &p)
	}
}

func ListPatients(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patients []models.Patient
		if err := db.Find(&patients).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, patients)
	}
}

func GetPatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var patient models.Patient
		if err := db.First(&patient, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Patient not found",
			})
			return
		}
		c.JSON(http.StatusOK, patient)
	}
}

func UpdatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var patient models.Patient
		if err := db.First(&patient, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Patient not found",
			})
			return
		}
		role := utils.GetRoleFromContext(c)
		if role != string(models.Doctor) && role != string(models.Receptionist) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		var input models.Patient
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		patient.Name = input.Name
		patient.Age = input.Age
		patient.Address = input.Address
		patient.Phone = input.Phone
		patient.Details = input.Details
		db.Save(&patient)
		c.JSON(http.StatusOK, patient)
	}
}

func DeletePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := utils.GetRoleFromContext(c)
		if role != string(models.Receptionist) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Only receptionist can delete patients",
			})
			return
		}
		id := c.Param("id")
		if err := db.Delete(&models.Patient{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Patient deleted",
		})
	}
}
