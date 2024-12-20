package main

import (
    "fmt"
    "sync"
)

type userData struct {
    // Replace this with actual user data structure
    Name string
}

var (
    userDB sync.Map
    nextKey int64
    mu sync.Mutex // Mutex to protect nextKey increment
)

func getNextKey() int64 {
    mu.Lock()
    defer mu.Unlock()
    nextKey++
    return nextKey
}

func addUser(name string) (int64, error) {
    // Ensure name is not empty
    if name == "" {
        return 0, fmt.Errorf("name cannot be empty")
    }
    // Generate a unique key using atomic.AddInt64
    key := getNextKey()
    // Store the user data in the sync.Map
    userDB.Store(key, &userData{Name: name})
    return key, nil
}

func getUser(key int64) (*userData, bool) {
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

func deleteUser(key int64) {
    // Delete the user data from the sync.Map
    userDB.Delete(key)
}

func main() {
    // Example usage:
    key1, err := addUser("Alice")
    if err != nil {
        panic(err)
    }
    fmt.Println("User added with key:", key1)

    key2, err := addUser("Bob")
    if err != nil {
        panic(err)
    }
    fmt.Println("User added with key:", key2)

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
