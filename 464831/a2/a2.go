package main

import (
	_"fmt"
	"testing"
)

func BenchmarkDynamicAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0)
		for j := 0; j < 100000; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkPreallocatedAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100000)
		for j := 0; j < 100000; j++ {
			slice[j] = j
		}
	}
}

func main() {
	testing.Init()
	testing.Main()
}