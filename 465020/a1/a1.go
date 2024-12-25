package main

import (
	"log"
	"sync/atomic"
	"time"
)

// UserDataMap extends a standard map to provide resilience and versioning features.
type UserDataMap struct {
	data    map[string]interface{}
	version int64
	logger  *log.Logger
}

// NewUserDataMap initializes a new UserDataMap with logging enabled.
func NewUserDataMap() *UserDataMap {
	return &UserDataMap{
		data:    make(map[string]interface{}),
		version: 0,
		logger:  log.New(log.New(os.Stderr, "UserDataMap: ", log.LstdFlags), "", 0),
	}
}

// Get retrieves a value for a given key, logging the operation.
func (udm *UserDataMap) Get(key string) (interface{}, bool) {
	udm.logger.Printf("Fetching data for key: %s", key)
	return udm.data[key], udm.data[key] != nil
}

// Set sets or updates a key-value pair, logging the operation and updating the version.
func (udm *UserDataMap) Set(key string, value interface{}) {
	udm.logger.Printf("Setting data for key: %s, value: %v", key, value)
	udm.data[key] = value
	atomic.AddInt64(&udm.version, 1)
}

// Delete removes a key-value pair, logging the operation.
func (udm *UserDataMap) Delete(key string) {
	udm.logger.Printf("Deleting data for key: %s", key)
	delete(udm.data, key)
	atomic.AddInt64(&udm.version, 1)
}

// GetVersion returns the current version of the map, useful for data consistency checks.
func (udm *UserDataMap) GetVersion() int64 {
	return udm.version
}

// Snapshot returns a copy of the current map state, with version information.
func (udm *UserDataMap) Snapshot() map[string]interface{} {
	snapshot := make(map[string]interface{})
	for key, value := range udm.data {
		snapshot[key] = value
	}
	return snapshot
}

// Main function to demonstrate the usage of UserDataMap.
func main() {
	udm := NewUserDataMap()

	// Set user data
	udm.Set("user1", map[string]string{"name": "Alice", "email": "alice@example.com"})

	// Get user data
	data, exists := udm.Get("user1")
	if exists {
		user := data.(map[string]string)
		log.Printf("User: %s, Email: %s", user["name"], user["email"])
	} else {
		log.Println("User not found")
	}

	// Update user data
	udm.Set("user1", map[string]string{"name": "Alice", "email": "alice@newdomain.com"})

	// Delete user data
	udm.Delete("user1")

	// Display current version
	log.Printf("Current version: %d", udm.GetVersion())

	// Create snapshot of the map
	snapshot := udm.Snapshot()
	log.Println("Snapshot of current state:", snapshot)
}