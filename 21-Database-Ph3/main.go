package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User model maps directly to a "users" table in Postgres
type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email;unique"`
}

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
	// 2. Connect and Open connection to database
	db, err := connectDB()
	// Handle DB Connection Error
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection established successfully!")

	// 3. Auto-Migrate: Creates the "users" table automatically based on the User struct
	db.AutoMigrate(&User{})

	// CRUD Operations

	// 1. CREATE (Insert a new user)
	newUser := User{Name: "Lakshman", Email: "lakshman@example.com"}
	db.Create(&newUser)
	fmt.Println("Created User ID:", newUser)

	// 2. READ (Find all users)
	var users []User
	db.Find(&users)
	fmt.Println("All Users:", users)

	// 3. READ ONE (Find a user by ID)
	var user User
	db.First(&user, 1) // Finds user with ID 1
	fmt.Printf("Found User: %s (%s)\n", user.Name, user.Email)

	// 4. UPDATE (Modify a user's record)
	user.Email = "updated_lakshman@example.com"
	db.Save(&user)
	fmt.Println("Updated User Email:", user.Email)

	// 5. DELETE (Remove a user)
	db.Delete(&user, user.ID)
	fmt.Printf("Deleted User ID: %d\n", user.ID)
}
