package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
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

// FetchUsers simulates fetching users from a database with a retry mechanism
func FetchUsers(retries int, initialBackoff time.Duration) ([]User, error) {
	// Simulated data
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	for attempt := 0; attempt <= retries; attempt++ {
		// Simulated transient error
		if attempt%2 == 0 {
			return nil, WrapError("FetchUsers", errors.New("temporary database error"))
		}
		if len(users) == 0 {
			return nil, WrapError("FetchUsers", ErrNoUsersFound)
		}
		return users, nil
	}

	return nil, WrapError("FetchUsers", errors.New("maximum retries reached"))
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

// RetryWithExponentialBackoff retries a function with exponential backoff
func RetryWithExponentialBackoff(retries int, initialBackoff time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	for attempt := 0; attempt <= retries; attempt++ {
		result, err := fn()
		if err == nil {
			log.Printf("Attempt %d successful: %v\n", attempt, result)
			return result, nil
		}

		log.Printf("Attempt %d failed: %v. Retrying...\n", attempt, err)

		backoff := math.Pow(2.0, float64(attempt)) * initialBackoff.Seconds()
		time.Sleep(time.Duration(backoff))
	}
	log.Printf("Maximum retries reached. Aborting.\n")
	return nil, errors.New("maximum retries reached")
}

// UserDashboardHandler handles the /user/dashboard endpoint
func UserDashboardHandler(w http.ResponseWriter, r *http.Request) {
	retries := 3
	initialBackoff := 100 * time.Millisecond

	users, err := RetryWithExponentialBackoff(retries, initialBackoff, func() ([]User, error) {
		return FetchUsers(retries, initialBackoff)
	})
	if err != nil {
		log.Printf("Error: %v\n", err)
		http.Error(w, "Unable to fetch users. Please try again later.", http.StatusInternalServerError)
		return
	}

	userID := 3 // Simulate an invalid user ID
	user, err := FindUser(users.([]User), userID)
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