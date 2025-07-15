package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/1ShoukR/flashlight-backend/internal/models"
	"github.com/1ShoukR/flashlight-backend/internal/routes"
	"github.com/gin-contrib/cors"
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
	
	dsn := os.Getenv("DSN")
	log.Printf("Attempting to connect to database with DSN: %s", dsn)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}
	
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	
	log.Println("Successfully connected to database")
	log.Println("Running AutoMigrate for Student model...")
	if err := db.AutoMigrate(&models.Student{}); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}
	log.Println("AutoMigrate completed successfully")
	
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:5173", "http://localhost:5173", "http://127.0.0.1:3000", "http://localhost:5173", "https://flashlight-frontend-production.up.railway.app"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
	}))
	
	newServer := &Server{
		port:   port,    
		db:     db,      
		router: router,
	}
	
	// Set up routes after migration
	routes.Students(newServer.router, db)
	
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.router,
	}
	return server
}
