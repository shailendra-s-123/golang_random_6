package main

import (
    "fmt"
)

// Convert an array to a slice
func arrayToSlice() {
    var array [4]int
    array[0] = 10
    array[1] = 20
    array[2] = 30
    array[3] = 40

    // Convert array to slice using slice notation
    slice := array[:]

    // Output the original array and slice
    fmt.Println("Original Array:", array)
    fmt.Println("Converted Slice:", slice)

    // Modify the slice; changes are reflected in the array
    slice[1] = 200
    fmt.Println("Modified Slice:", slice)
    fmt.Println("Modified Array:", array)
}

// Convert a slice to an array
func sliceToArray() {
    slice := []int{10, 20, 30, 40}

    // Convert slice to array using slice notation
    array := slice[:]

    // Output the original slice and the converted array
    fmt.Println("Original Slice:", slice)
    fmt.Println("Converted Array:", array)

    // Modify the array; changes do not affect the slice
    array[0] = 100
    fmt.Println("Modified Array:", array)
    fmt.Println("Original Slice (Unchanged):", slice)
}

// Function demonstrating shared data pitfall
func sharedDataPitfall() {
    slice1 := []int{1, 2, 3}
    slice2 := slice1[:]

    // Modify slice2; changes are reflected in slice1
    slice2[0] = 100
    fmt.Println("Slice 1:", slice1)
    fmt.Println("Slice 2:", slice2)
}

// Function demonstrating avoiding shared data with a copy
func avoidSharedData() {
    slice1 := []int{1, 2, 3}
    slice2 := make([]int, len(slice1), cap(slice1))
    copy(slice2, slice1)

    // Modify slice2; slice1 remains unchanged
    slice2[0] = 100
    fmt.Println("Slice 1:", slice1)
    fmt.Println("Slice 2:", slice2)
}

func main() {
    fmt.Println("\nArray to Slice Example:")
    arrayToSlice()

    fmt.Println("\nSlice to Array Example:")
    sliceToArray()

    fmt.Println("\nShared Data Pitfall:")
    sharedDataPitfall()

    fmt.Println("\nAvoiding Shared Data:")
    avoidSharedData()
}