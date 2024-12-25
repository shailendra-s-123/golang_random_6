package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
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

// RetryableOperation represents an operation that can be retried
type RetryableOperation func() error

// Retry attempts the operation with exponential backoff up to the specified maximum attempts
func Retry(operation RetryableOperation, maxAttempts int, initialBackoff time.Duration) error {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err := operation(); err == nil {
			return nil
		}

		log.Printf("Operation failed (attempt %d/%d): %v\n", attempt, maxAttempts, err)

		// Exponential backoff
		backoff := time.Duration(rand.Int63n(int64(initialBackoff)))
		time.Sleep(backoff)
		initialBackoff *= 2
	}

	return errors.New("operation failed after maximum attempts")
}

// UserDashboardHandler handles the /user/dashboard endpoint
func UserDashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Retryable operation for FetchUsers
	fetchUsersOperation := func() error {
		users, err := FetchUsers()
		if err != nil {
			return WrapError("FetchUsersOperation", err)
		}
		return nil
	}

	if err := Retry(fetchUsersOperation, 3, 1*time.Second); err != nil {
		log.Printf("Error: %v\n", err)
		http.Error(w, "Unable to fetch users. Please try again later.", http.StatusInternalServerError)
		return
	}

	userID := 3 // Simulate an invalid user ID
	user, err := FindUser([]User{}, userID) // Using an empty user list for demonstration purposes
	if err != nil {