package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github/laxman67/models"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler handles user signup
func RegisterHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var input AuthInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid request payload"}`))
			return
		}

		var user models.User
		user.Email = input.Email
		if err := user.HashPassword(input.Password); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Failed to hash password"}`))
			return
		}

		if result := db.Create(&user); result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Email already exists or invalid data"}`))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "User registered successfully"}`))
	}
}

// LoginHandler authenticates user and returns a JWT token
func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var input AuthInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid request payload"}`))
			return
		}

		var user models.User
		if result := db.Where("email = ?", input.Email).First(&user); result.Error != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Invalid email or password"}`))
			return
		}

		if !user.CheckPassword(input.Password) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Invalid email or password"}`))
			return
		}

		// Generate JWT Token
		secret := os.Getenv("JWT_SECRET")
		claims := jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Failed to generate token"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
