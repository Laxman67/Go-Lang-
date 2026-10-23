package main

import (
	"encoding/json"
	"fmt"
	"github/laxman67/models"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {

	// 1. Connection string (DSN)
	dsn := "host=localhost user=postgres password=secret dbname=postgres port=5432 sslmode=disable"
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
	fmt.Println("Database connection established successfully!", db)

	//Connect DB Ends======================
	mux := http.NewServeMux()

	// Register routes with HTTP methods and path parameters (Go 1.22+ syntax)
	mux.HandleFunc("GET /tasks", getTasksHandler)
	// mux.HandleFunc("GET /tasks/{id}", getTaskByIDHandler)
	// mux.HandleFunc("POST /tasks", createTaskHandler)

}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []models.Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}
