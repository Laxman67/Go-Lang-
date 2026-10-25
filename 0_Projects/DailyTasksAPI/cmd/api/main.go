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
		handlers.WelcomeRouteHandler(c.Writer, c.Request)
	})

	router.GET("/tasks", func(c *gin.Context) {
		handlers.GetTasksHandler(db, c.Writer, c.Request)
	})

	router.GET("/tasks/:id", func(c *gin.Context) {
		handlers.GetTaskByIDHandler(db, c.Writer, c.Request)
	})

	router.POST("/tasks", func(c *gin.Context) {
		handlers.CreateTaskHandler(db, c.Writer, c.Request)
	})

	router.DELETE("/tasks/:id", func(c *gin.Context) {
		handlers.DeleteTaskHandler(db, c.Writer, c.Request)
	})

	router.POST("/register", func(c *gin.Context) {
		handlers.RegisterHandler(db)(c.Writer, c.Request)
	})

	router.POST("/login", func(c *gin.Context) {
		handlers.LoginHandler(db)(c.Writer, c.Request)
	})

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
