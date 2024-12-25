package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Custom error types for specific scenarios
type DashError struct {
	Context string
	Err     error
}

func (e *DashError) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

func WrapError(context string, err error) error {
	return &DashError{Context: context, Err: err}
}

var (
	ErrNoUsersFound     = errors.New("no users found in the database")
	ErrInvalidSliceIndex = errors.New("slice index out of bounds")
)

// User represents a user in the dashboard
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Fallback mechanism: Default user
var fallbackUser = User{ID: 999, Name: "Default User"}

// FetchUsers simulates fetching users from a database
func FetchUsers() ([]User, error) {
	// Simulated data
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	// Simulated edge case: Empty user list
	if len(users) == 0 {
		return nil, WrapError("FetchUsers", ErrNoUsersFound)
	}
	return users, nil
}

// FindUser searches for a user by ID, providing a fallback user if not found
func FindUser(users []User, userID int) (User, error) {
	for _, user := range users {
		if user.ID == userID {
			return user, nil
		}
	}
	// Provide fallback user
	log.Printf("FindUser: User with ID %d not found. Using fallback user.\n", userID)
	return fallbackUser, WrapError("FindUser", fmt.Errorf("user with ID %d not found", userID))
}

// CalculateSum computes the sum of a slice safely
func CalculateSum(data []int) (int, error) {
	if data == nil {
		return 0, WrapError("CalculateSum", errors.New("nil slice"))
	}
	total := 0
	for _, value := range data {
		total += value
	}
	return total, nil
}

// UserDashboardHandler handles the /user/dashboard endpoint
func UserDashboardHandler(w http.ResponseWriter, r *http.Request) {
	users, err := FetchUsers()
	if err != nil {
		log.Printf("Error: %v\n", err)
		http.Error(w, "Unable to fetch users. Please try again later.", http.StatusInternalServerError)
		return
	}

	userID := 3 // Simulate an invalid user ID
	user, err := FindUser(users, userID)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	data := []int{1, 2, 3}
	sum, err := CalculateSum(data)
	if err != nil {
		log.Printf("Error: %v\n", err)
		http.Error(w, "Unable to calculate slice sum.", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"user": user,
		"sum":  sum,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/user/dashboard", UserDashboardHandler)
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}