package main  
import (  
    "fmt"
    "log"
    "math/rand"
    "os"
    "runtime"
    "runtime/pprof"
    "sync"
    _"time"
)

const (  
    numGoroutines = 100
    numIterations = 100000
)

var (  
    slice  = make([]int, 0, numIterations*numGoroutines)
    sliceMutex sync.RWMutex
)

func appendToSlice(wg *sync.WaitGroup) {  
    defer wg.Done()
    for i := 0; i < numIterations; i++ {
        sliceMutex.Lock()
        slice = append(slice, rand.Intn(1000))
        sliceMutex.Unlock()
    }
}

func sliceAndIterate(wg *sync.WaitGroup) {  
    defer wg.Done()
    for i := 0; i < numIterations; i++ {
        sliceMutex.RLock()
        length := len(slice)
        randomIndex := rand.Intn(length)
        // Slicing using a random index
        slicedSlice := slice[randomIndex:length]
        sliceMutex.RUnlock()
        // Iterating over the sliced slice
        for _, value := range slicedSlice {
            _ = value
        }
    }
}

func main() {  
    fmt.Println("Starting profile...")
    
    // Start CPU profiling
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal("Could not create CPU profile:", err)
    }
    defer f.Close()
    if err := pprof.StartCPUProfile(f); err != nil {
        log.Fatal("Could not start CPU profile:", err)
    }
    defer pprof.StopCPUProfile()
    
    wg := &sync.WaitGroup{}
    wg.Add(2 * numGoroutines)

    // Start goroutines for concurrent slice operations
    for i := 0; i < numGoroutines; i++ {
        go appendToSlice(wg)
        go sliceAndIterate(wg)
    }

    // Wait for all goroutines to finish
    wg.Wait()
    
    fmt.Println("Finished benchmarks")
    
    // Start memory profiling
    runtime.GC()
    memoryF, err := os.Create("memory.prof")
    if err != nil {
        log.Fatal("Could not create memory profile:", err)
    }
    defer memoryF.Close()
    pprof.WriteHeapProfile(memoryF)
}
