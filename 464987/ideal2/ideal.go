package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"
)

// Custom error types for specific scenarios
type SliceError struct {
	Context string
	Err     error
}

func (e *SliceError) Error() string {
	return fmt.Sprintf("%s: %v", e.Context, e.Err)
}

func WrapError(context string, err error) error {
	return &SliceError{Context: context, Err: err}
}

// Constants for retries and backoff
const (
	MaxRetries     = 3
	InitialBackoff = 100 * time.Millisecond
)

// ProcessSlice processes a slice and retries in case of transient errors
func ProcessSlice(slice []int, index int) (int, error) {
	if slice == nil {
		return 0, WrapError("ProcessSlice", errors.New("slice is nil"))
	}
	if index < 0 || index >= len(slice) {
		return 0, WrapError("ProcessSlice", errors.New("index out of bounds"))
	}

	// Simulating potential transient failure
	for attempt := 0; attempt <= MaxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(math.Pow(2, float64(attempt-1))) * InitialBackoff
			time.Sleep(backoff)
			log.Printf("Retrying (%d/%d) after %v\n", attempt, MaxRetries, backoff)
		}

		// Simulating transient error with 50% success rate
		if attempt%2 == 0 {
			log.Printf("ProcessSlice transient failure on attempt %d\n", attempt)
			continue
		}

		// Success case
		log.Printf("ProcessSlice succeeded on attempt %d\n", attempt)
		return slice[index], nil
	}

	return 0, WrapError("ProcessSlice", errors.New("maximum retries reached"))
}

// Fallback function
func Fallback() int {
	log.Println("Using fallback value for slice operation.")
	return -1
}

// Dashboard logic integrating slice operation
func Dashboard() {
	slice := []int{10, 20, 30}
	index := 1

	result, err := ProcessSlice(slice, index)
	if err != nil {
		log.Printf("Error processing slice: %v\n", err)
		result = Fallback()
	}

	log.Printf("Final result: %d\n", result)
}

func main() {
	log.Println("Starting dashboard application.")
	Dashboard()
}