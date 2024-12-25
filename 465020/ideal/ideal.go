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

// AtomicMap wraps the UserData map and a counter for atomic read-write operations.
type AtomicMap struct {
	userData *UserData
	readOps  int64
	writeOps int64
}

// NewAtomicMap initializes a new AtomicMap.
func NewAtomicMap() *AtomicMap {
	return &AtomicMap{
		userData: NewUserData(),
	}
}

// SetAtomic safely sets a key-value pair with atomic tracking of write operations.
func (a *AtomicMap) SetAtomic(key, value string) {
	a.userData.Set(key, value)
	atomic.AddInt64(&a.writeOps, 1) // Atomically increment write operations counter
}

// GetAtomic retrieves the value with atomic tracking of read operations.
func (a *AtomicMap) GetAtomic(key string) (string, bool) {
	atomic.AddInt64(&a.readOps, 1) // Atomically increment read operations counter
	return a.userData.Get(key)
}

// ReportStats provides a snapshot of the current read and write operations.
func (a *AtomicMap) ReportStats() (int64, int64) {
	return atomic.LoadInt64(&a.readOps), atomic.LoadInt64(&a.writeOps)
}

func main() {
	// Initialize the atomic map structure
	atomicMap := NewAtomicMap()

	// Simulate concurrent access
	var wg sync.WaitGroup

	// Perform some write operations
	wg.Add(1)
	go func() {
		defer wg.Done()
		atomicMap.SetAtomic("user1", "Alice")
		atomicMap.SetAtomic("user2", "Bob")
	}()

	// Perform some read operations
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, _ = atomicMap.GetAtomic("user1")
		_, _ = atomicMap.GetAtomic("user2")
	}()

	wg.Wait()

	// Report stats
	readOps, writeOps := atomicMap.ReportStats()
	fmt.Printf("Read Operations: %d, Write Operations: %d\n", readOps, writeOps)
}