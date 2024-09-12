package controllers

import (
	"encoding/json"
	"fmt"
	"httpServer/models"
	"httpServer/utils"
	"net/http"
	"strconv"
)

// API register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// 1. Get and validate request body
	// get user from request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 2. Handle profile image
	// get pfofile image from request
	file, header, err := r.FormFile("profile")
	if err != nil {
		http.Error(w, "Failed to upload profile picture", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// store profile picture to folder uploads
	profilePath, err := utils.HandleFileUpload(file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Create new user
	newUser := models.User{
		ID:       models.GenerateID(),
		Username: username,
		Password: password,
		Profile:  profilePath,
	}

	// 4. Register new user
	err = models.Register(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User %s registered successfully!\n", username)
}

// API login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// 1. Get and validate request body
	// get user from request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	// check username and password
	user, err := models.Authenticate(username, password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome %s! Your profile is %s\n", user.Username, user.Profile)
}

// API view profile
func ViewProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// get ID from query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// get user from ID
	user, err := models.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// return user profile in json format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// API update profile
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// get ID from query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// get current user
	currentUser, err := models.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Get data from request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" {
		currentUser.Username = username
	}

	if password != "" {
		currentUser.Password = password
	}

	// handle profile image
	file, header, err := r.FormFile("profile")
	if err == nil {
		defer file.Close()

		// store profile picture to folder uploads
		profilePath, err := utils.HandleFileUpload(file, header)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		currentUser.Profile = profilePath
	}

	// create updated user
	err = models.UpdateUserProfile(id, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User %d updated successfully!\n", id)
}
