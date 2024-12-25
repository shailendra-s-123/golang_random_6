
package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Registered bool
}

var (
	userDB     = make(map[string]User) // User map
	userDBMu   sync.RWMutex            // Mutex for safe concurrent access
	nextUserID int64                   // Counter for unique user IDs
)

// GenerateUniqueUserID creates a unique user ID by incrementing the counter
func generateUniqueUserID() string {
	userDBMu.Lock() // Lock for thread-safe ID generation
	defer userDBMu.Unlock()
	nextUserID++
	return fmt.Sprintf("user_%d", nextUserID)
}

// AddUser adds a new user to the map with a unique ID
func addUser(name, email string) string {
	userID := generateUniqueUserID()
	userDBMu.Lock() // Lock for safe writing to the map
	defer userDBMu.Unlock()
	userDB[userID] = User{Name: name, Email: email, Registered: false}
	return userID
}

// GetUser retrieves user information based on userID
func getUser(userID string) (*User, error) {
	userDBMu.RLock() // Read lock for thread-safe reading from the map
	defer userDBMu.RUnlock()
	user, ok := userDB[userID]
	if !ok {
		return nil, fmt.Errorf("user not found: %s", userID)
	}
	return &user, nil
}

// UpdateUserRegistered updates the registration status of a user
func updateUserRegistered(userID string, registered bool) error {
	userDBMu.Lock() // Lock for safe writing to the map
	defer userDBMu.Unlock()
	if user, ok := userDB[userID]; ok {
		userDB[userID] = User{Name: user.Name, Email: user.Email, Registered: registered}
		return nil
	}
	return fmt.Errorf("user not found: %s", userID)
}

func main() {
	// Adding users concurrently
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			userID := addUser(fmt.Sprintf("User %d", i+1), fmt.Sprintf("user%d@example.com", i+1))
			fmt.Printf("User added: %s\n", userID)
		}(i)
	}
	wg.Wait()

	// Retrieving and printing user information
	for userID := range userDB {
		user, err := getUser(userID)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("User ID: %s, Name: %s, Email: %s, Registered: %t\n", userID, user.Name, user.Email, user.Registered)
		}
	}

	// Updating user registration status
	err := updateUserRegistered("user_1", true)
	if err != nil {
		fmt.Println(err)
	}

	// Retrieving updated user information
	user, err := getUser("user_1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Updated User ID: %s, Name: %s, Email: %s, Registered: %t\n", user.ID, user.Name, user.Email, user.Registered)
	}
}
