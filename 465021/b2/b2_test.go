package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

// Map structure with Get, Set, and Delete methods
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[string]int)}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, exists := sm.data[key]
	return val, exists
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Module-level tests for SafeMap

func TestSafeMap_Set(t *testing.T) {
	t.Run("Should set key-value pair", func(t *testing.T) {
		sm := NewSafeMap()
		sm.Set("key1", 10)
		val, exists := sm.Get("key1")
		if !exists || val != 10 {
			t.Errorf("Expected key1=10, got %d=%v", val, exists)
		}
	})
	t.Run("Should overwrite existing key-value pair", func(t *testing.T) {
		sm := NewSafeMap()
		sm.Set("key1", 10)
		sm.Set("key1", 20)
		val, exists := sm.Get("key1")
		if !exists || val != 20 {
			t.Errorf("Expected key1=20, got %d=%v", val, exists)
		}
	})
}

func TestSafeMap_Get(t *testing.T) {
	t.Run("Should get existing key-value pair", func(t *testing.T) {
		sm := NewSafeMap()
		sm.Set("key1", 10)
		val, exists := sm.Get("key1")
		if !exists || val != 10 {
			t.Errorf("Expected key1=10, got %d=%v", val, exists)
		}
	})
	t.Run("Should handle non-existing key", func(t *testing.T) {
		sm := NewSafeMap()
		_, exists := sm.Get("key1")
		if exists {
			t.Errorf("Expected key1 to not exist")
		}
	})
	t.Run("Should handle nil map", func(t *testing.T) {
		var sm *SafeMap
		_, exists := sm.Get("key1")
		if exists {
			t.Errorf("Expected panic for nil map")
		}
	})
}

func TestSafeMap_Delete(t *testing.T) {
	t.Run("Should delete existing key-value pair", func(t *testing.T) {
		sm := NewSafeMap()
		sm.Set("key1", 10)
		sm.Delete("key1")
		_, exists := sm.Get("key1")
		if exists {
			t.Errorf("Expected key1 to be deleted")
		}
	})
	t.Run("Should handle non-existing key", func(t *testing.T) {
		sm := NewSafeMap()
		sm.Delete("key1")
		_, exists := sm.Get("key1")
		if exists {
			t.Errorf("Expected key1 to not exist")
		}
	})
}

func TestSafeMap_Concurrency(t *testing.T) {
	t.Run("Should handle concurrent access", func(t *testing.T) {
		sm := NewSafeMap()