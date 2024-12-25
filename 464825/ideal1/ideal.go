
// Package main provides an example of key differences between slices and arrays
// in Golang, focusing on how data is transferred between them and examples of conversions.
package main

import "fmt"

func main() {
	// 1. Declare and initialize an array (fixed size)
	var arr [5]int = [5]int{1, 2, 3, 4, 5} // Arrays have fixed length defined at compile-time
	fmt.Println("Original array:", arr)

	// 2. Convert array to slice
	// Slices are dynamically-sized views into arrays. A slice does not copy data;
	// instead, it references the array's underlying memory.
	sliceFromArr := arr[:]
	fmt.Println("Slice created from array:", sliceFromArr)

	// Modify the slice and observe changes in the array
	sliceFromArr[0] = 10
	fmt.Println("Modified slice:", sliceFromArr)
	fmt.Println("Array after modifying the slice:", arr) // Notice the array is also updated

	// 3. Create a slice and copy data into a new array
	// The `copy` function creates a true copy of elements into the target array or slice.
	slice := []int{6, 7, 8, 9, 10} // Slices can grow and shrink dynamically
	fmt.Println("Original slice:", slice)

	// Create an array to hold the copied data
	var newArr [5]int
	copy(newArr[:], slice) // Copy slice elements to the array. Use [:] to convert array to slice.
	fmt.Println("New array after copying from slice:", newArr)

	// 4. Modify the original slice and show it does not affect the copied array
	slice[0] = 42
	fmt.Println("Modified slice:", slice)
	fmt.Println("New array remains unchanged:", newArr)

	// 5. Convert a slice back to an array-like structure (requires a copy)
	// Note: Direct assignment from slice to array is not possible since arrays have fixed sizes.
	// You must copy the elements explicitly.

	// Example of a partial copy (using fewer elements)
	var partialArr [3]int
	copy(partialArr[:], slice)
	fmt.Println("Partial array copied from slice:", partialArr)

	// Summary of differences between slices and arrays in Go:
	// - Arrays have fixed sizes, while slices are dynamically sized.
	// - Slices are references to underlying arrays and share memory, while arrays are independent.
	// - Modifying a slice backed by an array affects the array and vice versa.
	// - Explicit copying is required to create independent copies of slices or arrays.
}


 