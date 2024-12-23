package main

import (
	"fmt"
	"log"
)

// Custom error types for slice operations
type SliceError struct {
	Message string
	Context  string
}

func (e SliceError) Error() string {
	return fmt.Sprintf("%s: %s", e.Context, e.Message)
}

func (e SliceError) IsOutOfBounds() bool {
	return e.Message == "out of bounds"
}

func (e SliceError) IsNilDereference() bool {
	return e.Message == "nil slice dereference"
}

func (e SliceError) IsInvalidModification() bool {
	return e.Message == "invalid modification"
}

// Function to safely append to a slice
func safeAppend(slice *[]int, value int) error {
	if slice == nil {
		return SliceError{Message: "nil slice dereference", Context: "safeAppend"}
	}
	*slice = append(*slice, value)
	return nil
}

// Function to safely get a value from a slice
func safeGet(slice *[]int, index int) (int, error) {
	if slice == nil {
		return 0, SliceError{Message: "nil slice dereference", Context: "safeGet"}
	}
	if index < 0 || index >= len(*slice) {
		return 0, SliceError{Message: "out of bounds", Context: "safeGet"}
	}
	return (*slice)[index], nil
}

// Function to safely modify a slice element
func safeModify(slice *[]int, index int, newValue int) error {
	if slice == nil {
		return SliceError{Message: "nil slice dereference", Context: "safeModify"}
	}
	if index < 0 || index >= len(*slice) {
		return SliceError{Message: "out of bounds", Context: "safeModify"}
	}
	*slice[index] = newValue
	return nil
}

// Function to safely delete an element from a slice
func safeDelete(slice *[]int, index int) error {
	if slice == nil {
		return SliceError{Message: "nil slice dereference", Context: "safeDelete"}
	}
	if index < 0 || index >= len(*slice) {
		return SliceError{Message: "out of bounds", Context: "safeDelete"}
	}
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
	return nil
}

func main() {
	var mySlice []int
	
	// Attempt to append to a nil slice
	if err := safeAppend(&mySlice, 1); err != nil {
		log.Printf("Error appending to slice: %s\n", err)
	}

	// Attempt to get a value from a nil slice
	if value, err := safeGet(&mySlice, 0); err != nil {
		log.Printf("Error getting value from slice: %s\n", err)
	} else {
		log.Printf("Value at index 0: %d\n", value)
	}

	// Attempt to modify an element in a nil slice
	if err := safeModify(&mySlice, 0, 10); err != nil {
		log.Printf("Error modifying slice element: %s\n", err)
	}

	// Attempt to delete an element from a nil slice
	if err := safeDelete(&mySlice, 0); err != nil {
		log.Printf("Error deleting slice element: %s\n", err)
	}

	// Initialize the slice
	mySlice = []int{1, 2, 3}

	// Append a new value
	if err := safeAppend(&mySlice, 4); err != nil {
		log.Printf("Error appending to slice: %s\n", err)
	} else {
		log.Printf("Slice after append: %v\n", mySlice)
	}

	// Get a value from the slice
	if value, err := safeGet(&mySlice, 1); err != nil {
		log.Printf("Error getting value from slice: %s\n", err)
	} else {
		log.Printf("Value at index 1: %d\n", value)
	}

	// Modify an element in the slice
	if err := safeModify(&mySlice, 2, 20); err != nil {
		log.Printf("Error modifying slice element: %s\n", err)
	} else {
		log.Printf("Slice after modification: %v\n", mySlice)
	}

	// Delete an element from the slice
	if err := safeDelete(&mySlice, 0); err != nil {
		log.Printf("Error deleting slice element: %s\n", err)
	} else {
		log.Printf("Slice after deletion: %v\n", mySlice)
	}
}