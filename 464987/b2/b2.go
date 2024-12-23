package main

import (
	"fmt"
	"sync"
	"testing"
)

// Define a custom type `dashboard` as a slice of strings with a mutex for thread safety
type dashboard struct {
	items []string
	mu    sync.Mutex
}

// Custom error types for slice operations
type sliceError struct {
	msg string
}

func (e *sliceError) Error() string {
	return fmt.Sprintf("slice error: %s", e.msg)
}

// Function to append an element to the dashboard (thread-safe)
func (d *dashboard) Append(item string) error {
	if item == "" {
		return &sliceError{"cannot append an empty item"}
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	*d = append(*d, item)
	return nil
}

// Function to get an element from the dashboard by index (thread-safe)
func (d *dashboard) Get(index int) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.items) == 0 {
		return "", &sliceError{"dashboard is empty"}
	}
	if index < 0 || index >= len(d.items) {
		return "", &sliceError{"index out of bounds"}
	}
	return d.items[index], nil
}

// Function to update an element in the dashboard by index (thread-safe)
func (d *dashboard) Update(index int, item string) error {
	if item == "" {
		return &sliceError{"cannot update with an empty item"}
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.items) == 0 {
		return &sliceError{"dashboard is empty"}
	}
	if index < 0 || index >= len(d.items) {
		return &sliceError{"index out of bounds"}
	}
	d.items[index] = item
	return nil
}

func main() {
	// Initialize a dashboard
	var d dashboard

	// Test append
	err := d.Append("Widget A")
	if err != nil {
		fmt.Println("Error appending to dashboard:", err)
	} else {
		fmt.Println("After append:", d.items)
	}

	// Test get (from empty dashboard)
	item, err := d.Get(0)
	if err != nil {
		fmt.Println("Error getting from dashboard:", err)
	} else {
		fmt.Println("Item at index 0:", item)
	}

	// Test update (invalid index)
	err = d.Update(-1, "Widget B")
	if err != nil {
		fmt.Println("Error updating dashboard:", err)
	}

	// Append more elements and test further
	d.Append("Widget B")
	d.Append("Widget C")

	// Test get (valid index)
	item, err = d.Get(1)
	if err != nil {
		fmt.Println("Error getting from dashboard:", err)
	} else {
		fmt.Println("Item at index 1:", item)
	}

	// Test update (valid index)
	err = d.Update(1, "Widget E")
	if err != nil {
		fmt.Println("Error updating dashboard:", err)
	} else {
		fmt.Println("Dashboard after update:", d.items)
	}
}
 
 // Unit tests to validate edge cases
func TestDashboard(t *testing.T) {
	d := &dashboard{}

	t.Run("Append", func(t *testing.T) {
		err := d.Append("Widget A")
		if err != nil {
			t.Errorf("Append failed: %v", err)
		}
		if len(d.items) != 1 {
			t.Errorf("Expected 1 item, got %d", len(d.items))