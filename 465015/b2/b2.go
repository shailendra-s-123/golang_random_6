package main

import (
    "fmt"
    "sync"
)

type User struct {
    ID   int
    Name string
}

var (
    users    = make(map[int]*User)
    userLock sync.Mutex
)

func addUser(name string) {
    userLock.Lock()
    defer userLock.Unlock()

    // Generate a new unique ID for the user
    nextID := len(users) + 1

    // Create a new user and store it in the map with the unique ID
    users[nextID] = &User{
        ID:   nextID,
        Name: name,
    }
}

func getUser(id int) *User {
    userLock.Lock()
    defer userLock.Unlock()

    // Return the user associated with the given ID from the map
    return users[id]
}

func deleteUser(id int) {
    userLock.Lock()
    defer userLock.Unlock()

    // Delete the user associated with the given ID from the map
    delete(users, id)
}

func main() {
    // Add multiple users concurrently
    go addUser("Alice")
    go addUser("Bob")
    go addUser("Charlie")

    // Wait for all users to be added
    // In a real application, you would use proper synchronization
    // to wait for all goroutines to finish.
    fmt.Scanln()

    // Get and print user information
    for _, user := range users {
        fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
    }

    // Delete a user
    deleteUser(2)

    // Get and print user information after deletion
    fmt.Println("Users after deletion:")
    for _, user := range users {
        fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
    }
}