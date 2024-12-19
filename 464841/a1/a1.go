package main

import (
	"fmt"
	"log"
	_"net/http"
	"runtime/pprof"
	"sync"
	_"time"
)

var (
	syncSlice  = make([]int, 0, 10000)
	lock       = sync.Mutex{}
	unsyncSlice = make([]int, 0, 10000)
)

func benchmarkSyncSlice() {
	for i := 0; i < 1000000; i++ {
		lock.Lock()
		syncSlice = append(syncSlice, i)
		lock.Unlock()
	}
}

func benchmarkUnsyncSlice() {
	for i := 0; i < 1000000; i++ {
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
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile:", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
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