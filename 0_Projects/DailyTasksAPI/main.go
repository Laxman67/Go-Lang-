package main

import (
	"fmt"
	"github/laxman67/handlers"
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
	fmt.Println("Database connection established successfully!")

	//Connect DB Ends======================
	mux := http.NewServeMux()

	// Register routes with HTTP methods and path parameters (Go 1.22+ syntax)
	mux.HandleFunc("GET /", handlers.WelcomeRouteHandler)

	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(db, w, r)
	})
	mux.HandleFunc("GET /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTaskByIDHandler(db, w, r)
	})
	mux.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		log.Println("POST /tasks route hit")
		handlers.CreateTaskHandler(db, w, r)
	})

	mux.HandleFunc("DELETE /tasks/{Id}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Delete /tasks route hit")
		handlers.DeleteTaskHandler(db, w, r)
	})

	println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

}
