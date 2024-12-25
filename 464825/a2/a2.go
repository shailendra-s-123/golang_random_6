package main

import (
	"fmt"
)

func main() {
	// 1. Arrays in Go
	// - Arrays have a fixed size defined at compile-time.
	// - They store a sequence of elements of the same type.

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", numbers)

	// 2. Slices in Go
	// - Slices are dynamically sized and provide a flexible way to manage arrays.
	// - A slice is a reference to a subrange of an underlying array.

	// Create a slice from an array
	slice := numbers[:] // Slices share the same memory as the original array
	fmt.Println("Slice from array:", slice)

	// Modify the slice
	slice[0] = 10
	fmt.Println("Modified slice:", slice)
	fmt.Println("Array after modifying slice:", numbers) // Original array is modified too

	// 3. Converting a slice to an array
	// - To convert a slice to an array, you need to copy the slice's elements into an array.
	// - Direct assignment does not work because arrays have a fixed size.

	var newArray [5]int
	copy(newArray[:], slice) // The copy function copies elements from slice to array
	fmt.Println("New array after copying from slice:", newArray)

	// Modify the slice to see it doesn't affect the array
	slice[0] = 20
	fmt.Println("Modified slice after copying:", slice)
	fmt.Println("New array remains unchanged:", newArray)

	// 4. Creating a slice from scratch
	// - Slices can be created without a backing array.

	var anotherSlice []int = []int{6, 7, 8, 9, 10}
	fmt.Println("New slice created from scratch:", anotherSlice)

	// 5. Slicing a slice
	// - You can create a sub-slice that references a portion of the original slice.

	subSlice := anotherSlice[1:3] // Sub-slice referencing indices 1 to 2 (exclusive)
	fmt.Println("Sub-slice:", subSlice)

	// Modify the sub-slice
	subSlice[0] = 77
	fmt.Println("Modified sub-slice:", subSlice)
	fmt.Println("Original slice after modifying sub-slice:", anotherSlice) // Original slice is modified too

	// Summary of key differences:
	// - Arrays have fixed sizes; slices are dynamic.
	// - Slices share memory with arrays when created from them, allowing modifications to affect the original array.
	// - To convert a slice to an array, use the copy function, as direct assignment does not work.
	// - Slices can be created independently of arrays.
	// - Slicing a slice creates a reference to the original data, allowing modifications to propagate.
}