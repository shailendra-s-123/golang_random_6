package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// UserData holds the user data and the read-write lock for concurrency control.
type UserData struct {
	sync.RWMutex
	data map[string]string
}

// NewUserData initializes a new UserData object with an empty map.
func NewUserData() *UserData {
	return &UserData{
		data: make(map[string]string),
	}
}

// Set adds or updates a key-value pair in the map with write access.
func (u *UserData) Set(key, value string) {
	u.Lock() // Exclusive lock for writing
	defer u.Unlock()
	u.data[key] = value
}

// Get retrieves the value for a given key with read access.
func (u *UserData) Get(key string) (string, bool) {
	u.RLock() // Shared lock for reading
	defer u.RUnlock()
	value, exists := u.data[key]
	return value, exists
}

// AtomicMap wraps the UserData map and atomic counters for read/write operations.
type AtomicMap struct {
	userData *UserData
	readOps  atomic.Int64 // Using atomic.Int64 for thread-safe counter operations
	writeOps atomic.Int64
}

// NewAtomicMap initializes a new AtomicMap with atomic operation counters.
func NewAtomicMap() *AtomicMap {
	return &AtomicMap{
		userData: NewUserData(),
	}
}

// SetAtomic safely sets a key-value pair with atomic tracking of write operations.
func (a *AtomicMap) SetAtomic(key, value string) {
	a.userData.Set(key, value)
	a.writeOps.Add(1) // Atomically increment write operations counter
}

// GetAtomic retrieves the value with atomic tracking of read operations.
func (a *AtomicMap) GetAtomic(key string) (string, bool) {
	a.readOps.Add(1) // Atomically increment read operations counter
	return a.userData.Get(key)
}

// ReportStats provides a snapshot of the current read and write operations.
func (a *AtomicMap) ReportStats() (int64, int64) {
	return a.readOps.Load(), a.writeOps.Load() // Use Load method for atomic access to the counter values
}

// Simulate high-concurrency environment with multiple goroutines.
func main() {
	// Initialize the atomic map structure
	atomicMap := NewAtomicMap()

	// Use WaitGroup to ensure all goroutines finish before reporting stats
	var wg sync.WaitGroup

	// Simulate high-concurrency environment with many read and write operations
	numGoroutines := 100
	wg.Add(numGoroutines)

	// Launch multiple goroutines to simulate concurrent access
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			// Alternate between read and write operations
			if id%2 == 0 {
				atomicMap.SetAtomic(fmt.Sprintf("user%d", id), fmt.Sprintf("User%d", id))
			} else {
				_, _ = atomicMap.GetAtomic(fmt.Sprintf("user%d", id-1))
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Report the statistics on read and write operations.
	readOps, writeOps := atomicMap.ReportStats()
	fmt.Printf("Read Operations: %d, Write Operations: %d\n", readOps, writeOps)
}