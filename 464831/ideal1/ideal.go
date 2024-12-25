package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"runtime/pprof"
	"os"
)

func dynamicGrowth() {
	// Simulate a slice growing dynamically
	var slice []int
	for i := 0; i < 100000; i++ {
		slice = append(slice, rand.Int())
	}
	fmt.Println("Dynamic Growth: Final length of slice =", len(slice))
}

func preallocate() {
	// Preallocate slice to avoid reallocations
	slice := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		slice = append(slice, rand.Int())
	}
	fmt.Println("Preallocation: Final length of slice =", len(slice))
}

func main() {
	// Profiling setup
	file, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	// Run dynamic growth and measure time
	start := time.Now()
	dynamicGrowth()
	fmt.Println("Time taken for dynamic growth:", time.Since(start))

	// Run preallocation and measure time
	start = time.Now()
	preallocate()
	fmt.Println("Time taken for preallocation:", time.Since(start))
}