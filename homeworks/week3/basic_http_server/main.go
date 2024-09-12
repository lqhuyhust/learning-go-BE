package main

import (
	"fmt"
	"httpServer/controllers"
	"httpServer/models"
	"net/http"
	"os"
)

func main() {
	// Load users data
	if err := models.LoadUsers(); err != nil {
		fmt.Println("Error loading users:", err)
	}

	// Ensure folder exists
	err := os.MkdirAll("uploads", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}

	// routes
	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/user", controllers.ViewProfileHandler)
	http.HandleFunc("/user/update", controllers.UpdateProfileHandler)

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
