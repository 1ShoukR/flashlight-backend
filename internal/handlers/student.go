package handlers

import (
	"fmt"
	"net/http"

	"github.com/1ShoukR/flashlight-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserData struct {
	Name 	string		`json:"name" binding:"required"`
	Grade	uint		`json:"grade" binding:"required"`
}

func GetAllStudents(db *gorm.DB, router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var students []models.Student
		
		result := db.Find(&students)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
			return
		}
		if len(students) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No students found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": "This is working! ", "students": students})
	}
}

func CreateStudent(db *gorm.DB, router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UserData
		if err := c.ShouldBindJSON(&req); err != nil {
			errString := fmt.Sprintf("Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": errString})
			return
		}
		var newUser models.Student
		newUser.Grade = req.Grade
		newUser.Name = req.Name
		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
			return
		}
		c.JSON(200, gin.H{"status": "success", "message": "user has bee created!", "data": newUser})
		
	}
}