package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const bufferSize = 4096

func readFileEfficiently(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, bufferSize)

	var data []byte
	for {
		n, err := reader.Read(data[:cap(data)])
		if n == 0 {
			break // End of file
		}
		if err != nil && err != io.EOF {
			return err
		}

		data = data[:n] // Slice the data to the actual number of bytes read
		// Process the data here
		fmt.Println(string(data))
	}

	return nil
}

func writeFileEfficiently(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, bufferSize)

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return writer.Flush()
}

func main() {
	start := time.Now()

	// Read a file
	err := readFileEfficiently("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Write data to a file
	data := []byte("Hello, world!")
	err = writeFileEfficiently("output.txt", data)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Println("File I/O completed in:", elapsed)
}