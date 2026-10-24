package handlers

import (
	"encoding/json"
	"github/laxman67/models"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func WelcomeRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Message: "Home Route /GET",
		Success: true,
		Data:    nil,
	}

	json.NewEncoder(w).Encode(res)
}

func GetTasksHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks []models.Task

	if err := db.Find(&tasks).Error; err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	res := Response{
		Message: "All tasks fetched Successfully",
		Success: true,
		Data:    tasks,
	}

	json.NewEncoder(w).Encode(res)
}

func GetTaskByIDHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var task models.Task

	result := db.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func CreateTaskHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask models.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	log.Println(r.Method)
	log.Println(r.URL.Path)
	log.Println(r.TLS)

	// Log parsed request body
	log.Printf("Request Body: %+v\n", newTask)

	// Insert into database
	result := db.Create(&newTask)
	if result.Error != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	res := Response{
		Message: "Task Added Successfully!",
		Success: true,
		Data:    newTask,
	}

	json.NewEncoder(w).Encode(res)
}

// func ErrorMessageGenerator

func DeleteTaskHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("Id")
	// id, err := strconv.Atoi(idStr)

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(`{"error": "Invalid Task ID"}`))
	// 	return
	// }
	// 3. Database se delete request run karte hain
	result := db.Delete(&models.Task{}, idStr)
	log.Println(result)
	// Check karte hain ki match record mila ya nahi
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Task not found"}`))
		return
	}

	// 4. Success message return karte hain
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Task deleted successfully"}`))

}
