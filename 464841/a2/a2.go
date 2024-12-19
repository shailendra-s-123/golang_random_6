package main

import (
	"fmt"
	"log"
	_"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	_"sync/atomic"
	_"time"
)

const (
	numGoroutines = 20
	iterations    = 1_000_000
	sliceSize     = 10_000
)

var (
	syncSlice    sync.Slice // Thread-safe slice
	unsyncSlice = make([]int, 0, sliceSize)
)

// Thread-safe slice using sync.Slice
type syncSliceImpl struct {
	sync.Mutex
	s []int
}

func (s *syncSliceImpl) Append(x ...int) {
	s.Lock()
	s.s = append(s.s, x...)
	s.Unlock()
}

func (s *syncSliceImpl) Slice(start, end int) []int {
	s.Lock()
	defer s.Unlock()
	return s.s[start:end]
}

func (s *syncSliceImpl) Iterate() {
	s.Lock()
	defer s.Unlock()
	for _, v := range s.s {
		// Placeholder for operations
	}
}

func benchmarkSyncSlice() {
	for i := 0; i < iterations; i++ {
		syncSlice.Append(i)
		syncSlice.Slice(0, 10)
		syncSlice.Iterate()
	}
}

func benchmarkUnsyncSlice() {
	for i := 0; i < iterations; i++ {
		unsyncSlice = append(unsyncSlice, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Start synchronous slice benchmark in a new goroutine
	go func() {
		benchmarkSyncSlice()
		wg.Done()
	}()

	// Start asynchronous slice benchmark in a new goroutine
	go func() {
		benchmarkUnsyncSlice()
		wg.Done()
	}()

	// Start CPU profile
	cpuF, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile:", err)
	}
	defer cpuF.Close()
	if err := pprof.StartCPUProfile(cpuF); err != nil {
		log.Fatal("Could not start CPU profile:", err)
	}
	defer pprof.StopCPUProfile()

	// Start memory profile
	runtime.GC() // make sure GC runs before memory profiling
	memoryF, err := os.Create("memory.prof")
	if err != nil {
		log.Fatal("Could not create memory profile:", err)
	}
	defer memoryF.Close()
	pprof.WriteHeapProfile(memoryF)

	// Wait for both benchmarks to complete
	wg.Wait()

	fmt.Println("Finished benchmarks")
}