package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

type userData struct {
    // Replace this with actual user data structure
    Name string
}

var (
    userDB   sync.Map
    nextKey uint64
)

func generateUniqueKey() uint64 {
    return atomic.AddUint64(&nextKey, 1)
}

func addUser(name string) uint64 {
    // Ensure name is not empty
    if name == "" {
        return 0 // Return 0 to indicate invalid input
    }
    // Generate a unique key using atomic.AddUint64
    key := generateUniqueKey()
    // Store the user data in the sync.Map
    userDB.Store(key, &userData{Name: name})
    return key
}

func getUser(key uint64) (*userData, bool) {
    // Load the user data from the sync.Map
    value, ok := userDB.Load(key)
    if !ok {
        return nil, false // User not found
    }
    // Assert the value to userData type
    user, ok := value.(*userData)
    if !ok {
        return nil, false // Internal error, unexpected data type
    }
    return user, true
}

func deleteUser(key uint64) {
    // Delete the user data from the sync.Map
    userDB.Delete(key)
}

func main() {
    // Example usage:
    key1 := addUser("Alice")
    if key1 == 0 {
        fmt.Println("Failed to add user: Alice")
    } else {
        fmt.Println("User added with key:", key1)
    }
    
    key2 := addUser("Bob")
    if key2 == 0 {
        fmt.Println("Failed to add user: Bob")
    } else {
        fmt.Println("User added with key:", key2)
    }
    
    // Get user data
    user1, ok := getUser(key1)
    if !ok {
        fmt.Println("User not found:", key1)
    } else {
        fmt.Println("User found:", user1.Name)
    }

    // Delete user
    deleteUser(key2)
}  
                     