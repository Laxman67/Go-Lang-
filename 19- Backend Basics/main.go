package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Task represents our data model with JSON tags
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// In-memory database mock
var tasks = []Task{
	{ID: 1, Title: "Master Go Fundamentals", Done: true},
	{ID: 2, Title: "Learn net/http and Routing", Done: false},
}

func main() {
	mux := http.NewServeMux()
	// Yeh aapka Router (ek receptionist) hai. Iska kaam hai requests ko sahi jagah bhejna.

	// Register routes with HTTP methods and path parameters (Go 1.22+ syntax)
	mux.HandleFunc("GET /tasks", getTasksHandler)
	mux.HandleFunc("GET /tasks/{id}", getTaskByIDHandler)
	mux.HandleFunc("POST /tasks", createTaskHandler)

	println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

// GET /tasks - Retrieve all tasks
func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Optional query parameter example: /tasks?done=true
	doneParam := r.URL.Query().Get("done")
	if doneParam != "" {
		// Filter logic could go here
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GET /tasks/{id} - Retrieve a single task by ID
func getTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter introduced in Go 1.22
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.NotFound(w, r)
}

// POST /tasks - Create a new task
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	// Decode incoming JSON request body
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign simple auto-increment ID
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
