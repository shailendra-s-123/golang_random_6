package main

import (
	"fmt"
	"sync"
	"testing"
)

// Define a custom type `dashboard` as a slice of strings
type dashboard []string

// Custom error type for slice operations
type sliceError struct {
	msg string
}

func (e *sliceError) Error() string {
	return fmt.Sprintf("slice error: %s", e.msg)
}

// Function to append an element to the dashboard
func (d *dashboard) Append(item string) error {
	if d == nil {
		return &sliceError{"cannot append to nil dashboard"}
	}
	*d = append(*d, item)
	return nil
}

// Function to get an element from the dashboard by index
func (d dashboard) Get(index int) (string, error) {
	if len(d) == 0 {
		return "", &sliceError{"dashboard is empty"}
	}
	if index < 0 || index >= len(d) {
		return "", &sliceError{"index out of bounds"}
	}
	return d[index], nil
}

// Function to update an element in the dashboard by index
func (d *dashboard) Update(index int, item string) error {
	if d == nil || len(*d) == 0 {
		return &sliceError{"dashboard is empty or nil"}
	}
	if index < 0 || index >= len(*d) {
		return &sliceError{"index out of bounds"}
	}
	(*d)[index] = item
	return nil
}

// Function to remove an element from the dashboard by index
func (d *dashboard) Remove(index int) error {
	if d == nil || len(*d) == 0 {
		return &sliceError{"dashboard is empty or nil"}
	}
	if index < 0 || index >= len(*d) {
		return &sliceError{"index out of bounds"}
	}
	*d = append((*d)[:index], (*d)[index+1:]...)
	return nil
}

// Thread-safe wrapper for dashboard operations
type dashboardManager struct {
	dashboard *dashboard
	mu        *sync.Mutex
}

func newDashboardManager() *dashboardManager {
	return &dashboardManager{
		dashboard: &dashboard{},
		mu:        &sync.Mutex{},
	}
}

// Thread-safe append
func (dm *dashboardManager) Append(item string) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	return dm.dashboard.Append(item)
}

// Thread-safe get
func (dm *dashboardManager) Get(index int) (string, error) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	return dm.dashboard.Get(index)
}

// Thread-safe update
func (dm *dashboardManager) Update(index int, item string) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	return dm.dashboard.Update(index, item)
}

// Thread-safe remove
func (dm *dashboardManager) Remove(index int) error {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	return dm.dashboard.Remove(index)
}

func main() {
	dm := newDashboardManager()

	// Test thread-safe append
	err := dm.Append("Widget A")
	if err != nil {
		fmt.Println("Error appending to dashboard:", err)
	} else {
		fmt.Println("After append:", *dm.dashboard)
	}

	// Test thread-safe get (from empty or nil dashboard)
	item, err := dm.Get(0)
	if err != nil {
		fmt.Println("Error getting from dashboard:", err)
	} else {
		fmt.Println("Item at index 0:", item)
	}

	// Test thread-safe update (invalid index)
	err = dm.Update(-1, "Widget B")
	if err != nil {
		fmt.Println("Error updating dashboard:", err)
	}

	// Append more elements and test further
	err = dm.Append("Widget C")
	err = dm.Append("Widget D")

	// Test thread-safe get (valid index)
	item, err = dm.Get(1)
	if err != nil {
		fmt.Println("Error getting from dashboard:", err)
	} else {
		fmt.Println("Item at index 1:", item)
	}

	// Test thread-safe update (valid index)
	err = dm.Update(1, "Widget E")
	if err != nil {
		fmt.Println("Error updating dashboard:", err)
	} else {
		fmt.Println("Dashboard after update:", *dm.dashboard)
	}

	// Test thread-safe remove (valid index)
	err = dm.Remove(1)
	if err != nil {
		fmt.Println("Error removing from dashboard:", err)
	} else {
		fmt.Println("Dashboard after remove:", *dm.dashboard)
	}
}