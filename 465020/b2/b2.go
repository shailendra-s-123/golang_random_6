package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// UserData represents a robust, concurrent map for user data.
type UserData struct {
	rwMutex sync.RWMutex
	data    map[string]interface{}
	version uint64
}

// NewUserData creates a new UserData instance.
func NewUserData() *UserData {
	return &UserData{
		data: make(map[string]interface{}),
	}
}

// LoadOrStore loads the value stored for the given key from UserData.
// If the key does not exist, it stores and returns the given value.
func (ud *UserData) LoadOrStore(key string, value interface{}) interface{} {
	ud.rwMutex.Lock()
	defer ud.rwMutex.Unlock()

	if val, ok := ud.data[key]; ok {
		return val
	}

	ud.data[key] = value
	return value
}

// CompareAndSwap checks whether the current value for the given key is equal to oldval.
// If so, it atomically swaps it with newval.
func (ud *UserData) CompareAndSwap(key string, oldval, newval interface{}) bool {
	ud.rwMutex.Lock()
	defer ud.rwMutex.Unlock()

	if val, ok := ud.data[key]; ok && val == oldval {
		ud.data[key] = newval
		atomic.AddUint64(&ud.version, 1)
		return true
	}
	return false
}

// Load reads the value stored for the given key from UserData.
func (ud *UserData) Load(key string) (value interface{}, ok bool) {
	ud.rwMutex.RLock()
	defer ud.rwMutex.RUnlock()

	value, ok = ud.data[key]
	return
}

// Store stores a key-value pair in the UserData.
func (ud *UserData) Store(key string, value interface{}) {
	ud.rwMutex.Lock()
	defer ud.rwMutex.Unlock()

	ud.data[key] = value
	atomic.AddUint64(&ud.version, 1)
}

// Delete deletes the key-value pair for the given key from UserData.
func (ud *UserData) Delete(key string) {
	ud.rwMutex.Lock()
	defer ud.rwMutex.Unlock()

	delete(ud.data, key)
	atomic.AddUint64(&ud.version, 1)
}

// Version returns the current version of the UserData.
func (ud *UserData) Version() uint64 {
	return atomic.LoadUint64(&ud.version)
}
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