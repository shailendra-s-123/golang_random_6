package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// SafeSlice is a thread-safe wrapper for a slice
type SafeSlice struct {
	mu    sync.RWMutex
	slice []int
}

// Append adds an element to the slice in a thread-safe manner
func (ss *SafeSlice) Append(val int) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.slice = append(ss.slice, val)
}

// Get retrieves an element at a specific index in a thread-safe manner
func (ss *SafeSlice) Get(index int) (int, bool) {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	if index < 0 || index >= len(ss.slice) {
		return 0, false
	}
	return ss.slice[index], true
}

// ThreadSafeBuffer is a thread-safe buffer using a channel
type ThreadSafeBuffer struct {
	buf chan int
}

// NewThreadSafeBuffer creates a new ThreadSafeBuffer with a specified capacity
func NewThreadSafeBuffer(capacity int) *ThreadSafeBuffer {
	return &ThreadSafeBuffer{buf: make(chan int, capacity)}
}

// Write appends an element to the buffer
func (tsb *ThreadSafeBuffer) Write(val int) {
	tsb.buf <- val
}

// Read retrieves an element from the buffer
func (tsb *ThreadSafeBuffer) Read() (int, bool) {
	select {
	case val, ok := <-tsb.buf:
		return val, ok
	default:
		return 0, false
	}
}

// BenchmarkSliceAppend benchmarks the append operation on a slice
func BenchmarkSliceAppend(b *testing.B) {
	slice := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
}

// BenchmarkSliceSlicing benchmarks the slicing operation on a slice
func BenchmarkSliceSlicing(b *testing.B) {
	slice := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = slice[:i] // Use the slice to avoid unused value error
	}
}

// BenchmarkSliceIteration benchmarks the iteration operation on a slice
func BenchmarkSliceIteration(b *testing.B) {
	slice := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, val := range slice {
			_ = val // Simulate processing the value
		}
	}
}

// ProfileSliceOperations profiles slice append and iteration operations
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
	// Run benchmarks and display results
	fmt.Println("Running benchmarks...")
	fmt.Println("BenchmarkSliceAppend:", testing.Benchmark(BenchmarkSliceAppend))
	fmt.Println("BenchmarkSliceSlicing:", testing.Benchmark(BenchmarkSliceSlicing))
	fmt.Println("BenchmarkSliceIteration:", testing.Benchmark(BenchmarkSliceIteration))

	// Profile operations
	fmt.Println("\nProfiling slice operations...")
	ProfileSliceOperations(1000000, 10)

	// Demonstrate thread-safe buffer with concurrent access
	fmt.Println("\nDemonstrating thread-safe buffer...")
	tsb := NewThreadSafeBuffer(1000)
	var wg sync.WaitGroup

	start := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				tsb.Write(rand.Intn(1000))
			}
		}(i)
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				_, _ = tsb.Read()
			}
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time taken for concurrent buffer operations: %v\n", elapsed)
}
