package main
import (
    "fmt"
    "sync"
)

// UserData represents a safe, concurrent map for user data.
type UserData struct {
    data sync.Map
}

// Store stores a key-value pair in the UserData.
func (ud *UserData) Store(key string, value interface{}) {
    ud.data.Store(key, value)
}

// Load loads the value stored for the given key from UserData.
func (ud *UserData) Load(key string) (value interface{}, ok bool) {
    return ud.data.Load(key)
}

// Delete deletes the key-value pair for the given key from UserData.
func (ud *UserData) Delete(key string) {
    ud.data.Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (ud *UserData) Range(f func(key, value interface{}) bool) {
    ud.data.Range(f)
}

// main is the entry point of the program.
func main() {
    // Create a new UserData instance.
    userData := new(UserData)

    // Store user data concurrently from multiple goroutines.
    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        userData.Store("user1", map[string]interface{}{
            "name":  "Alice",
            "age":   25,
            "email": "alice@example.com",
        })
        wg.Done()
    }()

    go func() {
        userData.Store("user2", map[string]interface{}{
            "name":  "Bob",
            "age":   30,
            "email": "bob@example.com",
        })
        wg.Done()
    }()

    go func() {
        userData.Store("user3", map[string]interface{}{
            "name":  "Charlie",
            "age":   22,
            "email": "charlie@example.com",
        })
        wg.Done()
    }()

    // Wait for all stores to complete.
    wg.Wait()

    // Retrieve and print user data.
    user1Data, ok := userData.Load("user1")
    if ok {
        fmt.Println("User 1 Data:", user1Data)
    } else {
        fmt.Println("User 1 Data not found.")
    }

    // Delete user2 data.
    userData.Delete("user2")

    // Range over all remaining user data.
    fmt.Println("Remaining User Data:")
    userData.Range(func(key, value interface{}) bool {
        fmt.Println(key, ":", value)
        return true // Continue iterating
    })
}