package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// Thread-safe buffer using a slice and a mutex
type ThreadSafeBuffer struct {
	mu   sync.Mutex
	data []int
}

func (tb *ThreadSafeBuffer) Append(val int) {
	tb.mu.Lock()
	tb.data = append(tb.data, val)
	tb.mu.Unlock()
}

func (tb *ThreadSafeBuffer) Read(index int) (int, bool) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if index < 0 || index >= len(tb.data) {
		return 0, false
	}
	return tb.data[index], true
}

func (tb *ThreadSafeBuffer) Clear() {
	tb.mu.Lock()
	tb.data = tb.data[:0]
	tb.mu.Unlock()
}

// Benchmark functions for different slice operations
func BenchmarkSliceAppend(b *testing.B) {
	slice := make([]int, 0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			slice = append(slice, rand.Intn(1000))
		}
	})
}

func BenchmarkThreadSafeBufferAppend(b *testing.B) {
	buffer := &ThreadSafeBuffer{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buffer.Append(rand.Intn(1000))
		}
	})
}

func BenchmarkSliceSlice(b *testing.B) {
	slice := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		slice = slice[i:]
	}
}

func BenchmarkThreadSafeBufferRead(b *testing.B) {
	buffer := &ThreadSafeBuffer{}
	for i := 0; i < b.N; i++ {
		buffer.Append(rand.Intn(1000))
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = buffer.Read(rand.Intn(b.N))
		}
	})
}

func main() {
	// Run the benchmarks
	testing.Benchmark(BenchmarkSliceAppend)
	testing.Benchmark(BenchmarkThreadSafeBufferAppend)
	testing.Benchmark(BenchmarkSliceSlice)
	testing.Benchmark(BenchmarkThreadSafeBufferRead)

	// Concurrency safety testing (prevent race conditions)
	fmt.Println("Concurrency Safety Testing...")
	buffer := &ThreadSafeBuffer{}
	var wg sync.WaitGroup

	const concurrencyLevel = 10
	wg.Add(concurrencyLevel + 1)

	// Start concurrent appends
	for i := 0; i < concurrencyLevel; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				buffer.Append(rand.Intn(1000))
			}
		}()
	}

	// Start reading from the buffer simultaneously
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			_, _ = buffer.Read(rand.Intn(10000))
		}
	}()

	wg.Wait()
	fmt.Println("Concurrency Safety Test Completed.")
}