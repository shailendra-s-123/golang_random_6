package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/golang/snappy"
)

const bufferSize = 64 * 1024

func main() {
	start := time.Now()

	// Compress the input file
	err := compress("input.txt", "tmp.snappy")
	if err != nil {
		log.Fatalf("Compression failed: %v", err)
	}

	// Decompress the file
	err = decompress("tmp.snappy", "output.txt")
	if err != nil {
		log.Fatalf("Decompression failed: %v", err)
	}

	log.Printf("File operations completed in %v\n", time.Since(start))
}

func compress(inputPath, tmpPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening input file: %w", err)
	}
	defer inputFile.Close()

	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}
	defer tmpFile.Close()

	writer := snappy.NewBufferedWriter(tmpFile)
	defer writer.Close()

	reader := bufio.NewReaderSize(inputFile, bufferSize)
	for {
		buf := make([]byte, bufferSize)
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading input file: %w", err)
		}
		if _, err := writer.Write(buf[:n]); err != nil {
			return fmt.Errorf("error writing compressed data: %w", err)
		}
	}
	return nil
}

func decompress(tmpPath, outputPath string) error {
	tmpFile, err := os.Open(tmpPath)
	if err != nil {
		return fmt.Errorf("error opening temporary file: %w", err)
	}
	defer tmpFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer outputFile.Close()

	reader := snappy.NewReader(tmpFile)
	if _, err := io.Copy(outputFile, reader); err != nil {
		return fmt.Errorf("error writing decompressed data: %w", err)
	}
	return nil
}
