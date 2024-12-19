
package main  
import "fmt"

func main() {  
    // Arrays in Go are fixed-size data structures that store elements of the same data type.
    // The size of an array is specified during its declaration, and it cannot be changed later.
    array := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array:", array) // Output: Array: [1 2 3 4 5]

    // Slices in Go are dynamically-sized data structures that point to a specific portion of an array.
    // A slice does not store its data itself; instead, it contains a pointer to the first element of the array, the length of the slice, and the capacity of the underlying array.
    slice := array[:]  // Creates a slice that references the entire array
    fmt.Println("Slice:", slice)  // Output: Slice: [1 2 3 4 5]

    // Data Transfer between Slices and Arrays:
    // Data is transferred between slices and arrays using assignment statements. 
    // When you assign a slice to an array, the data in the slice is copied into the corresponding elements of the array.
    // However, when you assign an array to a slice, the slice references the same underlying data as the array.

    // Example 1: Converting Array to Slice
    // The length and capacity of the slice are set to the same as the array.
    var newArray [3]int = [3]int{10, 20, 30}
    fmt.Println("Original Array:", newArray) // Output: Original Array: [10 20 30]

    newSlice := newArray[:]
    fmt.Println("Slice from Array:", newSlice) // Output: Slice from Array: [10 20 30]

    // Modifying the slice will also modify the array since they share the same underlying data
    newSlice[1] = 40
    fmt.Println("Modified Array:", newArray)  // Output: Modified Array: [10 40 30]

    // Example 2: Converting Slice to Array
    // Creating a new array with the same length as the slice and assigning the slice to it.
    // This creates a copy of the data in the slice.
    slice2 := []int{6, 7, 8, 9}
    fmt.Println("Original Slice:", slice2)  // Output: Original Slice: [6 7 8 9]

    var arrayFromSlice [4]int
    arrayFromSlice = slice2[:] // Here, slice2[:] creates a copy of the data from slice2 to arrayFromSlice
    fmt.Println("Array from Slice:", arrayFromSlice) // Output: Array from Slice: [6 7 8 9]

    // Modifying the array will not affect the slice
    arrayFromSlice[2] = 15
    fmt.Println("Modified Array:", arrayFromSlice) // Output: Modified Array: [6 7 15 9]
    fmt.Println("Original Slice:", slice2) // Output: Original Slice: [6 7 8 9]
} 
  