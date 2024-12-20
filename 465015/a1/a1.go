package main

import (
	"fmt"
	"sync"
)

var (
	userData = make(map[string]string)
	mu       = sync.Mutex{}
)

func addUser(userID string, userData string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := userData[userID]; exists {
		return fmt.Errorf("userID %s already exists", userID)
	}

	userData[userID] = userData
	fmt.Printf("Added user: %s\n", userID)
	return nil
}

func getUserData(userID string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	data, exists := userData[userID]
	if !exists {
		return "", fmt.Errorf("userID %s not found", userID)
	}

	return data, nil
}

func deleteUser(userID string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := userData[userID]; !exists {
		return fmt.Errorf("userID %s not found", userID)
	}

	delete(userData, userID)
	fmt.Printf("Deleted user: %s\n", userID)
	return nil
}

func main() {
	// Simulate concurrent access
	go func() {
		addUser("user1", "data1")
		getUserData("user1")
		deleteUser("user1")
	}()

	go func() {
		addUser("user2", "data2")
		getUserData("user2")
		deleteUser("user2")
	}()

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
}