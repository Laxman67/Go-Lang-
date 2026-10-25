package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github/laxman67/internals/database"
	"github/laxman67/internals/handlers"
	"github/laxman67/internals/models"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// Connect DB Start
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate Models
	if err := db.AutoMigrate(&models.Task{}, &models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database connection established successfully!")
	// Connect DB Ends

	// Set Gin mode (Release or Debug)
	gin.SetMode(gin.DebugMode)

	// Create Gin router with default Logger and Recovery middleware
	router := gin.Default()

	// Register routes using Gin syntax
	router.GET("/", func(c *gin.Context) {
		handlers.WelcomeRouteHandler(c)
	})

	router.GET("/tasks", handlers.GetTasksHandler(db))

	router.GET("/tasks/:id", handlers.GetTaskByIDHandler(db))

	router.POST("/tasks", handlers.CreateTaskHandler(db))

	router.DELETE("/tasks/:id", handlers.DeleteTaskHandler(db))

	router.POST("/register", handlers.RegisterHandler(db))

	router.POST("/login", handlers.LoginHandler(db))

	// Server Starts and listening using Gin's Run method
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
