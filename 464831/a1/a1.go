package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Reset goroutine stats
	runtime.GC()
	var start = time.Now()

	// Dynamic growth
	dynamicSlice := make([]int, 0)
	for i := 0; i < 1000000; i++ {
		dynamicSlice = append(dynamicSlice, i)
	}

	// Preallocation
	preallocatedSlice := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		preallocatedSlice[i] = i
	}

	duration := time.Since(start)
	fmt.Printf("Dynamic growth time: %v\n", duration)

	// Reset goroutine stats
	runtime.GC()
	start = time.Now()

	// Preallocation
	preallocatedSlice = make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		preallocatedSlice[i] = i
	}

	duration = time.Since(start)
	fmt.Printf("Preallocation time: %v\n", duration)
}