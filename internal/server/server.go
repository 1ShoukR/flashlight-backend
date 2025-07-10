package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/1ShoukR/flashlight-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	port   int        
	db     *gorm.DB   
	router *gin.Engine 
}

func NewServer() *http.Server {
	_ = godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("PORT")) 
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	router := gin.Default()
	newServer := &Server{
		port:   port,    
		db:     db,      
		router: router,
	}
	db.AutoMigrate(&models.Student{})
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.router,
	}
	return server
}

