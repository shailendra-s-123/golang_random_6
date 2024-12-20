package main

import (
	"fmt"
)

// Function to demonstrate converting an array to a slice.
func arrayToSlice() {
	// Declare a fixed-size array
	var intArray [4]int
	intArray[0], intArray[1], intArray[2], intArray[3] = 10, 20, 30, 40

	// Convert the array to a slice (creates a reference to the underlying array)
	intSlice := intArray[:]

	// Print original array and slice
	fmt.Println("Original Array:", intArray)
	fmt.Println("Converted Slice:", intSlice)

	// Modify slice (modifies the underlying array as well)
	intSlice[1] = 200
	fmt.Println("Modified Slice:", intSlice)
	fmt.Println("Modified Array:", intArray)
}

// Function to demonstrate converting a slice to an array (sized fixed-length array).
func sliceToArray() {
	// Declare and initialize a slice
	intSlice := []int{10, 20, 30, 40}

	// Convert the slice to an array (does not change the original slice)
	var intArray [4]int
	copy(intArray[:], intSlice) // Copy values from slice to array

	// Print the original slice and the converted array
	fmt.Println("Original Slice:", intSlice)
	fmt.Println("Converted Array:", intArray)

	// Modify array (does not affect the slice)
	intArray[0] = 100
	fmt.Println("Modified Array:", intArray)
	fmt.Println("Original Slice (Unchanged):", intSlice)
}

// Function to demonstrate appending to a slice.
func sliceAppending(s *[]int) {
	// Append a new element to the slice
	*s = append(*s, 50)
}

// Function to demonstrate the behavior when slices are passed to functions.
func arrayPassed() {
	// Initialize slices
	a := []int{10, 20, 30}
	b := []int{10, 20, 30} // Independent slice copy

	fmt.Println("Original A:", a)

	// Pass slice a to function (modifies the original)
	functionWithArray(a)

	// Pass slice b to function (modifies a copy)
	functionWithArrayCopy(b)

	// Print modified slices
	fmt.Println("Modified A:", a)
	fmt.Println("Modified B:", b)
}

// Function that modifies the original slice.
func functionWithArray(s []int) {
	s[0] = 100 // Modify the original slice
}

// Function that modifies a copy of the slice.
func functionWithArrayCopy(s []int) {
	s = append(s, 100) // Adding value to a new slice (no impact on the original)
}

func main() {
	fmt.Println("\nArray to Slice Example:")
	arrayToSlice()

	fmt.Println("\nSlice to Array Example:")
	sliceToArray()

	fmt.Println("\nAppending to Slice:")
	s := []int{1, 2, 3}
	fmt.Println("Original Slice:", s)
	sliceAppending(&s)
	fmt.Println("Slice After Appending:", s)

	fmt.Println("\nArray Passing to Functions:")
	arrayPassed()
}
