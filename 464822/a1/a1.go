package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const (
	filePath     = "largefile.txt"
	bufferSize    = 1024 * 64 // 64KB buffer size
	numWorkers    = 4
	bufferThread = 2
)

func readBuf(ctx context.Context, reader *bufio.Reader, outputChan chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf := make([]byte, bufferSize)
			n, err := reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Printf("Error reading file: %v\n", err)
				return
			}
			outputChan <- buf[:n]
		}
	}
}

func writeBuf(ctx context.Context, writer *bufio.Writer, inputChan <-chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			writer.Flush() // Ensure any buffered data is written out
			return
		default:
			buf := <-inputChan
			_, err := writer.Write(buf)
			if err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("File not found: %v\n", filePath)
		return
	}

	inputChan := make(chan []byte, bufferThread)
	outputChan := make(chan []byte, bufferThread)

	var wg sync.WaitGroup
	wg.Add(numWorkers + 2) // Number of reader goroutines + 1 writer goroutine + 1 for closing

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for i := 0; i < numWorkers; i++ {
		go readBuf(ctx, reader, inputChan, &wg)
	}

	destinationFile, err := os.CreateTemp("", "out-")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer destinationFile.Close()

	writer := bufio.NewWriter(destinationFile)
	go writeBuf(ctx, writer, outputChan, &wg)

	close(inputChan)
	close(outputChan)
	wg.Wait()

	content, err := ioutil.ReadFile(destinationFile.Name())
	if err != nil {
		fmt.Printf("Error reading destination file: %v\n", err)
		return
	}

	fmt.Println(string(content))
}