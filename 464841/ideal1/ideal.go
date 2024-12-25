package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeSlice struct {
	mu    sync.Mutex
	slice []int
}

// Append safely appends an element to the slice.
func (ss *SafeSlice) Append(val int) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.slice = append(ss.slice, val)
}

// Get safely retrieves an element from the slice.
func (ss *SafeSlice) Get(index int) (int, bool) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	if index < 0 || index >= len(ss.slice) {
		return 0, false
	}
	return ss.slice[index], true
}

// ProfileSliceOperations profiles slice append and iteration operations.
func ProfileSliceOperations(sliceSize int, concurrencyLevel int) {
	ss := SafeSlice{slice: make([]int, 0, sliceSize)}
	var wg sync.WaitGroup

	// Start concurrent appends
	start := time.Now()
	wg.Add(concurrencyLevel)
	for i := 0; i < concurrencyLevel; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < sliceSize/concurrencyLevel; j++ {
				ss.Append(rand.Intn(1000))
			}
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time taken for concurrent appends: %v\n", elapsed)

	// Start concurrent reads
	start = time.Now()
	wg.Add(concurrencyLevel)
	for i := 0; i < concurrencyLevel; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < sliceSize/concurrencyLevel; j++ {
				if val, ok := ss.Get(rand.Intn(sliceSize)); ok {
					_ = val // Simulate processing the value
				}
			}
		}(i)
	}
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Printf("Time taken for concurrent reads: %v\n", elapsed)
}

func main() {
	// Profile operations
	ProfileSliceOperations(1000000, 10)
}