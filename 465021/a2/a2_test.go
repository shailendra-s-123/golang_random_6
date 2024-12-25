package main

import (
	"sync"
	"testing"
)

// SafeMap holds a map and a mutex for thread safety
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

// AssertMapContains checks if the map contains the expected key-value pair
func AssertMapContains(t *testing.T, sm *SafeMap, key string, expectedValue int) {
	t.Helper()
	value, exists := sm.Get(key)
	if !exists {
		t.Errorf("Expected key '%s' to exist", key)
	}
	if value != expectedValue {
		t.Errorf("Expected value %d for key '%s', got %d", expectedValue, key, value)
	}
}

// AssertMapDoesNotContain checks if the map does not contain the specified key
func AssertMapDoesNotContain(t *testing.T, sm *SafeMap, key string) {
	t.Helper()
	_, exists := sm.Get(key)
	if exists {
		t.Errorf("Expected key '%s' not to exist", key)
	}
}

// TestSet tests the Set operation
func TestSet(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", 10)

	AssertMapContains(t, sm, "key1", 10)
}

// TestGet tests the Get operation
func TestGet(t *testing.T) {
	sm := NewSafeMap()

	// Test nil map
	if _, exists := sm.Get("key1"); exists {
		t.Errorf("Expected key not to exist in a nil map")
	}

	sm.Set("key1", 10)
	AssertMapContains(t, sm, "key1", 10)

	// Test key not found
	_, exists := sm.Get("nonexistent")
	if exists {
		t.Errorf("Expected 'nonexistent' key to not exist")
	}
}

// TestDelete tests the Delete operation
func TestDelete(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", 10)
	AssertMapContains(t, sm, "key1", 10)

	sm.Delete("key1")
	AssertMapDoesNotContain(t, sm, "key1")
}

// TestConcurrency tests map operations under concurrent access
func TestConcurrency(t *testing.T) {
	sm := NewSafeMap()

	var wg sync.WaitGroup
	const numIterations = 1000

	for i := 0; i < numIterations; i++ {
		key := strconv.Itoa(i)
		value := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Set(key, value)
		}()
	}
	wg.Wait()

	for i := 0; i < numIterations; i++ {
		key := strconv.Itoa(i)
		value, exists := sm.Get(key)
		if !exists {
			t.Errorf("Expected key '%s' to exist", key)
		}
		if value != i {
			t.Errorf("Expected value %d for key '%s', got %d", i, key, value)
		}
	}
}

// TestConsistency ensures map consistency during operations
func TestConsistency(t *testing.T) {
	sm := NewSafeMap()
	sm.Set("key1", 10)
	sm.Set("key2", 20)

	AssertMapContains(t, sm, "key1", 10)
	AssertMapContains(t, sm, "key2", 20)

	sm.Delete("key2")
	AssertMapDoesNotContain(t, sm, "key2")
	AssertMapContains(t, sm, "key1", 10)
}

func main() {
	testing.Main()
}
