package main

import (
	"fmt"
	"sync"
	"time"
)

// UserData represents a safe, concurrent map for user data using a read-write mutex.
type UserData struct {
	data  map[string]interface{}
	mu    sync.RWMutex
}

// NewUserData initializes a new UserData instance.
func NewUserData() *UserData {
	return &UserData{data: make(map[string]interface{})}
}

// Store stores a key-value pair in the UserData.
func (ud *UserData) Store(key string, value interface{}) {
	ud.mu.Lock()
	defer ud.mu.Unlock()
	ud.data[key] = value
}

// Load loads the value stored for the given key from UserData.
func (ud *UserData) Load(key string) (value interface{}, ok bool) {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	return ud.data[key], ok
}

// Delete deletes the key-value pair for the given key from UserData.
func (ud *UserData) Delete(key string) {
	ud.mu.Lock()
	defer ud.mu.Unlock()
	delete(ud.data, key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (ud *UserData) Range(f func(key, value interface{}) bool) {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	for key, value := range ud.data {
		if !f(key, value) {
			return
		}
	}
}

// Main is the entry point of the program.
func main() {
	// Create a new UserData instance.
	userData := NewUserData()

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
	fmt.Println("\nRemaining User Data:")
	userData.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true // Continue iterating
	})

	// Simulate high concurrency with read operations.
	fmt.Println("\nHigh Concurrency Reads:")
	go func() {
		for {
			user1Data, _ := userData.Load("user1")
			fmt.Println("User 1 Data:", user1Data)
			time.Sleep(100 * time.Microsecond)
		}
	}()

	go func() {
		for {
			user2Data, _ := userData.Load("user2")
			fmt.Println("User 2 Data:", user2Data)
			time.Sleep(100 * time.Microsecond)
		}
	}()

	// Occasionally update user1 data.
	go func() {
		for {
			user1Data, _ := userData.Load("user1")
			user1Map := user1Data.(map[string]interface{})
			user1Map["age"] = user1Map["age"].(int) + 1
			userData.Store("user1", user1Map)
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait forever.
	select {}
}