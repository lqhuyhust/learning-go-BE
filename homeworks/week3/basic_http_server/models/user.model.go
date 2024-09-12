package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

type Users struct {
	sync.Mutex
	Users []User `json:"users"`
}

const DBUrl = "data/users.json"

var users Users

// load users data from json file
func LoadUsers() error {
	file, err := os.Open(DBUrl)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", DBUrl)
		} else {
			return fmt.Errorf("failed to open file: %w", err)
		}
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	err = json.Unmarshal(data, &users.Users)
	if err != nil {
		return fmt.Errorf("failed to decode users: %w", err)
	}

	return nil
}

// store users data to json file
func StoreUsers() error {
	data, err := json.Marshal(users.Users)
	fmt.Println(string(data))
	fmt.Println(err)
	if err != nil {
		fmt.Println("Error encoding users:", err)
	}

	if err := os.WriteFile(DBUrl, data, 0644); err != nil {
		fmt.Println("Error writing users:", err)
	}

	return nil
}

// generate new user ID
func GenerateID() int {
	maxID := 0
	for _, user := range users.Users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	return maxID + 1
}

// register new user
func Register(newUser User) error {
	users.Lock()
	defer users.Unlock()

	// check username exists
	for _, user := range users.Users {
		if user.Username == newUser.Username {
			return fmt.Errorf("username %s already exists", newUser.Username)
		}
	}

	// add new user
	newUser.ID = GenerateID()
	users.Users = append(users.Users, newUser)

	fmt.Println("New user registered:", newUser)

	// store users
	return StoreUsers()
}

// get user by ID
func GetUserByID(id int) (User, error) {
	users.Lock()
	defer users.Unlock()

	for _, user := range users.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("user with ID %d not found", id)
}

// update user profile
func UpdateUserProfile(id int, updatedUser User) error {
	users.Lock()
	defer users.Unlock()

	for i, user := range users.Users {
		if user.ID == id {
			users.Users[i].Username = updatedUser.Username
			users.Users[i].Password = updatedUser.Password
			users.Users[i].Profile = updatedUser.Profile
			return StoreUsers()
		}
	}

	return fmt.Errorf("user with ID %d not found", id)
}

// authenticate user
func Authenticate(username, password string) (User, error) {
	for _, user := range users.Users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("invalid username or password")
}
