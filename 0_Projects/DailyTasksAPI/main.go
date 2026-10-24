package main

import (
	"fmt"
	"github/laxman67/handlers"
	"github/laxman67/middleware"
	"github/laxman67/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
	// Read from environment variables
	// 1. Connection string (DSN)
	// dsn := "host=localhost user=postgres password=secret dbname=postgres port=5432 sslmode=disable"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	//Connect DB Start======================
	db, err := connectDB()
	// Handle DB Connection Error
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established successfully!")

	//Connect DB Ends======================
	mux := http.NewServeMux()

	// Register routes with HTTP methods and path parameters (Go 1.22+ syntax)
	mux.HandleFunc("GET /", handlers.WelcomeRouteHandler)

	mux.HandleFunc("GET /tasks", middleware.Logger(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(db, w, r)
	}))
	mux.HandleFunc("GET /tasks/{id}", middleware.Logger(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTaskByIDHandler(db, w, r)
	}))
	mux.HandleFunc("POST /tasks", middleware.Logger(func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTaskHandler(db, w, r)
	}))

	mux.HandleFunc("DELETE /tasks/{Id}", middleware.Logger(func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTaskHandler(db, w, r)
	}))

	// Server Starts and listening
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}

}
