package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

const (
	bufferSize = 64 * 1024 // 64 KB buffer
)

func readFileConcurrently(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bufReader := bufio.NewReaderSize(file, bufferSize)

	var wg sync.WaitGroup

	for {
		chunk := make([]byte, bufferSize)
		n, err := bufReader.Read(chunk)

		if n == 0 {
			if err != io.EOF {
				return err
			}
			break
		}

		wg.Add(1)
		go func(chunk []byte) {
			processChunk(chunk)
			wg.Done()
		}(chunk[:n])
	}

	wg.Wait()
	return nil
}

func processChunk(chunk []byte) {
	// Replace this with your processing logic
	for i, b := range chunk {
		// fmt.Printf("%c", b) // Uncomment to print data
		if b == '\n' {
			runtime.GC() // Force GC for demonstration purposes
			fmt.Println("GC triggered after newline")
		}
	}
}

func main() {
	filePath := "example.txt" // Replace with your file path
	err := readFileConcurrently(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
	} else {
		fmt.Println("File processed successfully.")
	}
}
