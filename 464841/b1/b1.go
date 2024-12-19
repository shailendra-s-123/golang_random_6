package main  
import (  
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Benchmark slice operations using safe concurrency patterns
func main() {  
    // Set the number of goroutines for testing
    numGoroutines := 100

    // Set the number of elements to add per goroutine
    elementsPerGoroutine := 10000

    // Create a buffer slice using make to allocate memory
    buffer := make([]int, 0, numGoroutines*elementsPerGoroutine)

    // Initialize a wait group to synchronize concurrent operations
    var wg sync.WaitGroup
    wg.Add(numGoroutines)

    // Start the performance profiling
    start := time.Now()

    // Launch goroutines to append elements concurrently
    for i := 0; i < numGoroutines; i++ {
        go func() {
            for j := 0; j < elementsPerGoroutine; j++ {
                // Safely append an element to the buffer using a mutex
                mutex.Lock()
                buffer = append(buffer, j)
                mutex.Unlock()
            }
            wg.Done()
        }()
    }

    // Wait for all goroutines to complete their operations
    wg.Wait()

    // End the performance profiling
    elapsed := time.Since(start)

    // Print the profiling results
    fmt.Printf("Append operation performance: %s\n", elapsed)
    fmt.Printf("Buffer size: %d\n", len(buffer))

    // Reset the buffer for the next test
    buffer = make([]int, 0, numGoroutines*elementsPerGoroutine)

    // Start the performance profiling for slicing
    start = time.Now()

    // Launch goroutines to slice the buffer concurrently
    for i := 0; i < numGoroutines; i++ {
        go func() {
            for j := 0; j < elementsPerGoroutine; j++ {
                // Safely slice the buffer using a mutex
                mutex.Lock()
                slice := buffer[:j+1]
                mutex.Unlock()

                // Perform some operation on the sliced data (e.g., sum)
                sum := 0
                for _, value := range slice {
                    sum += value
                }
            }
            wg.Done()
        }()
    }

    // Wait for all goroutines to complete their operations
    wg.Wait()

    // End the performance profiling for slicing
    elapsed = time.Since(start)

    // Print the profiling results for slicing
    fmt.Printf("Slice operation performance: %s\n", elapsed)

    // Reset the buffer for the next test
    buffer = make([]int, numGoroutines*elementsPerGoroutine)

    // Start the performance profiling for iteration
    start = time.Now()

    // Launch goroutines to iterate over the buffer concurrently
    for i := 0; i < numGoroutines; i++ {
        go func() {
            for j := 0; j < elementsPerGoroutine; j++ {
                // Safely iterate over the buffer using a mutex
                mutex.Lock()
                for _, value := range buffer {
                    _ = value // Do something with the value
                }
                mutex.Unlock()
            }
            wg.Done()
        }()
    }

    // Wait for all goroutines to complete their operations
    wg.Wait()

    // End the performance profiling for iteration
    elapsed = time.Since(start)

    // Print the profiling results for iteration
    fmt.Printf("Iteration operation performance: %s\n", elapsed)
}

// Use a sync.Mutex to ensure safe access to the shared buffer
var mutex sync.Mutex