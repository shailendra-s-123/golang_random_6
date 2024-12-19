package main

import (
    "fmt"
)

func arrayToSlice() {
    // Declare a fixed-size array
    var intArray [4]int
    intArray[0] = 10
    intArray[1] = 20
    intArray[2] = 30
    intArray[3] = 40

    // Convert the array to a slice
    intSlice := intArray[:]

    // Print original array and slice values
    fmt.Println("Original Array:", intArray)
    fmt.Println("Converted Slice:", intSlice)

    // Since slice points to the underlying array, modifications are shared
    intSlice[1] = 200
    fmt.Println("Modified Slice:", intSlice)
    fmt.Println("Modified Array:", intArray)
}

func sliceToArray() {
    // Declare a slice
    intSlice := []int{10, 20, 30, 40}

    // Convert the slice to an array
    intArray := intSlice[:]

    // Print the original slice and the converted array
    fmt.Println("Original Slice:", intSlice)
    fmt.Println("Converted Array:", intArray)

    // Modifications to the array do not affect the slice
    intArray[0] = 100
    fmt.Println("Modified Array:", intArray)
    fmt.Println("Original Slice (Unchanged):", intSlice)
}

func sliceAppending(s *[]int) {
    // Append a new element to the slice
    *s = append(*s, 50)
}

func arrayPassed() {
    a := []int{10, 20, 30}
    b := []int{10, 20, 30} // Another independent copy of array

    fmt.Println("Original A:", a)

    functionWithArray(a)
    functionWithArrayCopy(b)

    fmt.Println("Modified A:", a)
    fmt.Println("Modified B:", b)
}

func functionWithArray(s []int) {
    s[0] = 100 // Modifying the original slice
}

func functionWithArrayCopy(s []int) {
    // Passing a copy of the slice; Modifications won't affect the original
    s = s[:]
    s[0] = 100
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