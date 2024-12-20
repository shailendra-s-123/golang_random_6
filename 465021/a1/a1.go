package main

import (
	"fmt"
	"testing"
)

// MapOperations is a struct that provides operations on a map.
type MapOperations struct {
	Data map[int]string
}

// Get retrieves the value associated with the given key.
func (m *MapOperations) Get(key int) string {
	return m.Data[key]
}

// Set stores the value for the given key in the map.
func (m *MapOperations) Set(key int, value string) {
	m.Data[key] = value
}

// Delete removes the key-value pair from the map.
func (m *MapOperations) Delete(key int) {
	delete(m.Data, key)
}

// TestMapOperations is a testing function to ensure MapOperations works correctly.
func TestMapOperations(t *testing.T) {
	m := &MapOperations{Data: make(map[int]string)}

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

	// Delete operation
	m.Delete(1)
	if m.Get(1) != "" {
		t.Error("Delete failed for key 1")
	}

	// Test for concurrency
	var wg sync.WaitGroup
	var mapValue string

	wg.Add(2)

	go func() {
		m.Set(1, "new_apple")
		wg.Done()
	}()

	go func() {
		mapValue = m.Get(1)
		wg.Done()
	}()

	wg.Wait()

	if mapValue != "new_apple" {
		t.Error("Concurrency test failed for key 1")
	}

	// Test data consistency after multiple updates
	m.Set(1, "apple")
	m.Set(2, "banana")
	m.Set(3, "cherry")

	if m.Get(1) != "apple" || m.Get(2) != "banana" || m.Get(3) != "cherry" {
		t.Error("Data consistency test failed")
	}

	fmt.Println("All tests passed!")
}