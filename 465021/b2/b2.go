
package main

import (
	"fmt"
	"sync"
	"testing"
)

// MapOperations is a thread-safe struct that provides operations on a map.
type MapOperations struct {
	mu   sync.Mutex
	Data map[int]string
}

// NewMapOperations creates a new instance of MapOperations.
func NewMapOperations() *MapOperations {
	return &MapOperations{Data: make(map[int]string)}
}

// Get retrieves the value associated with the given key from the map.
func (m *MapOperations) Get(key int) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.Data[key]
}

// Set stores the value for the given key in the map.
func (m *MapOperations) Set(key int, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Data[key] = value
}

// Delete removes the key-value pair from the map.
func (m *MapOperations) Delete(key int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Data, key)
}

// TestMapOperations is a testing function to ensure MapOperations works correctly.
func TestMapOperations(t *testing.T) {
	m := NewMapOperations()

	// Test normal functionality
	m.Set(1, "apple")
	if m.Get(1) != "apple" {
		t.Error("Get failed for key 1")
	}
	m.Set(2, "banana")
	if m.Get(2) != "banana" {
		t.Error("Get failed for key 2")
	}

	// Test edge case: empty map
	if m.Get(3) != "" {
		t.Error("Get failed for non-existing key in an empty map")
	}

	// Test edge case: non-existing key
	m.Set(4, "cherry")
	if m.Get(5) != "" {
		t.Error("Get failed for non-existing key 5")
	}

	// Test concurrency
	var wg sync.WaitGroup
	var mapValue string

	wg.Add(2)

	go func() {
		m.Set(10, "new_apple")
		wg.Done()
	}()

	go func() {
		mapValue = m.Get(10)
		wg.Done()
	}()

	wg.Wait()

	if mapValue != "new_apple" {
		t.Error("Concurrency test failed for key 10")
	}

	// Test data consistency after multiple updates
	m.Set(1, "apple")
	m.Set(2, "banana")
	m.Set(3, "cherry")

	if m.Get(1) != "apple" || m.Get(2) != "banana" || m.Get(3) != "cherry" {
		t.Error("Data consistency test failed")
	}

	// Test Delete operation
	m.Delete(2)
	if m.Get(2) != "" {
		t.Error("Delete failed for key 2")
	}

	fmt.Println("All tests passed!")
}
  