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

// readFileConcurrently reads a file in chunks concurrently
func readFileConcurrently(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bufReader := bufio.NewReaderSize(file, bufferSize)
	var wg sync.WaitGroup

	for {
		// Create a buffer to hold the chunk
		chunk := make([]byte, bufferSize)
		n, err := bufReader.Read(chunk)

		// Break loop on EOF or handle read errors
		if n == 0 {
			if err != io.EOF {
				return err
			}
			break
		}

		// Process the chunk concurrently
		wg.Add(1)
		go func(chunk []byte) {
			defer wg.Done()
			processChunk(chunk)
		}(chunk[:n])
	}

	// Wait for all goroutines to complete
	wg.Wait()
	return nil
}

// processChunk processes a chunk of data
func processChunk(chunk []byte) {
	// Iterate over the bytes in the chunk
	for _, b := range chunk {
		if b == '\n' {
			// Demonstration: Trigger garbage collection on newline
			runtime.GC()
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
