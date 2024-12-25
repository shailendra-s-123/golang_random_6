package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "large_dataset.txt" // Replace with the path to your dataset

	err := processLargeDataset(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing dataset: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Dataset processed successfully.")
}

func processLargeDataset(filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReaderSize(file, 16*1024) // 16 KB buffer size

	lineCount := 0
	for {
		// Read a line from the file
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" { // Handle end of file gracefully
				break
			}
			return fmt.Errorf("failed to read line: %w", err)
		}

		// Process the line (e.g., counting lines)
		processLine(line)
		lineCount++
	}

	fmt.Printf("Processed %d lines in total.\n", lineCount)
	return nil
}

func processLine(line string) {
	// Simulated processing (e.g., log the first 50 characters of the line)
	if len(line) > 50 {
		line = line[:50] + "..."
	}
	fmt.Printf("Processed line: %s\n", line)
}
