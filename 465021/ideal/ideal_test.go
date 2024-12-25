package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

// Map structure that holds a map and a mutex for thread safety
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewSafeMap initializes a new SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[string]int)}
}

// Set adds a key-value pair to the map
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get retrieves the value for a given key
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, exists := sm.data[key]
	return val, exists
}

// Delete removes a key-value pair from the map
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Test functions

// TestSet tests the Set operation
func TestSet(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", 10)

	val, exists := sm.Get("key1")
	if !exists {
		t.Errorf("Expected key 'key1' to exist")
	}
	if val != 10 {
		t.Errorf("Expected value 10, got %d", val)
	}
}

// TestGet tests the Get operation
func TestGet(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", 10)

	val, exists := sm.Get("key1")
	if !exists {
		t.Errorf("Expected key 'key1' to exist")
	}
	if val != 10 {
		t.Errorf("Expected value 10, got %d", val)
	}

	// Test non-existing key
	_, exists = sm.Get("nonexistent")
	if exists {
		t.Errorf("Expected 'nonexistent' key to not exist")
	}
}

// TestDelete tests the Delete operation
func TestDelete(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", 10)
	sm.Delete("key1")

	_, exists := sm.Get("key1")
	if exists {
		t.Errorf("Expected 'key1' to be deleted")
	}
}

// TestConcurrency tests map operations under concurrent access
func TestConcurrency(t *testing.T) {
	sm := NewSafeMap()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(strconv.Itoa(i), i) // Convert int to string using strconv.Itoa
		}(i)
	}
	wg.Wait()

	// Verify all values are set correctly
	for i := 0; i < 100; i++ {
		val, exists := sm.Get(strconv.Itoa(i)) // Convert int to string using strconv.Itoa
		if !exists {
			t.Errorf("Expected key '%d' to exist", i)
		}
		if val != i {
			t.Errorf("Expected value '%d', got %d", i, val)
		}
	}
}

// TestConsistency ensures map consistency during operations
func TestConsistency(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", 10)
	sm.Set("key2", 20)

	// Verify consistency between multiple operations
	val1, exists1 := sm.Get("key1")
	val2, exists2 := sm.Get("key2")
	if !exists1 || !exists2 {
		t.Errorf("Expected both keys to exist")
	}
	if val1 != 10 || val2 != 20 {
		t.Errorf("Expected values 10 and 20, got %d and %d", val1, val2)
	}
}

func main() {
	// Run the tests
	t := &testing.T{}
	TestSet(t)
	TestGet(t)
	TestDelete(t)
	TestConcurrency(t)
	TestConsistency(t)
	if t.Failed() {
		fmt.Println("Tests failed")
	} else {
		fmt.Println("All tests passed")
	}
}
