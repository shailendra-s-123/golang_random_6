package main

import (
	"fmt"
)

// main is the entry point of the program.
func main() {
	// Example 1: Array to Slice
	arr := [5]int{1, 2, 3, 4, 5} // Declare an array with 5 elements
	slice1 := arr[:]            // Create a slice that covers the entire array

	// Print the original array
	fmt.Println("Original array:", arr)
	// Output: Original array: [1 2 3 4 5]

	// Modify the slice
	slice1[0] = 10

	// Print the modified array and slice
	fmt.Println("Modified array after slice modification:", arr)
	fmt.Println("Modified slice:", slice1)
	// Output:
	// Modified array after slice modification: [10 2 3 4 5]
	// Modified slice: [10 2 3 4 5]

	// Both the array and the slice share the same underlying data
	// Changes to one affect the other.

	// Example 2: Slice to Array
	slice2 := []int{1, 2, 3}
	arr2 := slice2[:]  // Copy the slice's data to a new array

	// Print the original slice
	fmt.Println("Original slice:", slice2)
	// Output: Original slice: [1 2 3]

	// Modify the slice
	slice2 = append(slice2, 4)

	// Print the modified slice and the new array
	fmt.Println("Modified slice after appending:", slice2)
	fmt.Println("New array after copying slice:", arr2)
	// Output:
	// Modified slice after appending: [1 2 3 4]
	// New array after copying slice: [1 2 3]

	// When we create an array from a slice using slice notation `slice[:]`,
	// It creates a copy of the slice's underlying data to the new array.
	// Changes to the slice after this copying won't affect the array.

	// Example 3: Avoiding Unintentional Modifications
	arr3 := [3]int{1, 2, 3}
	slice3 := arr3[:] // Create a slice from the array

	// Modify the slice incorrectly (assuming it won't affect the array)
	slice3 = slice3[1:]  // This is a problem!
	slice3[0] = 100      // This will modify the middle element of arr3

	// Print the modified array
	fmt.Println("Original array after incorrect slice modification:", arr3)
	// Output: Original array after incorrect slice modification: [1 100 3]

	// To avoid this issue, be careful with slice notation and avoid 
	// modifying slices beyond their intended boundaries.

	// Best Practice:
	// If you need to make modifications to a slice without affecting the original array,
	// it's better to create a copy of the slice before modifying it.
	// You can use the `make` function to allocate a new slice with the same length as the original.
	// Then, use the `copy` function to copy the original slice's data to the new slice.

	// Example 4: Copying a Slice for Modifications
	slice4 := []int{1, 2, 3}

	// Create a copy of the slice
	slice4Copy := make([]int, len(slice4))
	copy(slice4Copy, slice4)

	// Modify the copy
	slice4Copy[0] = 100

	// Print the original slice and the copy
	fmt.Println("Original slice:", slice4)
	fmt.Println("Copied slice after modification:", slice4Copy)
	// Output:
	// Original slice: [1 2 3]
	// Copied slice after modification: [100 2 3]

	// This way, modifications to the copy won't affect the original slice.
}