package main

import (
	"fmt"
	"sync"
	"testing"
)

// ThreadSafeMapOperations is a struct that provides synchronized map operations.
type ThreadSafeMapOperations struct {
	Data    map[int]string
	mutex   sync.RWMutex
}

// Get retrieves the value associated with the given key.
func (m *ThreadSafeMapOperations) Get(key int) string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.Data[key]
}

// Set stores the value for the given key in the map.
func (m *ThreadSafeMapOperations) Set(key int, value string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.Data[key] = value
}

// Delete removes the key-value pair from the map.
func (m *ThreadSafeMapOperations) Delete(key int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.Data, key)
}

// TestThreadSafeMapOperations ensures thread-safe MapOperations works correctly.
func TestThreadSafeMapOperations(t *testing.T) {
	// Create an instance of ThreadSafeMapOperations
	m := &ThreadSafeMapOperations{Data: make(map[int]string)}

	// Basic set and get operations
	m.Set(1, "apple")
	if m.Get(1) != "apple" {
		t.Error("Get failed for key 1")
	}

	m.Set(2, "banana")
	if m.Get(2) != "banana" {
		t.Error("Get failed for key 2")
	}

	// Edge case: non-existing key
	if m.Get(3) != "" {
		t.Error("Get failed for non-existing key 3")
	}

	// Edge case: empty map
	if len(m.Data) != 2 {
		t.Error("Expected map to have 2 items after set operations")
	}

	m.Delete(1)
	if m.Get(1) != "" {
		t.Error("Delete failed for key 1")
	}

	// Test for concurrency
	var wg sync.WaitGroup
	var mapValue string

	wg.Add(2)

	go func() {
		m.Set(4, "new_value")
		wg.Done()
	}()

	go func() {
		mapValue = m.Get(4)
		wg.Done()
	}()

	wg.Wait()

	if mapValue != "new_value" {
		t.Error("Concurrency test failed for key 4")
	}

	// Test data consistency after multiple updates
	m.Set(4, "another_value")
	m.Set(5, "cheese")

	if m.Get(4) != "another_value" || m.Get(5) != "cheese" {
		t.Error("Data consistency test failed")
	}

	fmt.Println("All tests passed!")
}