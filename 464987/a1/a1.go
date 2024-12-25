package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// User represents a user in the dashboard.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUsers fetches a list of users from a mock API.
func GetUsers() ([]User, error) {
	// Simulate a mock API call
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	// Simulate an edge case with a nil response
	if len(users) == 0 {
		return nil, ErrNoUsersFound
	}

	return users, nil
}

// ErrNoUsersFound is a custom error type for scenarios where no users are found.
type ErrNoUsersFound struct {
	message string
}

func (e ErrNoUsersFound) Error() string {
	return e.message
}

func newNoUsersFoundError(message string) error {
	return &ErrNoUsersFound{message}
}

// ErrInvalidUserID is a custom error type for invalid user IDs.
type ErrInvalidUserID struct {
	id      int
	message string
}

func (e ErrInvalidUserID) Error() string {
	return fmt.Sprintf("Invalid user ID %d: %s", e.id, e.message)
}

func newInvalidUserIDError(id int, message string) error {
	return &ErrInvalidUserID{id, message}
}

// UserDashboardHandler handles the /user/dashboard endpoint.
func UserDashboardHandler(w http.ResponseWriter, r *http.Request) {
	users, err := GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err == ErrNoUsersFound {
			log.Printf("Detailed Log: %v\n", err)
			w.Write([]byte("No users found in the database."))
		} else {
			log.Printf("Detailed Log: %v\n", err)
			w.Write([]byte("An internal server error occurred."))
		}
		return
	}

	fallbackUsers := []User{
		{ID: 999, Name: "Default User"},
	}

	// Simulate a scenario where user not found
	userID := 3 // Non-existent user ID
	user, err := FindUser(users, userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Detailed Log: %v\n", err)
		w.Write([]byte("User not found."))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Detailed Log: Failed to encode user data: %v\n", err)
		w.Write([]byte("An error occurred while encoding the user data."))
		return
	}
}

// FindUser searches for a user by ID and returns an error if not found.
func FindUser(users []User, userID int) (*User, error) {
	for _, user := range users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, newInvalidUserIDError(userID, "User not found")
}

func main() {
	http.HandleFunc("/user/dashboard", UserDashboardHandler)
	log.Println("Dashboard server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}