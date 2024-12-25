package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime/pprof"
	"os"
	"log"
)

func exponentialGrowth() {
	// Simulate slice growing dynamically with exponential growth
	var slice []int
	for i := 0; i < 100000; i++ {
		slice = append(slice, rand.Int())
	}
	fmt.Println("Exponential Growth: Final length of slice =", len(slice))
}

func fixedSizeIncrements(increment int) {
	// Simulate slice growing with fixed-size increments
	var slice []int
	for i := 0; i < 100000; i++ {
		slice = append(slice, rand.Int())
		if len(slice)%increment == 0 {
			capacity := len(slice) + increment
			if capacity > cap(slice) {
				slice = make([]int, len(slice), capacity)
			}
		}
	}
	fmt.Println("Fixed Size Increments (" + fmt.Sprintf("%d", increment) + "): Final length of slice =", len(slice))
}

func main() {
	// Profiling setup
	file, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	// Run exponential growth and measure time
	start := time.Now()
	exponentialGrowth()
	fmt.Println("Time taken for exponential growth:", time.Since(start))

	// Run fixed-size increments with an increment of 1000
	start = time.Now()
	fixedSizeIncrements(1000)
	fmt.Println("Time taken for fixed size increments (1000):", time.Since(start))

	// Run fixed-size increments with an increment of 100
	start = time.Now()
	fixedSizeIncrements(100)
	fmt.Println("Time taken for fixed size increments (100):", time.Since(start))
}