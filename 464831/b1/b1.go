package main  
import (  
    "fmt"
    "time"
)

func preallocSlice(n int) {
    data := make([]int, 0, n)
    for i := 0; i < n; i++ {
        data = append(data, i)
    }
}

func dynamicSlice(n int) {
    data := []int{}
    for i := 0; i < n; i++ {
        data = append(data, i)
    }
}

func main() {
    n := 10000000

    fmt.Println("Preallocation:")
    start := time.Now()
    preallocSlice(n)
    elapsed := time.Since(start)
    fmt.Println("Time taken:", elapsed)

    fmt.Println("\nDynamic Growth:")
    start = time.Now()
    dynamicSlice(n)
    elapsed = time.Since(start)
    fmt.Println("Time taken:", elapsed)
}