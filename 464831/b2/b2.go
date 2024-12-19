package main

import (
	_"fmt"
	"testing"
)

func BenchmarkDynamicGrowth(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dynamicSlice []int
		for i := 0; i < 1000000; i++ {
			dynamicSlice = append(dynamicSlice, i)
		}
	}
}

func BenchmarkPreallocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		preallocatedSlice := make([]int, 1000000)
		for i := 0; i < 1000000; i++ {
			preallocatedSlice[i] = i
		}
	}
}

func main() {
	testing.Benchmark(BenchmarkDynamicGrowth)
	testing.Benchmark(BenchmarkPreallocation)
}