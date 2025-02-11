package main

import (
	"fmt"
	"sync"
	"time"
)

type userData struct {
	Name     string
	Email    string
	Registered bool
}

var (
	userDB     map[string]userData
	userDBMu   sync.RWMutex
	nextUserID int64
)

func generateUniqueUserID() string {
	// Use a mutex to ensure thread-safe generation of unique IDs
	nextUserID++
	return fmt.Sprintf("user_%d", nextUserID)
}

func addUser(name, email string) string {
	userID := generateUniqueUserID()
	userDBMu.Lock()
	defer userDBMu.Unlock()
	userDB[userID] = userData{Name: name, Email: email, Registered: false}
	return userID
}

func getUser(userID string) (*userData, error) {
	userDBMu.RLock()
	defer userDBMu.RUnlock()
	user, ok := userDB[userID]
	if !ok {
		return nil, fmt.Errorf("user not found: %s", userID)
	}
	return &user, nil
}

func updateUserRegistered(userID string, registered bool) error {
	userDBMu.Lock()
	defer userDBMu.Unlock()
	if user, ok := userDB[userID]; ok {
		userDB[userID] = userData{Name: user.Name, Email: user.Email, Registered: registered}
		return nil
	}
	return fmt.Errorf("user not found: %s", userID)
}

func main() {
	userDB = make(map[string]userData)

	// Adding users concurrently
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Duration(i) * time.Millisecond)
			userID := addUser(fmt.Sprintf("User %d", i+1), fmt.Sprintf("user%d@example.com", i+1))
			fmt.Printf("User added: %s\n", userID)
			wg.Done()
		}()
	}
	wg.Wait()

	// Getting user information
	for _, userID := range userDB {
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

	user, err := getUser("user_1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Updated User ID: %s, Name: %s, Email: %s, Registered: %t\n", user.ID, user.Name, user.Email, user.Registered)
	}
}
