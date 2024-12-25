
// Package main showcases the key distinctions between slices and arrays in Golang.
package main

import "fmt"

func main() {
	// 1. Declare and initialize an array (fixed size)
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arr)

	// 2. Convert array to slice
	// Slices share data with the underlying array, making them mutable.
	sliceFromArr := arr[:]
	fmt.Println("Slice created from array:", sliceFromArr)

	// Modifying the slice affects the array
	sliceFromArr[0] = 10
	fmt.Println("Modified slice:", sliceFromArr)
	fmt.Println("Array after modifying the slice:", arr) // Output: Array after modifying the slice: [10 2 3 4 5]

	// 3. Create a slice with make()
	// The make() function creates a slice with a specified length and capacity.
	// The capacity is optional and defaults to the length if not provided.
	slice := make([]int, 3, 5) // Length = 3, Capacity = 5
	fmt.Println("Slice created with make():", slice)
	slice[0] = 6
	slice[1] = 7
	slice[2] = 8
	fmt.Println("Appended elements to slice:", slice)

	// 4. Convert slice to array (requires copying)
	// Direct assignment from slice to array is not possible, as arrays have fixed sizes.
	// You must copy the elements explicitly.
	var newArr [5]int
	copy(newArr[:], slice)
	fmt.Println("Copied slice into new array:", newArr)

	// 5. Modify the original slice and see how it affects the copied array
	slice[0] = 42
	fmt.Println("Modified slice after copying:", slice)
	fmt.Println("Copied array remains unchanged:", newArr) // Output: Copied array remains unchanged: [6 7 8 0 0]

	// 6. Create a slice and append elements to it, then convert to array
	slice = append(slice, 9, 10)
	fmt.Println("Slice after appending elements:", slice)

	// Using copy again to demonstrate how to handle the changed slice length
	var newArr2 [5]int
	copy(newArr2[:], slice)
	fmt.Println("Copied slice (after appending) into new array:", newArr2) // Output: Copied slice (after appending) into new array: [6 7 8 9 10]

	// 7. Create a slice with slice literal and convert to array
	slice3 := []int{11, 12, 13}
	var newArr3 [3]int
	copy(newArr3[:], slice3)
	fmt.Println("Slice literal copied into new array:", newArr3)
}

