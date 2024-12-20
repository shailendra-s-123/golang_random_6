package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

type UserData struct {
    Name    string
    Email   string
    Created int64 // UNIX timestamp, for example
}

var (
    userDB    sync.Map
    nextUserID atomic.Int64
)

// Generate a unique user ID using atomic operations
func getNextUserID() int64 {
    return atomic.AddInt64(&nextUserID, 1)
}

// Add a new user and return the unique ID
func addUser(name string, email string) int64 {
    // Validate input
    if name == "" || email == "" {
        fmt.Println("Error: Name and email cannot be empty")
        return 0
    }

    // Generate a unique user ID
    userID := getNextUserID()

    // Store the new user data in the sync.Map
    userDB.Store(userID, &UserData{
        Name:    name,
        Email:   email,
        Created: atomic.LoadInt64(&nextUserID),
    })

    return userID
}

// Retrieve user data by ID
func getUser(userID int64) (*UserData, error) {
    // Retrieve user data from the sync.Map
    value, ok := userDB.Load(userID)
    if !ok {
        return nil, fmt.Errorf("user not found: %d", userID)
    }

    // Cast the loaded value to UserData type
    userData, ok := value.(*UserData)
    if !ok {
        return nil, fmt.Errorf("internal error: unexpected data type")
    }

    return userData, nil
}

// Delete a user by ID
func deleteUser(userID int64) {
    userDB.Delete(userID)
}

func main() {
    // Example usage:
    userID1 := addUser("Alice Johnson", "alice@example.com")
    fmt.Println("User added with ID:", userID1)

    userID2 := addUser("Bob Smith", "bob@example.com")
    fmt.Println("User added with ID:", userID2)

    // Retrieve user data
    user1, err := getUser(userID1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("User found:", user1)
    }

    // Delete a user
    deleteUser(userID2)
}