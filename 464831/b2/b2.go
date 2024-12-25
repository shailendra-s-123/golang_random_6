package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Function to perform exponential growth
func exponentialGrowth(slice []int) []int {
	if cap(slice) == 0 {
		return make([]int, 1, 1)
	}
	newCap := cap(slice) * 2
	newSlice := make([]int, len(slice), newCap)
	copy(newSlice, slice)
	return newSlice
}

// Function to perform fixed-size growth
func fixedSizeGrowth(slice []int, increment int) []int {
	if cap(slice) == 0 {
		return make([]int, 0, increment)
	}
	newCap := cap(slice) + increment
	newSlice := make([]int, len(slice), newCap)
	copy(newSlice, slice)
	return newSlice
}

func main() {
	// Measure memory usage before starting the test
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)
	fmt.Println("Memory usage before test:", memstats.Alloc/1024/1024, "MB")

	// Test with exponential growth
	testExponentialGrowth()

	// Test with fixed-size growth
	testFixedSizeGrowth()

	// Measure memory usage after the test
	runtime.ReadMemStats(&memstats)
	fmt.Println("Memory usage after test: ", memstats.Alloc/1024/1024, "MB")
}

func testExponentialGrowth() {
	slice := []int{}
	const numElements = 1_000_000
	start := time.Now()
	for i := 0; i < numElements; i++ {
		slice = append(slice, i)
		if i%10000 == 0 {
			slice = exponentialGrowth(slice)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken with exponential growth: ", elapsed)
	fmt.Println("Capacity after exponential growth: ", cap(slice))
}

func testFixedSizeGrowth() {
	slice := []int{}
	const numElements = 1_000_000
	const increment = 10000
	start := time.Now()
	for i := 0; i < numElements; i++ {
		slice = append(slice, i)
		if i%increment == 0 {
			slice = fixedSizeGrowth(slice, increment)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken with fixed-size growth: ", elapsed)
	fmt.Println("Capacity after fixed-size growth: ", cap(slice))
} 