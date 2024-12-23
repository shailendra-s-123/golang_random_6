package main

import (
	_"errors"
	"fmt"
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

func main() {
	// Initialize a dashboard
	var d dashboard

	// Test append
	err := d.Append("Widget A")
	if err != nil {
		fmt.Println("Error appending to dashboard:", err)
	} else {
		fmt.Println("After append:", d)
	}

	// Test get (from empty or nil dashboard)
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
	d.Append("Widget C")
	d.Append("Widget D")

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
		fmt.Println("Dashboard after update:", d)
	}
}
