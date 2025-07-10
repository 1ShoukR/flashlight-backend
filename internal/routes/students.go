package routes

import (
	"github.com/1ShoukR/flashlight-backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Students(router *gin.Engine, db *gorm.DB) {
	router.GET("/students/all", handlers.IsThisWorking(db, router))
}