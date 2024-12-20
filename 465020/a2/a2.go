package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// UserDataMap is an extended map type with versioning, logging, atomic operations, snapshots, and concurrency control.
type UserDataMap struct {
	data     map[string]interface{}
	version  int64
	logger   *log.Logger
	mutex    *sync.RWMutex
	snapshot map[string]interface{}
}

// NewUserDataMap initializes a new UserDataMap.
func NewUserDataMap() *UserDataMap {
	return &UserDataMap{
		data:    make(map[string]interface{}),
		version: 0,
		logger:  log.New(os.Stderr, "UserDataMap: ", log.LstdFlags),
		mutex:   &sync.RWMutex{},
		snapshot: make(map[string]interface{}),
	}
}

// Get retrieves a value for a given key, logging the operation.
func (udm *UserDataMap) Get(key string) (interface{}, bool) {
	udm.mutex.RLock()
	defer udm.mutex.RUnlock()

	udm.logger.Printf("Fetching data for key: %s", key)
	return udm.data[key], udm.data[key] != nil
}

// Set sets or updates a key-value pair, logging the operation and updating the version.
func (udm *UserDataMap) Set(key string, value interface{}) {
	udm.mutex.Lock()
	defer udm.mutex.Unlock()

	udm.logger.Printf("Setting data for key: %s, value: %v", key, value)
	udm.data[key] = value
	udm.version++
	udm.snapshot[key] = value
}

// Delete removes a key-value pair, logging the operation.
func (udm *UserDataMap) Delete(key string) {
	udm.mutex.Lock()
	defer udm.mutex.Unlock()

	udm.logger.Printf("Deleting data for key: %s", key)
	delete(udm.data, key)
	delete(udm.snapshot, key)
	udm.version++
}

// GetVersion returns the current version of the map, useful for data consistency checks.
func (udm *UserDataMap) GetVersion() int64 {
	udm.mutex.RLock()
	defer udm.mutex.RUnlock()

	return udm.version
}

// Snapshot returns a copy of the current map state.
func (udm *UserDataMap) Snapshot() map[string]interface{} {
	udm.mutex.RLock()
	defer udm.mutex.RUnlock()

	snapshot := make(map[string]interface{})
	for key, value := range udm.snapshot {
		snapshot[key] = value
	}
	return snapshot
}

// Main function to demonstrate the usage of UserDataMap.
func main() {
	udm := NewUserDataMap()

	// Set user data in multiple goroutines
	go func() {
		for i := 0; i < 10; i++ {
			udm.Set("user1", map[string]string{"name": "Alice", "email": fmt.Sprintf("alice%d@example.com", i)})
			time.Sleep(time.Millisecond * 10)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			udm.Set("user2", map[string]string{"name": "Bob", "email": fmt.Sprintf("bob%d@example.com", i)})
			time.Sleep(time.Millisecond * 5)
		}
	}()

	// Get user data in another goroutine
	go func() {
		for {
			data, exists := udm.Get("user1")
			if exists {
				user := data.(map[string]string)
				fmt.Printf("User: %s, Email: %s\n", user["name"], user["email"])
			}
			time.Sleep(time.Second * 2)
		}
	}()

	// Take snapshot after a delay
	time.Sleep(time.Second * 3)
	snapshot := udm.Snapshot()
	fmt.Printf("Current Snapshot:\n%v\n", snapshot)

	// Delete user data in a goroutine
	go func() {
		udm.Delete("user1")
		time.Sleep(time.Second)
	}()

	// Get final version
	version := udm.GetVersion()
	fmt.Printf("Final Version: %d\n", version)
}