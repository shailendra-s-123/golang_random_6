package main

import (
	"sync"
	"testing"
	"strconv" // Import the strconv package
)

// MapHandler struct to manage a map and provide thread-safe operations
type MapHandler struct {
	mu sync.RWMutex
	m  map[string]string
}

// NewMapHandler initializes a MapHandler with an empty map
func NewMapHandler() *MapHandler {
	return &MapHandler{
		m: make(map[string]string),
	}
}

// Set adds or updates a key-value pair in the map
func (mh *MapHandler) Set(key, value string) {
	mh.mu.Lock()
	defer mh.mu.Unlock()
	mh.m[key] = value
}

// Get retrieves the value associated with a key, returning an empty string if not found
func (mh *MapHandler) Get(key string) (string, bool) {
	mh.mu.RLock()
	defer mh.mu.RUnlock()
	value, exists := mh.m[key]
	return value, exists
}

// Delete removes a key-value pair from the map
func (mh *MapHandler) Delete(key string) {
	mh.mu.Lock()
	defer mh.mu.Unlock()
	delete(mh.m, key)
}

// Tests for the MapHandler functionality
func TestMapHandler(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		key      string
		value    string
		expected string
		expectOK bool
	}{
		{
			name:     "Set and Get existing key",
			action:   "set",
			key:      "key1",
			value:    "value1",
			expected: "value1",
			expectOK: true,
		},
		{
			name:     "Get key not found",
			action:   "get",
			key:      "key2",
			expected: "",
			expectOK: false,
		},
		{
			name:     "Set and Delete key",
			action:   "delete",
			key:      "key1",
			expected: "",
			expectOK: false,
		},
	}

	// Initialize MapHandler
	mh := NewMapHandler()

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.action {
			case "set":
				mh.Set(tt.key, tt.value)
			case "get":
				got, ok := mh.Get(tt.key)
				if ok != tt.expectOK || got != tt.expected {
					t.Errorf("Get() = %v, %v, want %v, %v", got, ok, tt.expected, tt.expectOK)
				}
			case "delete":
				mh.Delete(tt.key)
				got, ok := mh.Get(tt.key)
				if ok != tt.expectOK || got != tt.expected {
					t.Errorf("Delete() = %v, %v, want %v, %v", got, ok, tt.expected, tt.expectOK)
				}
			}
		})
	}
}

// Test for concurrent access
func TestConcurrentMapAccess(t *testing.T) {
	mh := NewMapHandler()

	// Run a number of concurrent Set and Get operations
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)

		go func(i int) {
			defer wg.Done()
			// Convert int to string using strconv.Itoa()
			mh.Set(strconv.Itoa(i), strconv.Itoa(i))
		}(i)

		go func(i int) {
			defer wg.Done()
			// Convert int to string using strconv.Itoa()
			mh.Get(strconv.Itoa(i))
		}(i)
	}

	wg.Wait()

	// Verify the map has all values set
	for i := 0; i < 100; i++ {
		if value, ok := mh.Get(strconv.Itoa(i)); !ok || value != strconv.Itoa(i) {
			t.Errorf("Expected value for key %d to be %d, got %v", i, i, value)
		}
	}
}
