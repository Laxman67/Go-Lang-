package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// import any other packages you need (like strconv)
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var users = []User{
	{ID: 1, Name: "Alice Smith", Email: "alice@example.com"},
	{ID: 2, Name: "Bob Jones", Email: "bob@example.com"},
}

func main() {
	mux := http.NewServeMux()

	//Routes & Handlers
	mux.HandleFunc("GET /", homeRouteHandler)
	mux.HandleFunc("GET /user", getUsersHandler)
	mux.HandleFunc("GET /user/{id}", getUserByIdHandler)
	mux.HandleFunc("POST /user", addUserHandler)
	mux.HandleFunc("GET /user/", queryUserHandler)

	// Listen to Port

	println("Server listen to Port on : 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Success: true,
		Message: "Data Fetched Successfully",
		Data:    users,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

}

func queryUserHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")
	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Name query parameter is required",
			Data:    nil,
		})
		return
	}

	var matchedUsers []User
	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(query)) {
			matchedUsers = append(matchedUsers, user)
		}
	}
	if matchedUsers == nil {
		matchedUsers = []User{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "User(s) fetched successfully",
		Data:    matchedUsers,
	})
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	// decode incoming data using json.encode
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)

}

func homeRouteHandler(w http.ResponseWriter, r *http.Request) {
	// Create your response object
	res := Response{
		Success: true,
		Message: "This is Home Route",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
